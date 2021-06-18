
/**
 * @author: sarail
 * @time: 2021/6/16 21:05
**/

package model

import (
	"oj-be/dao"
	"testing"
)

func TestProblem_Create(t *testing.T) {
	p := NewProblem()
	p.Status = problemNormal
	p.Tittle = "quick sort"
	p.Description = "This is a quick sort problem"
	p.InputDescription = "The first line is a integer n.\nT" +
		"he next line contains n integer split with a space."
	p.OutputDescription = "One line contains n integers."
	p.Tip = "Is there anyone who can`t write a quick sort? Bu hui ba?"
	p.UserID = 1

	tx := dao.NewTX()
	_ = p.Create(tx)
}

func TestProblem_FindOne(t *testing.T) {
	p := NewProblemWithID(1)
	tx := dao.NewTX()
	_ = p.FindOne(tx)
}
