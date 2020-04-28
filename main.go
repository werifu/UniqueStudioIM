package main

import (
	"im/routers"
	"log"
)

func main(){
	//加载路由器
	router := routers.InitRouter()

	err := router.Run(":8000")
	if err != nil{
		log.Println("engine start:", err)
	}

}
