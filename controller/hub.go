package main

type Hub struct {
	//记录在场的客户端
	clients map[*Client]bool

	//广播版
	broadcast chan []byte

	//注册请求
	register chan *Client
}

func NewHub() *Hub{
	return &Hub{
		broadcast:  make(chan []byte),
		clients: 	make(map[*Client]bool),
		register:	make(chan *Client),
	}
}

func (h *Hub) run(){
	for{
		select {
		case conn := <-h.register:	//有新连接注册进来
			h.clients[conn] = true
		}
	}
}

