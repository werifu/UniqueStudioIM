package v1

import (
	"github.com/gin-gonic/gin"
	"im/model"
	"im/pkg/util"
	"log"
	"net/http"
)


// 进入房间
func GetRoom(c *gin.Context) {
	roomName := c.Param("name")
	if _, ok := model.Rooms[roomName]; ok {
		c.HTML(http.StatusOK, "room.tmpl", nil)
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "房间不存在"})
	}
}


func GetCreateRoom(c *gin.Context) {
	c.HTML(http.StatusOK, "newroom.tmpl", nil)
}

// post创建房间
func PostCreateRoom(c *gin.Context) {

	password := c.PostForm("password")
	roomName := c.PostForm("room_name")

	// 判断房间是否已经存在
	if _, ok := model.Rooms[roomName]; ok {
		c.JSON(http.StatusOK, gin.H{"message": "房间名已被占用"})
		return
	}

	//让大水管跑起来
	var hub = model.NewHub()
	go hub.Run()

	creatorName := util.GetSessionUsername(c)

	var room = model.NewRoom(hub, password, roomName, &model.User{Username:creatorName})

	//注册到房间名单里
	model.Rooms[room.Name] = room
	log.Printf("房间创建成功：name:%s; psw:%s", room.Name, password)
}

// 修改房间信息
func EditRoom(c *gin.Context) {

}

//删除房间
func DeleteRoom(c *gin.Context) {
	roomName := c.Param("name")

	username := util.GetSessionUsername(c)

	if room, ok := model.Rooms[roomName]; ok {
		if room.CreatedBy.Username == username {
			delete(model.Rooms, roomName)
			room.Delete()
		} else {
			c.JSON(http.StatusForbidden, gin.H{"message": "权限不足"})
		}
	}
}
