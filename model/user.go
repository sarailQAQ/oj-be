
/**
 * @author: sarail
 * @time: 2021/6/8 22:07
**/

package model

import (
	"gorm.io/gorm"
	"log"
)

func NewUser() *User {
	return &User{}
}
type User struct {
	model

	Username    string
	Mail        string `gorm:"uniqueIndex"`
	PhoneNumber string `gorm:"uniqueIndex"`

	Nickname string
	Sex      int8
	Avatar   string
}

func (u *User) AutoMigrate(tx *gorm.DB) {
	err := tx.AutoMigrate(&u)
	if err != nil {
		log.Panicln(err)
	}
}


