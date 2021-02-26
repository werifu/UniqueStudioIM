package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thchat/model"
	"thchat/pkg/e"
	"thchat/pkg/logging"
	"thchat/pkg/util"
)


// @Summary 进入房间申请
// @Description 进入房间url时前端向后端发送密码申请许可
// @Tags v1
// @Produce json
// @Success 200 {string} json "{"code":200, "message":"可以进入"}"
// @Failure 403 {string} json "{"code":50001, "message":"未登录"}"
// @Failure 400 {string} json "{"code":10002, "message":"房间密码错误"}"
// @Failure 400 {string} json "{"code":10001, "message":"房间不存在"}"
// @Router /api/v1/room/{name} [get]
func GetRoom(c *gin.Context) {
	roomName := c.Param("name")
	psw := c.DefaultQuery("password", "")
	if _, ok := model.Rooms[roomName]; ok {

		if model.Rooms[roomName].GetPsw() == psw {
			c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": "ok"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"code": e.ErrRoomPassword, "message": "房间密码错误"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"code": e.ErrNotExistRoom, "message": e.MsgFlags[e.ErrNotExistRoom]})
	}
}


// post创建房间
func PostCreateRoom(c *gin.Context) {

	password := c.PostForm("password")
	roomName := c.PostForm("room_name")

	// 判断房间是否已经存在
	if _, ok := model.Rooms[roomName]; ok {
		c.JSON(http.StatusOK, gin.H{"code": e.ErrRoomExists, "message": e.MsgFlags[e.ErrRoomExists]})
		return
	}

	creatorName := util.GetSessionUsername(c)

	//让大水管跑起来
	var hub = model.NewHub()
	go hub.Run()


	var room = model.NewRoom(hub, password, roomName, &model.User{Username:creatorName})

	//注册到房间名单里
	model.Rooms[room.Name] = room
	logging.Info("房间创建成功：name:%s; psw:%s", room.Name, password)
	c.JSON(e.SUCCESS, "创建房间成功")
}

// 修改房间信息
func EditRoom(c *gin.Context) {

}

//删除房间
func DeleteRoom(c *gin.Context) {
	roomName := c.Param("name")

	username := util.GetSessionUsername(c)

	if room, ok := model.Rooms[roomName]; ok {
		if room.CreatedBy.Username == username {
			delete(model.Rooms, roomName)
			room.Delete()
			c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": "已删除"})
			return
		} else {
			c.JSON(http.StatusForbidden, gin.H{"code": e.ErrAuth, "message": "权限不足"})
			return
		}
	}
	c.JSON(http.StatusForbidden, gin.H{"code": e.ErrNotExistRoom, "message": "无该房间"})
}
