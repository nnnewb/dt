package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nnnewb/dt/internal/client"
	"github.com/nnnewb/dt/internal/middleware"
	"github.com/nnnewb/dt/internal/svc/bank"
	"github.com/nnnewb/dt/internal/svc/dm"
	"go.opentelemetry.io/otel"
)

func main() {
	tp := middleware.InitTracer("app")
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

	r := gin.Default()
	r.Use(middleware.Jaeger("app"))
	r.POST("/v1alpha1/transfer", transfer)
	r.Run(":5000")
}

type TransferReq struct {
	ID     int64 `json:"id,omitempty"`
	Bank   int32 `json:"bank,omitempty"`
	Amount int64 `json:"amount,omitempty"`
	ToID   int64 `json:"to_id,omitempty"`
	ToBank int32 `json:"to_bank,omitempty"`
}

type GeneralResp struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func transfer(c *gin.Context) {
	req := &TransferReq{}
	c.BindJSON(req)

	dmcli := client.NewDMClient("http://dm:5000")
	gid := dm.MustGenGID()
	resp, err := dmcli.CreateGlobalTx(c, &dm.CreateGlobalTxReq{GID: gid})
	if err != nil {
		log.Printf("create global tx failed, error %v", err)
		c.JSONP(500, &GeneralResp{
			Code:    -1,
			Message: "create global transaction failed",
		})
		return
	}

	if resp.Code != 0 {
		log.Printf("create global tx failed, code %d, %s", resp.Code, resp.Message)
		c.JSONP(500, &GeneralResp{
			Code:    -1,
			Message: "create global transaction failed",
		})
		return
	}

	cli1 := client.NewBankClient("http://bank1:5000")
	transInResp, err := cli1.TransIn(c, &bank.TransInReq{ID: req.ToID, Amount: req.Amount})
	if err != nil {
		// 失败的话就等着超时
		_, _ = dmcli.RollbackGlobalTx(c, &dm.RollbackGlobalTxReq{GID: gid})
		c.JSONP(500, &GeneralResp{
			Code:    -1,
			Message: "trans in failed",
		})
		return
	}

	if transInResp.Code != 0 {
		// 失败的话就等着超时
		_, _ = dmcli.RollbackGlobalTx(c, &dm.RollbackGlobalTxReq{GID: gid})

		c.JSONP(500, &GeneralResp{
			Code:    -1,
			Message: "trans in failed",
		})
		return
	}

	cli2 := client.NewBankClient("http://bank2:5000")
	transOutResp, err := cli2.TransIn(c, &bank.TransInReq{ID: req.ToID, Amount: req.Amount})
	if err != nil {
		// 失败的话就等着超时
		_, _ = dmcli.RollbackGlobalTx(c, &dm.RollbackGlobalTxReq{GID: gid})

		c.JSONP(500, &GeneralResp{
			Code:    -1,
			Message: "trans out failed",
		})
		return
	}

	if transOutResp.Code != 0 {
		// 失败的话就等着超时
		_, _ = dmcli.RollbackGlobalTx(c, &dm.RollbackGlobalTxReq{GID: gid})

		c.JSONP(500, &GeneralResp{
			Code:    -1,
			Message: "trans out failed",
		})
		return
	}

	commitResp, err := dmcli.CommitGlobalTx(c, &dm.CommitGlobalTxReq{GID: gid})
	if err != nil {
		// 失败的话就等着超时
		_, _ = dmcli.RollbackGlobalTx(c, &dm.RollbackGlobalTxReq{GID: gid})

		c.JSONP(500, &GeneralResp{
			Code:    -1,
			Message: "commit failed",
		})
		return
	}

	if commitResp.Code != 0 {
		// 失败的话就等着超时
		_, _ = dmcli.RollbackGlobalTx(c, &dm.RollbackGlobalTxReq{GID: gid})

		c.JSONP(500, &GeneralResp{
			Code:    -1,
			Message: "commit failed",
		})
		return
	}

	c.JSONP(200, map[string]interface{}{
		"code":    0,
		"message": "ok",
	})

}
