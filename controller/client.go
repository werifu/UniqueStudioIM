package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"time"
)

const (
	//客户端（pong）等待ping的最长时间
	pongWait = 60 * time.Second

	//hub发ping频率
	pingPeriod = 50 * time.Second
)

//声明一个upgrader（把http升级成ws（用于创建ws.conn
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	hub *Hub
	conn *websocket.Conn
	send chan []byte
}

func (c *Client)SetPong(){

	//超过该时间就断开连接
	err := c.conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil{
		log.Println("SetDDL:",err)
	}

	//听ping（不用发pong
	c.conn.SetPongHandler(func(string)error{
		err := c.conn.SetReadDeadline(time.Now().Add(pongWait))
		if err != nil{ return err}
		return nil
	})
}

//把消息泵到hub里
func (c *Client)PumpToHub(){
	defer func(){
		c.hub.unregister <- c	//注销
		c.conn.Close()
	}()
	c.SetPong()
	for{
		_, msg, err := c.conn.ReadMessage()
		if err != nil{
			log.Println("One connection closed.")
			break
		}
		log.Println(string(msg))
		c.hub.broadcast <- msg

	}
}

//从hub写到连接里
func (c *Client)ReadFromHub(){
	ticker := time.NewTicker(pingPeriod)
	for{
		select {

			case msg,ok := <-c.send:
				if !ok{
				//说明连接已经关了
				return
			}
				//写入conn的writer
				w, err := c.conn.NextWriter(websocket.TextMessage)
				if err != nil{
					log.Println("next writer:",err)
				}
				//过滤输入
				template.HTMLEscape(w, msg)
				err = w.Close()
				if err != nil{
					log.Println("w close:",err)
					}

			//保持心跳(hub发ping到客户端
			case <-ticker.C:
				err := c.conn.WriteMessage(websocket.PingMessage, nil)
				if err != nil{
					log.Println("tick err:", err)
					}
		}
	}
}


func ServeWS(hub *Hub, c *gin.Context){	//开启服务
	//创建连接
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil{
		fmt.Println("upgrade err:", err)
		return
	}
	//创建一个客户端

	client := &Client{hub:hub, conn:conn, send: make(chan []byte, 256)}
	client.hub.register<-client

	go client.PumpToHub()
	go client.ReadFromHub()
}