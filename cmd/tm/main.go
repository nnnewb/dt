package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/nnnewb/dt/internal/client"
	"github.com/nnnewb/dt/internal/middleware"
	"github.com/nnnewb/dt/internal/svc/bank"
	"github.com/nnnewb/dt/internal/svc/tm"
	"github.com/nnnewb/dt/internal/tracing/otelsql"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
)

func main() {
	tp := middleware.InitTracer("tm")
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

	driver := "otelmysql"
	dsn := "root:root@tcp(mysql:3306)/tm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		log.Fatal(err)
		return
	}

	r := gin.Default()
	r.Use(middleware.Jaeger("tm"))
	r.Use(middleware.WithDatabase(db))
	r.Use(middleware.LogPayloadAndResponse)
	r.POST("/v1alpha1/create_global_tx", createGlobalTx)
	r.POST("/v1alpha1/register_local_tx", registerLocalTx)
	r.POST("/v1alpha1/commit_global_tx", commitGlobalTx)
	r.POST("/v1alpha1/rollback_global_tx", rollbackGlobalTx)
	r.Run(":5000")
}

func createGlobalTx(c *gin.Context) {
	req := &tm.CreateGlobalTxReq{}
	c.BindJSON(req)

	db := c.MustGet("db").(*sqlx.DB)
	_, err := db.NamedExecContext(c.Request.Context(), `INSERT INTO global_tx(gid) VALUES(:gid)`, &tm.GlobalTx{GID: req.GID})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSONP(200, map[string]interface{}{
		"code":    0,
		"message": "ok",
	})
}

func registerLocalTx(c *gin.Context) {
	req := &tm.RegisterLocalTxReq{}
	c.BindJSON(req)

	db := c.MustGet("db").(*sqlx.DB)
	_, err := db.NamedExecContext(
		c.Request.Context(),
		`INSERT INTO local_tx(gid,branch_id,callback_url) values(:gid, :branch_id, :callback_url)`,
		&tm.LocalTx{
			GID:         req.GID,
			BranchID:    req.BranchID,
			CallbackUrl: req.CallbackUrl,
		},
	)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSONP(200, map[string]interface{}{
		"code":    0,
		"message": "ok",
	})
}

func commitGlobalTx(c *gin.Context) {
	req := &tm.CommitGlobalTxReq{}
	c.BindJSON(req)

	db := c.MustGet("db").(*sqlx.DB)
	allLocalTx := make([]tm.LocalTx, 0)
	err := db.SelectContext(c.Request.Context(), &allLocalTx, "SELECT * FROM local_tx WHERE gid=?", req.GID)
	if err != nil {
		c.Error(err)
		return
	}

	// TODO 极端情况下，回调 RM 时出现部分失败要如何处理？
	cli := &http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
	for _, tx := range allLocalTx {
		callbackPayload := &bank.TMCallbackReq{Action: "commit", GID: req.GID, BranchID: tx.BranchID}
		callbackResp := bank.TMCallbackResp{}
		err = client.WrappedPost(c.Request.Context(), cli, tx.CallbackUrl, callbackPayload, &callbackResp)
		if err != nil {
			c.Error(err)
			return
		}

		if callbackResp.Code != 0 {
			c.JSONP(500, &tm.CommitGlobalTxResp{
				Code:    -1,
				Message: fmt.Sprintf("commit local tx failed, response code %d, %s", callbackResp.Code, callbackResp.Message),
			})
			return
		}
	}

	c.JSONP(200, map[string]interface{}{
		"code":    0,
		"message": "ok",
	})
}

func rollbackGlobalTx(c *gin.Context) {
	req := &tm.RollbackGlobalTxReq{}
	c.BindJSON(req)

	db := c.MustGet("db").(*sqlx.DB)
	allLocalTx := make([]tm.LocalTx, 0)
	err := db.SelectContext(c.Request.Context(), &allLocalTx, "SELECT * FROM local_tx WHERE gid=?", req.GID)
	if err != nil {
		c.Error(err)
		return
	}

	// TODO 极端情况下，回调 RM 时出现部分失败要如何处理？
	cli := &http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
	for _, tx := range allLocalTx {
		callbackPayload := &bank.TMCallbackReq{Action: "rollback", GID: req.GID, BranchID: tx.BranchID}
		callbackResp := bank.TMCallbackResp{}
		err = client.WrappedPost(c.Request.Context(), cli, tx.CallbackUrl, callbackPayload, &callbackResp)
		if err != nil {
			c.Error(err)
			return
		}

		if callbackResp.Code != 0 {
			c.JSONP(500, &tm.RollbackGlobalTxResp{
				Code:    -1,
				Message: fmt.Sprintf("rollback local tx failed, response code %d, %s", callbackResp.Code, callbackResp.Message),
			})
			return
		}
	}

	c.JSONP(200, map[string]interface{}{
		"code":    0,
		"message": "ok",
	})
}
