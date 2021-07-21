package database

import (
	"log"

	"Pigeon/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var connectError error

func init() {
	DB, connectError = gorm.Open("mysql", config.MysqlDefaultConfig.ToConfigString())
	if connectError != nil {
		log.Fatal(connectError.Error())
	}
}
