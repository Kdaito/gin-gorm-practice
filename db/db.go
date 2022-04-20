package db

import (
	"example/api-template/config"
	"example/api-template/entity"
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
	err error
)

// データベースの初期化
func Init() {
	dbConnectInfo := fmt.Sprintf(
		`%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local`,
		config.Config.DbUserName,
		config.Config.DbUserPassword,
		config.Config.DbHost,
		config.Config.DbPort,
		config.Config.DbName,
	)

	db, err = gorm.Open(config.Config.DbDriverName, dbConnectInfo)

	if err != nil {
		fmt.Printf("failed connect database: %v", err)
	} else {
		fmt.Println("Successfully connect database...")
	}

	autoMigration()
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	if err := db.Close(); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Successfully disconnect database...")
	}
}

func autoMigration() {
	db.AutoMigrate(&entity.User{})
}