
/**
 * @author: sarail
 * @time: 2021/6/18 16:02
**/

package model

import (
	"oj-be/dao"
	"testing"
	"time"
)

func TestSubmission_Create(t *testing.T) {
	sub := NewSubmission()
	sub.ProblemID = 2
	sub.UserID = 1
	sub.SubmittedAt = time.Now().Unix()
	sub.Code = "print(hello world)"
	sub.Result = AC
	sub.Language = "Python"
	sub.TimeUsed = 0.03

	tx := dao.NewTX()
	_ = sub.Create(tx)
	tx.Commit()
}
