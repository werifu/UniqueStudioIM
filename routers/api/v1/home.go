package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"thchat/model"
)

//type Home struct {
//	Rooms map[string]*model.Room
//}
type RoomInfo struct {
	Name 		string
	CreatedBy 	string
	CreatedAt   string
}

//获取首页房间等的信息
func GetHome(c *gin.Context) {
	var rooms []RoomInfo
	for _, room := range model.Rooms {
		rooms = append(rooms, RoomInfo{
			Name:        room.Name,
			CreatedBy: room.CreatedBy.Username,
			CreatedAt: room.CreatedAt.Format("2006-01-02"),
		})
		fmt.Println("创建时间：", room.CreatedAt)
	}
	c.JSON(http.StatusOK, gin.H{
		"title": "首页",
		"rooms": rooms,
	})
}


