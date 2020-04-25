package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"im/model"
	"log"
	"net/http"
)

// 进入房间
func GetRoom(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}

// 创建房间
func CreateRoom(c *gin.Context) {
	password := c.Query("password")

	name := c.Query("name")

	//让大水管跑起来
	var hub = model.NewHub()
	go hub.Run()

	var room = model.NewRoom(hub, password, name, model.User{})

	//注册到房间名单里
	fmt.Println(room)
	model.Rooms[room.Name] = room
	log.Println("房间创建成功：", room.Name)

}

// 修改房间信息
func EditRoom(c *gin.Context) {

}

//删除房间
func DeleteRoom(c *gin.Context) {

}
