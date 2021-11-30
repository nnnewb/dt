package main

import (
	"fmt"

	"github.com/nnnewb/dt/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(mysql:3306)/shop?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// sync table schema every time
	db.AutoMigrate(&models.Order{})

	fmt.Println("vim-go")
}
