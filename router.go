package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hiyali/bilimger/handler/article"
	"github.com/hiyali/bilimger/handler/user"
)

const (
	ApiPath = "/api"
)

func RegisterRouter(router *gin.Engine) {
	// user
	user.RegisterRouter(router, ApiPath)

	// article
	article.RegisterRouter(router, ApiPath)

	// other
}
