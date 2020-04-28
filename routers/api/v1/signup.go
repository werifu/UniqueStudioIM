package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"im/model"
	"im/pkg/err"
	"log"
	"net/http"
)

func PostSignup(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	log.Println("注册参数：", username, password)

	valid := validation.Validation{}
	valid.Required(username, "name").Message("名称不能为空")
	valid.MaxSize(username, 20, "name").Message("名称最长20字符")
	valid.Required(password, "password").Message("密码不能为空")
	valid.Range(password,5, 40, "password").Message("密码长度在5到40个字符长度区间")

	var code int
	if model.UserExists(username) {
		log.Println(username)
		code = err.ErrUserExists
	} else {
		model.AddUser(username, password)
		code = err.SUCCESS
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": err.GetErrMsg(code),
	})

}

func GetSignup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.tmpl", gin.H{})
}