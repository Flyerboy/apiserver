package controller

import (
	"apiserver/errormsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func ResponseJson(c *gin.Context, e *errormsg.ErrorMsg, data interface{}) {
	code, msg := errormsg.DecodeError(e)
	response := &ResponseData {
		Code: code,
		Message: msg,
		Data: data,
	}
	c.JSON(http.StatusOK, response)
}