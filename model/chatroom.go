package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ChatroomGet(c *gin.Context){
	c.HTML(http.StatusOK, "room.tmpl", nil)
}
