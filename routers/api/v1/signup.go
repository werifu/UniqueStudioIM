package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"thchat/model"
	"thchat/pkg/e"
	"thchat/pkg/logging"
)

func PostSignup(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	logging.Info("注册参数：", username, password)

	valid := validation.Validation{}
	valid.Required(username, "name").Message("名称不能为空")
	valid.MaxSize(username, 20, "name").Message("名称最长20字符")
	valid.Required(password, "password").Message("密码不能为空")
	//valid.Range(password,5, 40, "password").Message("密码长度在5到40个字符长度区间")

	if valid.HasErrors() {
		var msg []string
		for _, err := range valid.Errors {
			logging.Error(err.Key, err.Message)
			msg = append(msg, err.Message)
		}
		c.JSON(http.StatusExpectationFailed, gin.H{"msgs": msg})
	}


	var code int
	if model.UserExists(username) {
		log.Println(username)
		code = e.ErrUserExists
	} else {
		model.AddUser(username, password)
		code = e.SUCCESS
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetErrMsg(code),
	})

}

func GetSignup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.tmpl", gin.H{"title": "注册"})
}