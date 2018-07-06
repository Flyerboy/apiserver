package model


import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"apiserver/config"
	"github.com/jinzhu/gorm"
)


func GetDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@/%s?charset=utf8", config.Get("db.user"), config.Get("db.password"), config.Get("db.name"))
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return db
}