package model


//定义用户
type User struct {
	Model

	Username 	string	`gorm:"size:32;unique;not null;unique"`		// 改变默认长度（size）,非空, 唯一
	Password	string	`gorm:"size:32;not null"`

}