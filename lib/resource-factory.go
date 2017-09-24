package lib

import (
	"github.com/gin-gonic/gin"
)

type Restful struct {
	Create        func(*gin.Context)
	AllOrPaginate func(*gin.Context)
	Show          func(*gin.Context)
	Update        func(*gin.Context)
	Delete        func(*gin.Context)
}

func (r *Restful) Apply(group *gin.RouterGroup) {
	group.POST("/", r.Create)
	group.GET("/", r.AllOrPaginate)
	group.GET("/:id", r.Show)
	group.PUT("/:id", r.Update)
	group.DELETE("/:id", r.Delete)
}
