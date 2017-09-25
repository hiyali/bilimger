package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hiyali/bilimger/handler/article"
	"github.com/hiyali/bilimger/handler/user"

	"net/http"
)

const (
	ApiPath = "/api"
)

func RegisterRouter(engine *gin.Engine) {
	user.RegisterRouter(engine, ApiPath)
	article.RegisterRouter(engine, ApiPath)

	// other

	// root ping
	engine.GET(ApiPath+"/salam", RootPing)
}

func RootPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": 123})
}
