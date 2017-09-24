package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hiyali/bilimger/lib"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

/*
func group(c *gin.Context) {
	//Migrate the schema
	db := getDB()
	db.AutoMigrate(&Todo{})
}
*/
func GetRestful() *lib.Restful {
	return &lib.Restful{Create, AllOrPaginate, Show, Update, Delete}
}

func Create(c *gin.Context) {
	item := User{
		Name: c.PostForm("name"),
	}

	db := lib.GetDB()
	db.Save(&item)

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "id": item.ID})
}

func AllOrPaginate(c *gin.Context) {
	var _items []User
	var items []User

	sort := c.Query("sort") // age desc, name   ===    age desc, name asc

	db := lib.GetDB()
	db = db.Order(sort)

	if c.Query("all") == "" {
		sizeStr := c.Query("size")
		size, sizeOk := strconv.Atoi(sizeStr)

		if sizeOk != nil {
			size = 10
			if sizeStr != "" {
				fmt.Println("User AllOrPaginate handler size not a integer")
			}
		}

		db.Limit(size).Find(&_items)
	} else {
		db.Find(&_items)
	}

	if len(_items) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	for _, item := range _items {
		items = append(items, item)
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": items})
}

func Show(c *gin.Context) {
	var item gorm.Model
	itemId := c.Param("id") // :id

	db := lib.GetDB()
	db.First(&item, itemId)

	if item.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": item})
}

func Update(c *gin.Context) {
	var item gorm.Model
	itemId := c.Param("id")

	db := lib.GetDB()
	db.First(&item, itemId)

	if item.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	db.Save(&item)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
}

func Delete(c *gin.Context) {
	var item gorm.Model
	itemId := c.Param("id")

	db := lib.GetDB()
	db.First(&item, itemId)

	if item.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	db.Delete(&item)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Pageable struct {
	Size int
	Sort string
}
