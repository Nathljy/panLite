package api

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	route    *gin.Engine
	initOnce sync.Once
)

func NewRoute() *gin.Engine {
	initOnce.Do(func() {
		route = gin.Default()
		InitRoot()
		user := route.Group("/auth")
		InitAuth(user)
		file := route.Group("/file")
		InitFile(file)
	})
	return route
}

func InitRoot() {
	route.GET("/introduce", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "panLite"})
	})
}
