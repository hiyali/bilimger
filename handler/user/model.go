package user

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Model struct {
	gorm.Model
	Name          string
	Password      string
	Email         string
	RememberToken string
	LevelId       string
}

func (Model) TableName() string {
	fmt.Println("----------------------------------sadfasdfasdfa")
	return Name
}

/*
func (u User) TableName() string {
	if u.Role == "admin" {
		return "admin_users"
	} else {
		return "users"
	}
}
*/
