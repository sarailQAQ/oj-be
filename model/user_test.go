
/**
 * @author: sarail
 * @time: 2021/6/16 20:25
**/

package model

import (
	"oj-be/dao"
	"testing"
)

func TestUserLogin(t *testing.T) {
	user := NewUser()
	user.Username = "Hapwitch"
	user.Mail = "sarail@qq.com"
	user.Password = "123456"
	user.Sex = 1

	tx := dao.NewTX()
	//_ = user.Register(tx)

	u := NewUser()
	u.Username = "sarail"
	u.Password = "123456"
	_ = u.Login(tx)
	tx.Commit()

}

func TestUser_Register(t *testing.T) {
	user := NewUser()
	user.Username = "Hapwitch"
	user.Mail = "sarail@qq.com"
	user.Password = "123456"
	user.Sex = 1

	tx := dao.NewTX()
	_ = user.Register(tx)


	tx.Commit()
}
