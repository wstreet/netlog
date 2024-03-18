package router

import (
	"netlog/pkg/item"

	"github.com/gin-gonic/gin"
)

func RegisterParseRouter(g *gin.RouterGroup) {
	g.POST("items", item.SaveItems)
	// g.GET("parse/:id", parse.GetItems)
	g.GET("items", item.GetItems)
}
