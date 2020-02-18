package controller

type Hub struct {
	//记录在场的客户端
	clients map[*Client]bool

	//广播版
	broadcast chan []byte

	//注册请求
	register chan *Client

	//注销请求
	unregister chan *Client
}

func NewHub() *Hub{
	return &Hub{
		broadcast:  make(chan []byte),
		clients: 	make(map[*Client]bool),
		register:	make(chan *Client),
		unregister:	make(chan *Client),
	}
}

func (h *Hub) Run(){
	for{
		select {
		case c := <-h.register:	//有新连接注册进来
			h.clients[c] = true
		case c := <-h.unregister:
			if _, ok := h.clients[c];ok{		//用户表里确实有注销用户
				delete(h.clients, c)
				close(c.send)
			}
		case msg := <-h.broadcast:	//取出广播板消息
			for c, _ := range h.clients {
				c.send <- msg		//把消息注入各个客户端的连接里
			}
		}
	}
}

