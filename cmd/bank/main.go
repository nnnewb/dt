package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/nnnewb/dt/internal/client"
	"github.com/nnnewb/dt/internal/middleware"
	"github.com/nnnewb/dt/internal/svc/bank"
	"github.com/nnnewb/dt/internal/svc/tm"
	"github.com/nnnewb/dt/internal/tracing/otelsql"
	"go.opentelemetry.io/otel"
)

var flgBankID int

func init() {
	flag.IntVar(&flgBankID, "bank-id", 0, "Bank ID for this server")
}

func main() {
	flag.Parse()
	if flgBankID == 0 {
		flag.Usage()
		return
	}

	tp := middleware.InitTracer(fmt.Sprintf("bank%d", flgBankID))
	otel.SetTracerProvider(tp)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	defer func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	otelsql.Register("otelmysql", &mysql.MySQLDriver{})
	dsn := fmt.Sprintf("root:root@tcp(mysql:3306)/bank%d?charset=utf8mb4&parseTime=True&loc=Local", flgBankID)
	db, err := sqlx.Open("otelmysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.Use(middleware.Jaeger(fmt.Sprintf("bank%d", flgBankID)))
	r.Use(middleware.WithDatabase(db))
	r.Use(middleware.LogPayloadAndResponse)
	r.POST("/v1alpha1/tm_callback", tmCallback)
	r.POST("/v1alpha1/trans_in", transIn)
	r.POST("/v1alpha1/trans_out", transOut)
	r.Run(":5000")
}

func tmCallback(c *gin.Context) {
	req := &bank.TMCallbackReq{}
	c.BindJSON(req)
	db := c.MustGet("db").(*sqlx.DB)

	// 业务逻辑
	xid := fmt.Sprintf("'%s','%s'", req.GID, req.BranchID)
	if req.Action == "commit" {
		// 提交 XA 事务
		_, err := db.ExecContext(c.Request.Context(), fmt.Sprintf("XA COMMIT %s", xid))
		if err != nil {
			c.Error(err)
			return
		}
	} else if req.Action == "rollback" {
		// 回滚 XA 事务
		_, err := db.ExecContext(c.Request.Context(), fmt.Sprintf("XA ROLLBACK %s", xid))
		if err != nil {
			c.Error(err)
			return
		}
	} else {
		c.JSONP(400, &bank.TMCallbackResp{
			Code:    1001,
			Message: "unknown action",
		})
		return
	}

	c.JSONP(200, map[string]interface{}{
		"code":    0,
		"message": "ok",
	})
}

func transIn(c *gin.Context) {
	req := &bank.TransInReq{}
	c.BindJSON(req)
	db := c.MustGet("db").(*sqlx.DB)
	cli := client.NewTMClient("http://tm:5000")
	branchID := tm.MustGenBranchID("TransIn")

	resp, err := cli.RegisterLocalTx(c.Request.Context(), &tm.RegisterLocalTxReq{
		GID:         req.GID,
		BranchID:    branchID,
		CallbackUrl: fmt.Sprintf("http://bank%d:5000/v1alpha1/tm_callback", flgBankID),
	})
	if err != nil {
		c.Error(err)
		return
	}
	if resp.Code != 0 {
		c.JSONP(500, &bank.TransInResp{
			Code:    -1,
			Message: "register local tx failed",
		})
		return
	}

	// 业务逻辑
	// 准备 XA 事务
	xid := fmt.Sprintf("'%s','%s'", req.GID, branchID)
	_, err = db.ExecContext(c.Request.Context(), fmt.Sprintf("XA BEGIN %s", xid))
	if err != nil {
		c.Error(err)
		return
	}

	_, err = db.ExecContext(c.Request.Context(), "UPDATE wallet SET balance=balance+? WHERE id=?", req.Amount, req.ID)
	if err != nil {
		c.Error(err)
		return
	}

	_, err = db.ExecContext(c.Request.Context(), fmt.Sprintf("XA END %s", xid))
	if err != nil {
		c.Error(err)
		return
	}

	_, err = db.ExecContext(c.Request.Context(), fmt.Sprintf("XA PREPARE %s", xid))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSONP(200, map[string]interface{}{
		"code":    0,
		"message": "ok",
	})
}

func transOut(c *gin.Context) {
	req := &bank.TransOutReq{}
	c.BindJSON(req)
	db := c.MustGet("db").(*sqlx.DB)
	cli := client.NewTMClient("http://tm:5000")
	branchID := tm.MustGenBranchID("TransOut")

	// 注册本地事务
	resp, err := cli.RegisterLocalTx(c.Request.Context(), &tm.RegisterLocalTxReq{
		GID:         req.GID,
		BranchID:    branchID,
		CallbackUrl: fmt.Sprintf("http://bank%d:5000/v1alpha1/tm_callback", flgBankID),
	})
	if err != nil {
		c.Error(err)
		return
	}
	if resp.Code != 0 {
		c.JSONP(500, &bank.TransInResp{
			Code:    -1,
			Message: "register local tx failed",
		})
		return
	}

	// 业务逻辑
	// 准备 XA 事务
	xid := fmt.Sprintf("'%s','%s'", req.GID, branchID)
	_, err = db.ExecContext(c.Request.Context(), fmt.Sprintf("XA BEGIN %s", xid))
	if err != nil {
		c.Error(err)
		return
	}

	_, err = db.ExecContext(c.Request.Context(), "UPDATE wallet SET balance=balance-? WHERE id=?", req.Amount, req.ID)
	if err != nil {
		c.Error(err)
		return
	}

	_, err = db.ExecContext(c.Request.Context(), fmt.Sprintf("XA END %s", xid))
	if err != nil {
		c.Error(err)
		return
	}

	_, err = db.ExecContext(c.Request.Context(), fmt.Sprintf("XA PREPARE %s", xid))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSONP(200, map[string]interface{}{
		"code":    0,
		"message": "ok",
	})
}
