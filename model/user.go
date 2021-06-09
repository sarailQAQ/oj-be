
/**
 * @author: sarail
 * @time: 2021/6/8 22:07
**/

package model

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"strings"
)

var (
	InvalidUser     = errors.New("invalid user")
	UsernameExit    = errors.New("username exit")
	MailUsed        = errors.New("mail has been used")
	PhoneNumberUsed = errors.New("phone number has been used")
)

func NewUser() *User {
	return &User{}
}

type User struct {
	Model

	Username    string `gorm:"index;type:varchar(30)"`
	Mail        string `gorm:"index;type:varchar(30)"`
	PhoneNumber string `gorm:"index;type:varchar(30)"`

	Nickname string
	Sex      int8
	Avatar   string
}

func (u *User) AutoMigrate(tx *gorm.DB) {
	err := tx.AutoMigrate(u)
	if err != nil {
		log.Panic(err)
	}
}

func (u *User) Invalid() bool {
	return u.Username == "" && u.Mail == "" && u.PhoneNumber == ""
}

// BeforeCreate 判断是否合法
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Invalid() {
		return InvalidUser
	}

	// team前缀用作比赛账号
	if strings.HasPrefix(u.Username, "team") {
		return InvalidUser
	}

	txUser, cnt := tx.Model(NewUser()), int64(0)
	if len(u.Username) > 0 {
		txUser.Where("username=?", u.Username).Count(&cnt)
		if cnt > 0 {
			return UsernameExit
		}
	}
	if len(u.PhoneNumber) > 0 {
		txUser.Where("phone_number=?", u.PhoneNumber).Count(&cnt)
		if cnt > 0 {
			return PhoneNumberUsed
		}
	}
	if len(u.Mail) > 0 {
		txUser.Where("mail=?", u.Mail).Count(&cnt)
		if cnt > 0 {
			return MailUsed
		}
	}

	return nil
}


// Login Username OR Mail OR PhoneNumber AND Password needed
func (u *User) Login(tx *gorm.DB) (ok bool){
	if  u.Invalid() {
		return false
	}

	tx.Select("id").Where(u).Find(&u)
	return u.ID > 0
}

// Register Username And Mail OR PhoneNumber AND Password needed
func (u *User) Register(tx *gorm.DB) (err error) {
	return tx.Create(u).Error
}


