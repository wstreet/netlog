package main

import (
	"netlog/config"
	"netlog/http/router"

	"github.com/gin-gonic/gin"
)

func main() {
	c := config.GetConfig()
	gin.SetMode(c.Server.Mode)

	r := router.NewRouter()

	r.Run(c.Server.Port)

}
