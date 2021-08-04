package database

import (
	"log"

	"Pigeon/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var connectError error

func init() {
	db, connectError = gorm.Open("mysql", config.MysqlDefaultConfig.ToConfigString())
	if connectError != nil {
		log.Fatal(connectError.Error())
	}
}

func GetDB() *gorm.DB {
	return db
}
