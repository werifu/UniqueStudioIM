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

	router.LoadHTMLGlob("tmpl/*")
	router.Static("/static", "./static")
	apiv1 := router.Group("/api/v1")
	{
		apiv1.GET("/", v1.GetHome)
		apiv1.GET("/newroom", v1.NewRoomAPI)
		//进入房间
		apiv1.GET("/room/:name", v1.GetRoom)
		//新建房间
		apiv1.POST("/room/", v1.CreateRoom)
		//更新房间信息
		apiv1.PUT("/room/:name", v1.EditRoom)
		//删除指定房间
		apiv1.DELETE("/room/:name", v1.DeleteRoom)

		//登录注册
		apiv1.GET("/login", v1.GetLogin)
		apiv1.POST("/login/", v1.PostLogin)
		apiv1.GET("/signup", v1.GetSignup)
		apiv1.POST("/signup", v1.PostSignup)

		//控制协议升级
		router.GET("room/:name/ws", func(c *gin.Context){
			model.SearchRoomWS(c)
		})
	}

	return router
}