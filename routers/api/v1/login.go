package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"im/model"
	"im/pkg/e"
	"im/pkg/logging"
	"log"
	"net/http"
)

func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{"title":"登录"})
}

func PostLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	logging.Info("登录：username:",username, ";password:", password)

	valid := validation.Validation{}
	valid.Required(username, "name").Message("名称不能为空")
	valid.MaxSize(username, 20, "name").Message("名称最长20字符")
	valid.Required(password, "password").Message("密码不能为空")
	valid.MinSize(password,5,  "password").Message("密码长度在5到40个字符长度区间")
	valid.MaxSize(password, 40, "password").Message("密码长度在5到40个字符长度区间")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "输入不符合规范"})
	}

	checkCode := model.LoginCheck(username, password)
	if checkCode != e.SUCCESS {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "用户名/密码错误"})
	} else {

		session := sessions.Default(c)
		session.Set("loginUser", username)
		err := session.Save()

		if err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "登录成功"})
	}


}

