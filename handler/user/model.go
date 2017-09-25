package user

import (
	"github.com/jinzhu/gorm"
)

type Model struct {
	gorm.Model
	Name          string
	Password      string
	Email         string
	RememberToken string
	LevelId       int
}

func (Model) TableName() string {
	return Name
}

/*
func (u User) TableName() string {
	if u.Role == "admin" {
		return "admin_" + Name
	} else {
		return Name
	}
}
*/
