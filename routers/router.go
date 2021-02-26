package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"thchat/middleware"
	"thchat/model"
	"thchat/pkg/config"
	v1 "thchat/routers/api/v1"
	"time"
)

func InitRouter() *gin.Engine{

	//创建路由
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	gin.SetMode(config.AppConfig.RunMode)


	// 设置session
	store := cookie.NewStore([]byte("loginUser"))
	router.Use(sessions.Sessions("mysession", store))


	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := router.Group("/api/v1")
	{
		apiv1.GET("/callback/github", v1.GithubOauthCallback)

		apiv1.GET("/", v1.GetHome)
		apiv1.GET("/home", v1.GetHome)

		//进入房间
		apiv1.GET("/room/:name", middleware.LoginValid(v1.GetRoom))

		//新建房间
		apiv1.POST("/newroom", middleware.LoginValid(v1.PostCreateRoom))

		//更新房间信息
		apiv1.PUT("/room/:name", v1.EditRoom)
		//删除指定房间
		apiv1.DELETE("/room/:name", v1.DeleteRoom)

		//登录注册
		apiv1.POST("/login", v1.PostLogin)
		apiv1.POST("/oauth/github", v1.OauthGithub)

		apiv1.GET("/signup", v1.GetSignup)
		apiv1.POST("/signup", v1.PostSignup)

		//状态
		apiv1.GET("/status", v1.GetStatus)

		//控制协议升级
		apiv1.GET("/room/:name/ws", func(c *gin.Context){
			model.SearchRoomWS(c)
		})
	}

	return router
}