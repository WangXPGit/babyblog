package v1

import (
	"babyblog/model"
	"babyblog/utils/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
	"net/http"
)

func UserExist(c *gin.Context) {
	// todo 查询用户是否存在
}

func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	var code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		code = model.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetUser(c *gin.Context) {
	// todo 查询指定用户
}

func GetUsers(c *gin.Context) {
	pageSize,_ := strconv.Atoi(c.Query("pageSize"))
	pageNum,_ := strconv.Atoi(c.Query("pageNum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetUsers(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status" : code,
		"data" : data,
		"message" : errmsg.GetErrMsg(code),
	})
}

func EditUser(c *gin.Context) {
	// todo 编辑指定用户
}

func DeleteUser(c *gin.Context) {
	// todo 删除用户
}

// 密码加密
func scryptPw(password string) string {
	const PwHashByte = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 11}

	scrypt.
}