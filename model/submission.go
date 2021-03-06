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

var (
	ErrInvalidSubmission = errors.New("invalid submission")
)

const (
	AC  = 1
	WA  = 2
	TLE = 3
	RE  = 4
	MLE = 5
	PE  = 6
)

func NewSubmission() *Submission {
	return &Submission{}
}

type Submission struct {
	gorm.Model

	ProblemID uint `gorm:"not null;index:sub_problem,priority:1"`
	UserID    uint `gorm:"not null;index:sub_user,priority:1"`

	TimeUsed    float32 `gorm:"not null"`
	Language    string  `gorm:"type:varchar(10) not null"`
	Code        string  `gorm:"not null"`
	SubmittedAt int64   `gorm:"not null;index:sub_problem,priority:2;index:sub_user,priority:2"`

	// 结果
	Result int
}

func (sub *Submission) AutoMigrate(tx *gorm.DB) {
	err := tx.AutoMigrate(&sub)
	if err != nil {
		log.Panicln(err)
	}
}

func (sub *Submission) BeforeCreate(tx *gorm.DB) error {
	if !NewProblemWithID(sub.ProblemID).Exist(tx) {
		return ErrInvalidSubmission
	}
	if !NewUserWithID(sub.UserID).Exist(tx) {
		return ErrInvalidUser
	}

	return nil
}

func (sub *Submission) Create(tx *gorm.DB) error {
	return tx.Create(sub).Error
}
