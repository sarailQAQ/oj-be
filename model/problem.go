/**
 * @author: sarail
 * @time: 2021/6/8 22:07
**/

package model

import (
	"errors"
	"gorm.io/gorm"
	"log"
)

// 问题的状态常量
const (
	// 正常状态，所有用户都可见
	problemNormal = 1

	// 仅自己可见
	problemOwnerOnly = 2

	// 比赛题目，比赛开始前
	problemBeforeContest = 3

	// 比赛题目，且比赛正在进行中
	problemInContest = 4
)

var (
	ErrInvalidProblem = errors.New("invalid problem")
)

func NewProblem() *Problem {
	return &Problem{}
}

func NewProblemWithID(id uint) *Problem {
	return &Problem{
		Model: Model{ID: id},
	}
}

type Problem struct {
	Model

	// UserID 上传题目的用户的ID
	UserID    uint

	// status: 题目的状态
	Status uint8

	// 标题、题目描述、输入描述、输出描述等
	Tittle            string `gorm:"type:varchar(255) not null"`
	Description       string
	InputDescription  string
	OutputDescription string
	Tip               string
}

func (p *Problem) AutoMigrate(tx *gorm.DB) {
	err := tx.AutoMigrate(&p)
	if err != nil {
		log.Panicln(err)
	}
}

func (p *Problem) BeforeCreate(tx *gorm.DB) error {
	if p.Status == 0 {
		p.Status = problemOwnerOnly
	}

	if err := p.Invalid(); err != nil {
		return err
	}
	if !NewUserWithID(p.UserID).Exist(tx) {
		return ErrInvalidProblem
	}

	return nil
}

func (p *Problem) Invalid() error {
	if p.UserID == 0 || p.Status == 0 || p.Status > problemInContest {
		return ErrInvalidProblem
	}

	return nil
}

func (p *Problem) Create(tx *gorm.DB) error {
	return tx.Create(p).Error
}

// FindOne ID required
func (p *Problem) FindOne(tx *gorm.DB) error {
	return tx.Where("id=?", p.ID).Find(p).Error
}

// Exist ID required
func (p *Problem) Exist(tx *gorm.DB) bool {
	cnt := int64(0)
	if tx.Where("id=?", p.ID).Count(&cnt); cnt > 0 {
		return true
	}

	return false
}
