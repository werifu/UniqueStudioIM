package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)


func SearchRoomWS(c *gin.Context) {
	roomName := c.Param("name")
	room := Rooms[roomName]
	log.Println("rooms:", Rooms)
	ServeWS(room.Hub, c)
}


func ServeWS(hub *Hub, c *gin.Context){ //开启服务
	//创建连接
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil{
		fmt.Println("upgrade err:", err)
		return
	}

	//创建一个客户端
	client := &Client{hub: hub, conn:conn, send: make(chan []byte, 256)}
	log.Println("register <- client")
	client.hub.register <- client

	go client.PumpToHub()
	go client.ReadFromHub()
}