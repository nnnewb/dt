package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/nnnewb/dt/internal/svc/bank"
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

	dsn := fmt.Sprintf("root:root@tcp(mysql:3306)/bank%d?charset=utf8mb4&parseTime=True&loc=Local", flgBankID)
	_, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.POST("/v1alpha1/bank/:BankID/TransIn", func(c *gin.Context) {
		req := &bank.TransInReq{}
		c.BindJSON(req)

		// TODO implements this

		c.JSONP(200, map[string]interface{}{
			"code":    0,
			"message": "ok",
		})
	})
	r.POST("/v1alpha1/bank/:BankID/TransOut", func(c *gin.Context) {
		req := &bank.TransOutReq{}
		c.BindJSON(req)

		// TODO implements this

		c.JSONP(200, map[string]interface{}{
			"code":    0,
			"message": "ok",
		})
	})
	r.Run(":5000")
}
