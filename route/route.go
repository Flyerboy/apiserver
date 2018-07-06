package route

import (
	"github.com/gin-gonic/gin"
	"apiserver/controller/user"
	"apiserver/controller/monitor"
)

func Load(g *gin.Engine) *gin.Engine {
	r := g.Group("/users")
	{
		r.GET("", user.GetLists)
		r.GET("/:id", user.GetUser)
	}
	r = g.Group("monitor")
	{
		r.GET("ping", monitor.Ping)
	}
	return g
}