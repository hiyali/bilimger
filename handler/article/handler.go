package article

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hiyali/bilimger/lib"
	"net/http"
	"strconv"
)

func Create(c *gin.Context) {
	item := Model{
		Name: c.PostForm("name"),
	}

	db := lib.GetDB()
	defer db.Close()
	db.Save(&item)

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": fmt.Sprintf("%v item created successfully!", Name), "id": item.ID})
}

func AllOrPaginate(c *gin.Context) {
	fmt.Println(Name + "handler")

	var _items []Model
	var items []Model

	sort := c.Query("sort") // age desc, name   ===    age desc, name asc

	db := lib.GetDB()
	defer db.Close()
	db = db.Order(sort)

	if c.Query("all") == "" {
		sizeStr := c.Query("size")
		size, sizeOk := strconv.Atoi(sizeStr)

		if sizeOk != nil {
			size = 10
			if sizeStr != "" {
				fmt.Println(Name + " AllOrPaginate handler size not a integer")
			}
		}

		db.Limit(size).Find(&_items)
	} else {
		db.Find(&_items)
	}

	if len(_items) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": fmt.Sprintf("No %v found!", Name)})
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
	var item Model
	itemId := c.Param("id") // :id
	fmt.Printf("getting item in [%v] for id [%v] \n", Name, itemId)

	if lib.GetEnv() == "test" && itemId == "ping" {
		Ping(c)
		return
	}

	db := lib.GetDB()
	defer db.Close()
	db.First(&item, itemId)
	// db.Unscoped().Table("items").Where("id = " + itemId).Find(&item)
	// db.Unscoped().Table(Name).Where("id = " + itemId).Find(&item)
	// fmt.Printf("table [%v] is %v\n", Name, db.HasTable(Name))
	// db.Table(Name).Where(Model{ID: 1}).First(&item)

	if item.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": fmt.Sprintf("No %v found!", Name)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": item})
}

func Update(c *gin.Context) {
	var item Model
	itemId := c.Param("id")

	db := lib.GetDB()
	defer db.Close()
	db.First(&item, itemId)

	if item.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": fmt.Sprintf("No %v found!", Name)})
		return
	}

	db.Save(&item)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": fmt.Sprintf("%v updated successfully!", Name)})
}

func Delete(c *gin.Context) {
	var item Model
	itemId := c.Param("id")

	db := lib.GetDB()
	defer db.Close()
	db.First(&item, itemId)

	if item.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": fmt.Sprintf("No %v found!", Name)})
		return
	}

	db.Delete(&item)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": fmt.Sprintf("%v deleted successfully!", Name)})
}
