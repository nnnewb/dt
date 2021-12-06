package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nnnewb/dt/internal/svc/bank"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// sync table schema every time
	db.AutoMigrate(&bank.Wallet{})

	r := gin.Default()
	r.POST("/v1alpha1/bank/:BankID/TransIn", func(c *gin.Context) {})
	r.POST("/v1alpha1/bank/:BankID/TransOut", func(c *gin.Context) {})
	r.Run()
}
