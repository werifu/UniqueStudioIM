package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"thchat/pkg/config"
	"thchat/pkg/logging"
	"time"
)

// 自定义model
type Model struct {
	ID int 						`gorm:"primary_key" json:"id"`
	CreatedAt time.Time			`gorm:"created_at" json:"created_at"`
	ModifiedAt time.Time 		`gorm:"modified_at" json:"modified_at"`
}



var db *gorm.DB

func init(){
	dbType := config.AppConfig.DataBase.Type
	dbName := config.AppConfig.DataBase.Name
	user := config.AppConfig.DataBase.User
	password := config.AppConfig.DataBase.Password
	host := config.AppConfig.DataBase.Host
	var err error
	//fmt.Println(host)
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		logging.Error(err)
		panic(err)
	}
	//初始化表格
	if !db.HasTable("users") {
		db.CreateTable(&User{})
	}


	db.SingularTable(true)

	// 设置连接池接入量
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	//fmt.Println("db init ok")
}