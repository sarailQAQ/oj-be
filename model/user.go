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
	ErrInvalidUser     = errors.New("invalid user")
	ErrUsernameExit    = errors.New("username exit")
	ErrMailUsed        = errors.New("mail has been used")
	ErrPhoneNumberUsed = errors.New("phone number has been used")
)

func NewUser() *User {
	return &User{}
}

func NewUserWithID(id uint) *User {
	return &User{
		Model: Model{ID: id},
	}
}

type User struct {
	Model

	Username    string `gorm:"index;type:varchar(30)"`
	Mail        string `gorm:"index;type:varchar(30)"`
	PhoneNumber string `gorm:"index;type:varchar(30)"`
	Password    string

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
	return u.ID == 0 && u.Username == "" && u.Mail == "" && u.PhoneNumber == ""
}

func (u *User) check(tx *gorm.DB) error {
	if u.Invalid() {
		return ErrInvalidUser
	}

	// team前缀用作比赛账号
	if strings.HasPrefix(u.Username, "team") {
		return ErrInvalidUser
	}

	txUser, cnt := tx.Model(NewUser()), int64(0)
	if len(u.Username) > 0 {
		txUser.Where("username=?", u.Username).Count(&cnt)
		if cnt > 0 {
			return ErrUsernameExit
		}
	}
	if len(u.PhoneNumber) > 0 {
		txUser.Where("phone_number=?", u.PhoneNumber).Count(&cnt)
		if cnt > 0 {
			return ErrPhoneNumberUsed
		}
	}
	if len(u.Mail) > 0 {
		txUser.Where("mail=?", u.Mail).Count(&cnt)
		if cnt > 0 {
			return ErrMailUsed
		}
	}

	return nil
}

// BeforeCreate 判断是否合法
func (u *User) BeforeCreate(tx *gorm.DB) error {
	return u.check(tx)
}

// Login Username OR Mail OR PhoneNumber AND Password required
func (u *User) Login(tx *gorm.DB) bool {
	if u.Invalid() {
		return false
	}

	tx.Select("id").Where(u).Find(&u)
	return u.ID > 0
}

// Register Username And Mail OR PhoneNumber AND Password required
func (u *User) Register(tx *gorm.DB) error {
	return tx.Create(u).Error
}

// Exist ID OR Username OR Mail OR PhoneNumber required
func (u *User) Exist(tx *gorm.DB) bool {
	cnt := int64(0)
	tx.Model(u).Where(u).Count(&cnt)
	return cnt > 0
}
