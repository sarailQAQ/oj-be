
/**
 * @author: sarail
 * @time: 2021/6/18 16:32
**/

package model

import (
	"fmt"
	"gorm.io/gorm/clause"
	"oj-be/dao"
	"testing"
	"time"
)

func TestContest_Create(t *testing.T) {
	c := NewContest()
	c.Description = "Test"
	c.Tittle = "system contest test"
	c.UserID = 1
	c.BeginAt = time.Date(2021, 6, 18, 12, 0, 0, 0, time.Local)
	c.EndAt = time.Date(2021, 6, 18, 17, 0, 0, 0, time.Local)
	c.Type = contestPublic
	c.Participants = []User{
		//{
		//	Model: Model{ID: 1},
		//},
		//{
		//	Model: Model{ID: 2},
		//},
	}
	c.Problems = []Problem{
		//{
		//	Model: Model{ID: 2},
		//},
		//{
		//	Model: Model{ID: 3},
		//},
	}



	tx := dao.NewTX()
	tx = tx.Clauses(clause.OnConflict{DoNothing: true})
	//err := c.Create(tx)
	//log.Println(err)
	c.ID = 1

	c.AddUser(tx, 1)
	c.AddUser(tx, 2)
	//
	//err := c.AddProblem(tx, 3)
	//log.Println(err)
	//c.AddProblem(tx, 2)
	tx.Commit()
}

func TestContest_RankList(t *testing.T) {
	c := NewContest()
	c.ID = 1

	tx := dao.NewTX()
	fmt.Println(c.RankList(tx))
}
