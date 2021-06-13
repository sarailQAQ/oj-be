/**
 * @author: sarail
 * @time: 2021/6/8 22:07
**/

package model

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"time"
)

// 比赛类型常量
const (
	// 对所有人开放
	contestPublic = 1

	// 指定用户可参加
	contestGrant = 2

	// 仅自己
	contestPrivate = 3
)

var (
	ErrInvalidContest = errors.New("invalid contest")
)

func NewContest() *Contest {
	return &Contest{}
}

func NewContestDefault() *Contest {
	return &Contest{
		Model:       gorm.Model{},
		UserID:      0,
		BeginAt:     time.Time{},
		EndAt:       time.Time{},
		Tittle:      "",
		Description: "",
		Type:        0,
		Problems:    nil,
	}
}

// Contest 比赛实体
type Contest struct {
	gorm.Model

	// 创建者
	UserID uint

	BeginAt time.Time `gorm:"index:idx_begin"`
	EndAt   time.Time

	// 标题和描述
	Tittle      string
	Description string

	Type uint8

	Problems     []Problem `gorm:"many2many:contest_problems"`
	Participants []User    `gorm:"many2many:contest_users"`
}

func (c *Contest) AutoMigrate(tx *gorm.DB) {
	err := tx.AutoMigrate(&c)
	if err != nil {
		log.Panicln(err)
	}
}

func (c *Contest) BeforeCreate(tx *gorm.DB) error {
	if c.BeginAt.Before(time.Now()) {
		return ErrInvalidContest
	}

	if !NewUserWithID(c.UserID).Exist(tx) {
		return ErrInvalidContest
	}

	return nil
}

func (c *Contest) Create(tx *gorm.DB) error {
	return tx.Create(c).Error
}

func (c *Contest) AddProblem(tx *gorm.DB, problemID uint) error {
	return tx.Model(c).Association("Problem").Append(&Problem{
		Model: Model{ID: problemID},
	})
}

func (c *Contest) DelProblem(tx *gorm.DB, problemID uint) error {
	return tx.Model(c).Association("Problem").Delete(&Problem{
		Model: Model{ID: problemID},
	})
}

func (c *Contest) AddUser(tx *gorm.DB, userID uint) error {
	return tx.Model(c).Association("User").Append(&User{
		Model: Model{ID: userID},
	})
}

func (c *Contest) DelUser(tx *gorm.DB, userID uint) error {
	return tx.Model(c).Association("User").Delete(&User{
		Model: Model{ID: userID},
	})
}

