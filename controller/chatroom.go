package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ChatroomGet(c *gin.Context){
	c.HTML(http.StatusOK, "home.html", nil)
}
