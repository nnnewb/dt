package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/nnnewb/dt/internal/middleware"
	"github.com/nnnewb/dt/internal/svc/dm"
	"go.opentelemetry.io/otel"
)

func main() {
	tp := middleware.InitTracer("dm")
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

	dsn := "root:root@tcp(mysql:3306)/dm?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.Use(middleware.Jaeger("dm"))
	r.POST("/v1alpha1/create_global_tx", createGlobalTx)
	r.POST("/v1alpha1/register_local_tx", registerLocalTx)
	r.POST("/v1alpha1/commit_global_tx", commitGlobalTx)
	r.POST("/v1alpha1/rollback_global_tx", rollbackGlobalTx)
	r.Run(":5000")
}

func createGlobalTx(c *gin.Context) {
	req := &dm.CreateGlobalTxReq{}
	c.BindJSON(req)

	// TODO implements this

	c.JSONP(200, map[string]interface{}{
		"code":    0,
		"message": "ok",
	})
}

func registerLocalTx(c *gin.Context) {
	req := &dm.RegisterLocalTxReq{}
	c.BindJSON(req)

	// TODO implements this

	c.JSONP(200, map[string]interface{}{
		"code":    0,
		"message": "ok",
	})
}

func commitGlobalTx(c *gin.Context) {
	req := &dm.CommitGlobalTxReq{}
	c.BindJSON(req)

	// TODO implements this

	c.JSONP(200, map[string]interface{}{
		"code":    0,
		"message": "ok",
	})
}

func rollbackGlobalTx(c *gin.Context) {
	req := &dm.RollbackGlobalTxReq{}
	c.BindJSON(req)

	// TODO implements this

	c.JSONP(200, map[string]interface{}{
		"code":    0,
		"message": "ok",
	})
}
