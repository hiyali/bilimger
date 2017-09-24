package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hiyali/bilimger/handler"
)

var ApiPath = "/api"

func RegisterRouter(router *gin.Engine) {
	group := router.Group(ApiPath + "/user")
	restful := handler.GetRestful()
	restful.Apply(group)
}
