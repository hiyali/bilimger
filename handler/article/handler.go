package article

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hiyali/bilimger/lib"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

func Create(c *gin.Context) {
	item := Model{
		Name: c.PostForm("name"),
	}

	db := lib.GetTable(Name)
	db.Save(&item)

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "id": item.ID})
}

func AllOrPaginate(c *gin.Context) {
	fmt.Println(Name + "handler")

	var _items []Model
	var items []Model

	sort := c.Query("sort") // age desc, name   ===    age desc, name asc

	db := lib.GetTable(Name)
	db = db.Order(sort)

	if c.Query("all") == "" {
		sizeStr := c.Query("size")
		size, sizeOk := strconv.Atoi(sizeStr)

		if sizeOk != nil {
			size = 10
			if sizeStr != "" {
				fmt.Println(Name + "AllOrPaginate handler size not a integer")
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

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": fmt.Sprintf("Ping to [%v]", Name)})
}

func Show(c *gin.Context) {
	var item gorm.Model
	itemId := c.Param("id") // :id

	if lib.GetEnv() == "test" && itemId == "ping" {
		Ping(c)
		return
	}

	db := lib.GetTable(Name)
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

	db := lib.GetTable(Name)
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

	db := lib.GetTable(Name)
	db.First(&item, itemId)

	if item.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	db.Delete(&item)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
}
