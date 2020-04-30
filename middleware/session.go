package middleware

import (
	"github.com/gin-gonic/gin"
	"im/pkg/util"
	"net/http"
)


func LoginValid(handle gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if util.IsLogin(c) {
			handle(c)
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 200, "message": "请先登录"})
		}
	}
}
