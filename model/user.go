package model

import (
	"thchat/pkg/e"
	"time"
)


//定义用户
type User struct {
	Model

	Username 	string	`gorm:"size:32;unique;not null;unique"`		// 改变默认长度（size）,非空, 唯一
	Password	string	`gorm:"size:32;not null"`
}

func (User) TableName() string {
	return "users"
}

// 插入User表
func AddUser(username, password string) bool {
	user := User{
		Model:    Model{
			ModifiedAt: time.Now(),
		},
		Username: username,
		Password: password,
	}
	db.Create(&user)
	return true
}

// User是否已经存在
func UserExists(username string) bool {
	user := User{}
	db.Where("username = ?", username).First(&user)

	if user.Username == username {
		return true
	}

	return false
}

func LoginCheck(username, password string) int {
	user := User{}
	db.Where("username = ?", username).Select("username, password").First(&user)

	if user.Username == username {
		if user.Password == password {
			return e.SUCCESS
		} else {
			return e.ErrUserPassword
		}
	} else {
		return e.ErrUserNotExists
	}
}