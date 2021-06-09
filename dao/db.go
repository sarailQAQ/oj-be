
/**
 * @author: sarail
 * @time: 2021/6/8 22:19
**/

package dao

import "gorm.io/gorm"

func NewTX() *gorm.DB {
	return db.Begin()
}


