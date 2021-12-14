package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/nnnewb/dt/internal/middleware"
	"github.com/nnnewb/dt/internal/svc/bank"
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

	dsn := fmt.Sprintf("root:root@tcp(mysql:3306)/bank%d?charset=utf8mb4&parseTime=True&loc=Local", flgBankID)
	_, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.Use(middleware.Jaeger(fmt.Sprintf("bank%d", flgBankID)))
	r.POST("/v1alpha1/dm_callback", dmCallback)
	r.POST("/v1alpha1/trans_in", transIn)
	r.POST("/v1alpha1/trans_out", transOut)
	r.Run(":5000")
}

func dmCallback(c *gin.Context) {
	req := &bank.TransInReq{}
	c.BindJSON(req)

	// TODO implements this

	c.JSONP(200, map[string]interface{}{
		"code":    0,
		"message": "ok",
	})
}

func transIn(c *gin.Context) {
	req := &bank.TransInReq{}
	c.BindJSON(req)

	// TODO implements this

	c.JSONP(200, map[string]interface{}{
		"code":    0,
		"message": "ok",
	})
}

func transOut(c *gin.Context) {
	req := &bank.TransOutReq{}
	c.BindJSON(req)

	// TODO implements this

	c.JSONP(200, map[string]interface{}{
		"code":    0,
		"message": "ok",
	})
}
