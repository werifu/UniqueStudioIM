package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"thchat/model"
	"thchat/pkg/e"
	"thchat/pkg/logging"
	"thchat/pkg/util"
)

func PostLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	logging.Info("登录：username:",username, ";password:", password)

	valid := validation.Validation{}
	valid.Required(username, "username").Message("名称不能为空")
	valid.MaxSize(username, 20, "username").Message("名称最长20字符")
	valid.Required(password, "password").Message("密码不能为空")
	valid.MinSize(password,5,  "password").Message("密码长度在5到40个字符长度区间")
	valid.MaxSize(password, 40, "password").Message("密码长度在5到40个字符长度区间")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			logging.Error(err.Key, err.Message)
		}
		c.JSON(http.StatusOK, gin.H{"code": e.ErrFormat, "message": "输入不符合规范"})
		return
	}

	checkCode := model.LoginCheck(username, password)
	if checkCode != e.SUCCESS {
		c.JSON(http.StatusOK, gin.H{"code": e.ErrAuth, "message": "用户名/密码错误"})
		return
	}
	err := util.SetSession(c, username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": "无法设置session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": "登录成功", "username": username})


}

