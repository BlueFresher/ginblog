package model

import (
	"fmt"
	"ginblog/utils"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// 配置数据库参数

var db *gorm.DB
var err error

func InitDb() {
	db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser, utils.DbPassword, utils.DbHost, utils.DbPort, utils.DbName))

	if err != nil {
		fmt.Println("连接数据库失败请检查：", err)
	}
	//
	db.SingularTable(true)
	db.AutoMigrate(&User{}, &Article{}, &Category{}, &Profile{}, &Comment{})

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(10 * time.Second)
}
