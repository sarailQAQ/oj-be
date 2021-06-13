
/**
 * @author: sarail
 * @time: 2021/6/13 17:21
**/

package model

import (
	"gorm.io/gorm"
	"log"
)

type Authority struct {
	gorm.Model

	UserID uint `gorm:"index:id_user_id"`

	// Admin 管理员
	// 管理员拥有所有权限
	Admin bool

	// Viewer
	// 拥有查看他人代码权限
	Viewer bool

	// ContestManager
	// 拥有添加/删除所有类型的比赛的权限
	// 同时可以修改所有比赛的信息
	ContestManager bool

	// ProblemManager
	// 拥有添加/删除所有类型的题目的权限
	// 同时可以修改所有题目的信息
	ProblemManager bool
}

func (a *Authority) AutoMigrate(tx *gorm.DB)  {
	err := tx.AutoMigrate(a)
	if err != nil {
		log.Fatalln(err)
	}
}