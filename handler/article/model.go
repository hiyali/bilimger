package article

import (
	"github.com/jinzhu/gorm"
)

type Model struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
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
