package routers

import (
	"github.com/gin-gonic/gin"
	"im/controller"

)

func InitRouter() *gin.Engine{

	//让大水管跑起来
	var hub = controller.NewHub()
	go hub.Run()

	//创建路由
	router := gin.Default()
	router.LoadHTMLGlob("html/*")
	//聊天室路由
	router.GET("/",controller.ChatroomGet)

	//控制协议升级
	router.GET("/ws", func(c *gin.Context){
		controller.ServeWS(hub, c)
	})

	return router
}