package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var DB = make(map[string]string)

func formatAsTime(t time.Time) string {
	hour, minute, second := t.Clock()
	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
}
func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.Delims("{[{", "}]}")
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
		"formatAsTime": formatAsTime,
	})
	r.LoadHTMLFiles("assets/templates/index.tmpl")
	// router.LoadHTMLGlob("assets/templates/*")

	r.GET("/sur", func(c *gin.Context) {
		t := time.Now()
		c.HTML(http.StatusOK, "index.tmpl", map[string]interface{}{
			"now":  t,
			"year": t.Year(),
		})
	})

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := DB[user]
		if ok {
			c.JSON(200, gin.H{"user": user, "value": value})
		} else {
			c.JSON(200, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.GET("admin", func(c *gin.Context) {
		// user := c.Params('user') // MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		// var json struct {
		// 	Value string `json:"value" binding:"required"`
		// }
		type Account struct {
			Account  string `form:"account"`
			Password string `form:"password"`
		}
		var account Account

		if c.Bind(&account) == nil {
			DB[account.Account] = account.Password
			c.JSON(200, gin.H{"status": "ok"})
		} else {
			c.JSON(200, gin.H{"status": "not ok"})
		}
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
