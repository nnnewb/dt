package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/nnnewb/dt/internal/svc/dm"
)

func main() {
	dsn := "root:root@tcp(mysql:3306)/dm?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.POST("/v1alpha1/create_global_tx", func(c *gin.Context) {
		req := &dm.CreateGlobalTxReq{}
		c.BindJSON(req)

		// TODO implements this

		c.JSONP(200, map[string]interface{}{
			"code":    0,
			"message": "ok",
		})
	})
	r.POST("/v1alpha1/register_local_tx", func(c *gin.Context) {
		req := &dm.RegisterLocalTxReq{}
		c.BindJSON(req)

		// TODO implements this

		c.JSONP(200, map[string]interface{}{
			"code":    0,
			"message": "ok",
		})
	})
	r.POST("/v1alpha1/commit_global_tx", func(c *gin.Context) {
		req := &dm.CommitGlobalTxReq{}
		c.BindJSON(req)

		// TODO implements this

		c.JSONP(200, map[string]interface{}{
			"code":    0,
			"message": "ok",
		})
	})
	r.POST("/v1alpha1/rollback_global_tx", func(c *gin.Context) {
		req := &dm.RollbackGlobalTxReq{}
		c.BindJSON(req)

		// TODO implements this

		c.JSONP(200, map[string]interface{}{
			"code":    0,
			"message": "ok",
		})
	})
	r.Run(":5000")
}
