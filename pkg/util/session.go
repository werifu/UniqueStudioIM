package util

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"thchat/pkg/logging"
)


func IsLogin(c *gin.Context) bool {
	session := sessions.Default(c)
	loginUser := session.Get("loginUser")

	logging.Info("session[登陆者]:", loginUser)
	if loginUser == nil {
		return false
	}
	return true
}

func GetSessionUsername(c *gin.Context) string {
	session := sessions.Default(c)
	loginUser := session.Get("loginUser")

	if loginUser == nil {
		fmt.Println("未登录")
		return ""
	}
	return loginUser.(string)
}



func SetSession(c *gin.Context, username string) error {
	session := sessions.Default(c)
	session.Options(sessions.Options{
		Path:     "/",
		Domain:   "",
		MaxAge:   3600,
		Secure: false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	session.Set("loginUser", username)

	err := session.Save()

	if err != nil {
		logging.Error(err)
		return err
	}
	return nil
}