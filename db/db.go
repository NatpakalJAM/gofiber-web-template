package db

import (
	"fmt"
	"gofiber-web-template/cfg"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dbType = "mysql"
)

var db *gorm.DB

//Init init db
func Init() {
	var err error
	dbCfg := cfg.C.DB

	switch dbType {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local",
			dbCfg.Username, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Database)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		break
	}
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}

}
