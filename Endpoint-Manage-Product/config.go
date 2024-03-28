package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDB() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("coba")
	fmt.Println("coba")
	DB = db
	return DB, nil
}
