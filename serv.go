package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hiyali/bilimger/lib"
)

const (
	port = 5555
)

func main() {
	lib.InitialEnvironment()
	if lib.GetEnv() == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.Default()
	RegisterRouter(engine)

	engine.Run(fmt.Sprintf("0.0.0.0:%d", port))
}
