package main

import (
	"github.com/gin-gonic/gin"

	"apiserver/route"
	"apiserver/config"
	"apiserver/tool"

	"log"
	"fmt"
)

func main() {

	err := config.Init()
	if err != nil {
		fmt.Printf("err:%s", err.Error())
		return
	}

	go func() {
		err := tool.PingServer()
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Printf("server running successfully")
	}()

	g := gin.New()
	r := route.Load(g)

	r.Run(config.Get("host") + ":" + config.Get("addr"))

}
