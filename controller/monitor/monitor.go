package monitor

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func Ping(c *gin.Context) {
	c.Status(http.StatusOK)
}