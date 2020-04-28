package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{})
}

func PostLogin(c *gin.Context) {
	c.JSON(http.StatusOK, "已阅")
	log.Println(c.Query("username"), c.Query("password"))
}

