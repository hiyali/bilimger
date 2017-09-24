package article

import (
	"github.com/gin-gonic/gin"
	"github.com/hiyali/bilimger/lib"
	"strings"
)

/*
func group(c *gin.Context) {
	//Migrate the schema
	db := getDB()
	db.AutoMigrate(&Todo{})
}
*/
const (
	Name = "article"
)

func RegisterRouter(engine *gin.Engine, apiPath string) {
	group := engine.Group(strings.Join([]string{apiPath, Name}, "/"))
	// group.POST("/", r.Create)
	// group.GET("/", r.AllOrPaginate)
	// group.GET("/:id", r.Show)
	// group.PUT("/:id", r.Update)
	// group.DELETE("/:id", r.Delete)
	restful := lib.Restful{Create, AllOrPaginate, Show, Update, Delete}
	restful.Apply(group)

	// other
}
