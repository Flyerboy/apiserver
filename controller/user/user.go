package user

import (
	"github.com/gin-gonic/gin"

	"apiserver/errormsg"
	"apiserver/model"
	. "apiserver/controller"

	"strconv"
)


func GetUser(c *gin.Context) {
	id := c.Param("id")
	userId, _ := strconv.Atoi(id)
	userModel := model.User{}
	user := userModel.GetUserById(userId)
	ResponseJson(c, errormsg.OK, user)
}

func GetLists(c *gin.Context) {
	userModel := model.User{}
	data := userModel.GetUserList(0, 10)
	ResponseJson(c, errormsg.OK, data)
}