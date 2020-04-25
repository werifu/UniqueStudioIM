package routers

import (
	"github.com/gin-gonic/gin"
	"im/model"
	"im/pkg/setting"
	v1 "im/routers/api/v1"
)



func InitRouter() *gin.Engine{

	//创建路由
	router := gin.Default()
	gin.SetMode(setting.RunMode)

	router.LoadHTMLGlob("html/*")

	apiv1 := router.Group("/api/v1")
	{
		apiv1.GET("/", v1.GetHome)

		//进入房间
		apiv1.GET("/room/:name", v1.GetRoom)
		//新建房间
		apiv1.POST("/room/", v1.CreateRoom)
		//更新房间信息
		apiv1.PUT("/room/:name", v1.EditRoom)
		//删除指定房间
		apiv1.DELETE("room/:name", v1.DeleteRoom)
		//控制协议升级
		router.GET("room/:name/ws", func(c *gin.Context){
			model.SearchRoomWS(c)
		})
	}

	return router
}