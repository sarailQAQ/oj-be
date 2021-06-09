
/**
 * @author: sarail
 * @time: 2021/6/8 22:07
**/

package model

type User struct {
	model

	Username    string
	Mail        string `gorm:"uniqueIndex"`
	PhoneNumber string `gorm:"uniqueIndex"`

	Nickname string
	Sex      int8
	Avatar   string
}
