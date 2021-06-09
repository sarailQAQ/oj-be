
/**
 * @author: sarail
 * @time: 2021/6/8 22:07
**/

package model

import (
	"gorm.io/gorm"
	"log"
)

func NewContestProblem() *ContestProblem {
	return &ContestProblem{}
}

// ContestProblem 一个比赛包含的题目
type ContestProblem struct {
	gorm.Model

	ContestID uint
	ProblemID uint
}

func (cp *ContestProblem) AutoMigrate(tx *gorm.DB) {
	err := tx.AutoMigrate(&cp)
	if err != nil {
		log.Panicln(err)
	}
}