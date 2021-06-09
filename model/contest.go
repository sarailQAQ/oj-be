
/**
 * @author: sarail
 * @time: 2021/6/8 22:07
**/

package model

import (
	"gorm.io/gorm"
	"time"
)

// 比赛状态常量
const (
	// 对所有人开放
	contestPublic = 1

	// 指定用户可参加
	contestGrant = 2

	// 仅自己
	contestPrivate = 3
)

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
