package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
)

const (
	port = 5555
)

func main() {

	env := flag.String("env", "dev", "Environment")
	flag.Parse()

	if *env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.Default()

	RegisterRouter(engine)

	engine.Run(fmt.Sprintf("0.0.0.0:%d", port))
}
