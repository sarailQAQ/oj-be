
/**
 * @author: sarail
 * @time: 2021/6/8 22:07
**/

package model

import (
	"gorm.io/gorm"
	"log"
)

func NewSubmission() *Submission {
	return &Submission{}
}

type Submission struct {
	gorm.Model

	ProblemID uint `gorm:"index:sub_problem,priority:1"`
	UserID    uint `gorm:"index:sub_user,priority:1"`

	TimeUsed    float32
	Language    string
	Code        string
	SubmittedAt int64 `gorm:"index:sub_problem,priority:2;index:sub_user,priority:2"`
}

func (sub *Submission) AutoMigrate(tx *gorm.DB) {
	err := tx.AutoMigrate(&sub)
	if err != nil {
		log.Panicln(err)
	}
}