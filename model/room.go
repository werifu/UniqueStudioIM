package model

import "time"

type Room struct {
	Model

	Name 		string			`gorm:"size:16;unique;not null;unique" json:"name"`		// 改变默认长度（size）,非空, 唯一
	password	string
	CreatedBy	*User
	Hub			*Hub 			`gorm:"-"`
}


func NewRoom(hub *Hub, password, name string, creator *User) *Room {
	return &Room{
		Model:     Model{
			CreatedAt: time.Now(),
		},
		Name:      name,
		password:  password,
		CreatedBy: creator,
		Hub:       hub,
	}
}

func (room *Room) GetPsw() string{
	return room.password
}

func (room *Room) EditName(newName string) {
	room.Name = newName
}

func (room *Room) Delete() {
	room.Hub.Stop()
}
