package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nnnewb/dt/internal/svc/dm"
)

func main() {
	r := gin.Default()
	r.POST("/v1alpha1/create-global-transaction", func(c *gin.Context) {
		req := &dm.CreateGlobalTransactionReq{}
		c.BindJSON(req)
	})
	r.POST("/v1alpha1/register-local-transaction", func(c *gin.Context) {
		req := &dm.RegisterLocalTransactionReq{}
		c.BindJSON(req)
	})
	r.POST("/v1alpha1/commit-global-transaction", func(c *gin.Context) {
		req := &dm.CommitGlobalTransactionReq{}
		c.BindJSON(req)
	})
	r.POST("/v1alpha1/rollback-global-transaction", func(c *gin.Context) {
		req := &dm.RollbackGlobalTransactionReq{}
		c.BindJSON(req)
	})
	r.Run()
}
