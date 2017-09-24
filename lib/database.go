package lib

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GetDB() *gorm.DB {
	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	if err == nil {
		panic("failed to connect database")
	}

	return db
}
