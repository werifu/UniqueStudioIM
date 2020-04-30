package util

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

func IsLogin(c *gin.Context) bool {
	session := sessions.Default(c)
	loginUser := session.Get("loginUser")

	log.Println("session[登陆者]:", loginUser)
	if loginUser == nil {
		return false
	}
	return true
}

func GetSessionUsername(c *gin.Context) string {
	session := sessions.Default(c)
	loginUser := session.Get("loginUser")

	if loginUser == nil {
		return ""
	}
	return loginUser.(string)
}



