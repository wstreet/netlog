package router

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// r.SetTrustedProxies([]string{"127.0.0.1"})
	v1 := r.Group("/api/v1")
	RegisterParseRouter(v1)
	return r
}
