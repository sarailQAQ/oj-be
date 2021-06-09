
/**
 * @author: sarail
 * @time: 2021/6/8 22:07
**/

package model

import (
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

	UserID  uint
	BeginAt time.Time
	EndAt   time.Time

	Tittle      string
	Description string

	Type uint8

	Problems []Problem `gorm:""`
}

func (c *Contest) AutoMigrate(tx *gorm.DB) {
	err := tx.AutoMigrate(&c)
	if err != nil {
		log.Panicln(err)
	}
}
