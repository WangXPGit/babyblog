package v1

import (
	"babyblog/model"
	"babyblog/utils/errmsg"

	"github.com/gin-gonic/gin"
	"net/http"
)

func UserExist(c *gin.Context) {
	// todo 查询用户是否存在
}

func AddUser(c *gin.Context) {
	// todo 添加用户

	// 获取前端参数
	// 检查参数中用户名是否存在
	// 存在则报错
	// 不存在则插入数据返回
	var data model.User
	_ = c.ShouldBindJSON(&data)
	var code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	// if code == errmsg.ERROR_USERNAME_USED {
	// 	code = errmsg.ERROR_USERNAME_USED
	// }
	c.JSON(http.StatusOK, gin.H{
		"status":code,
		"data":data,
		"message" : errmsg.GetErrMsg(code),
	})
}

func GetUser(c *gin.Context) {
	// todo 查询指定用户
}

func GetUsers(c *gin.Context) {
	// todo 查询用户列表
}

func EditUser(c *gin.Context) {
	// todo 编辑指定用户
}

func DeleteUser(c *gin.Context) {
	// todo 删除用户
}
