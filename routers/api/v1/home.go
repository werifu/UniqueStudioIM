package v1

import (
	"github.com/gin-gonic/gin"
	"im/model"
	"net/http"
)

//type Home struct {
//	Rooms map[string]*model.Room
//}


//获取首页房间等的信息
func GetHome(c *gin.Context) {
	var roomNames string
	for _, room := range model.Rooms {
		roomNames += room.Name
	}
	c.JSON(http.StatusOK, gin.H{
		"title": "首页",
		"rooms": roomNames,
	})
}

