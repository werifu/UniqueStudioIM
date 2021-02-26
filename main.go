package main

import (
	"log"
	"thchat/routers"
)

func main(){
	//加载路由器
	router := routers.InitRouter()
	err := router.Run(":8000")
	if err != nil{
		log.Println("engine star	t:", err)
	}
}
