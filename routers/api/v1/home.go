package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//type Home struct {
//	Rooms map[string]*model.Room
//}


//获取首页房间等的信息
func GetHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "首页",
	})
}

