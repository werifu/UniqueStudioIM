package model

import (
	"fmt"
	"im/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//定义用户
type User struct {
	gorm.Model

	Username 	string	`gorm:"size:32;unique;not null;unique"`		// 改变默认长度（size）,非空, 唯一
	Password	string	`gorm:"size:32;not null"`

}

type Room struct {
	gorm.Model

	Name 		string	`gorm:"size:16;unique;not null;unique"`		// 改变默认长度（size）,非空, 唯一
	Password	int

}

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

	db, err := gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
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