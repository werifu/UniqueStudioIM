package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"im/pkg/setting"
	"log"
)

// 自定义model
type Model struct {
	ID int 				`gorm:"primary_key" json:"id"`
	CreatedOn int 		`json:"created_on"`
	ModifiedOn int 		`json:"modified_on"`
}



var (
	db *gorm.DB
)

func init() {
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "fail to get section 'database': %v", err)
	}

	dbType := sec.Key("type").String()
	dbName := sec.Key("name").String()
	user := sec.Key("user").String()
	password := sec.Key("password").String()
	host := sec.Key("host").String()
	tablePrefix := sec.Key("table_frefix").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName;
	}

	db.SingularTable(true)

	// 设置连接池接入量
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}