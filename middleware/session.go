package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thchat/pkg/e"
	"thchat/pkg/util"
)





func LoginValid(handle gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if util.IsLogin(c) {
			handle(c)
		} else {
			c.JSON(http.StatusForbidden, gin.H{"code": e.NotLogin, "message": e.MsgFlags[e.NotLogin]})
		}
	}
}
