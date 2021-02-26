package v1

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"thchat/pkg/e"
)

func GetStatus(c *gin.Context) {
	session := sessions.Default(c)
	loginUser := session.Get("loginUser")

	if loginUser == nil {
		c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": "未登陆", "islogin": 0})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": "已登陆", "islogin": 1, "username": loginUser.(string)})
	return
}
