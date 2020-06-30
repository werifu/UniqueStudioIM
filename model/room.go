package model

type Room struct {
	Model

	Name 		string			`gorm:"size:16;unique;not null;unique" json:"name"`		// 改变默认长度（size）,非空, 唯一
	password	string
	CreatedBy	*User
	Hub			*Hub 			`gorm:"-"`
}

//获得各房间数据
//注:函数尾部已经 声明了返回变量，所以return可以不带变量
func GetRooms(maps interface{}) (rooms []Room) {
	db.Where(maps).Find(&rooms)
	return
}

//获得房间数量
func GetRoomTotal(maps interface{}) (count int) {
	db.Model(&Room{}).Where(maps).Count(&count)
	return
}

func NewRoom(hub *Hub, password, name string, creator *User) *Room {
	return &Room{
		Model:     Model{},
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
