
/**
 * @author: sarail
 * @time: 2021/6/13 17:21
**/

package model

import (
	"errors"
	"gorm.io/gorm"
	"log"
)

var (
	ErrInvalidAuthority = errors.New("invalid authority")
)

func NewAuthority() *Authority {
	return &Authority{}
}

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

// CanView UserID required
// return true if the user could view code with public auth field
func (a *Authority) CanView(tx *gorm.DB) bool {
	if err := tx.Where(a).First(a); err != nil {
		return false
	}
	return a.Admin || a.Viewer
}

// CanEditContest UserID required
func (a *Authority) CanEditContest(tx *gorm.DB) bool {
	if err := tx.Where(a).First(a); err != nil {
		return false
	}
	return a.Admin || a.ContestManager
}

func (a *Authority) CanEditProblem(tx *gorm.DB) bool {
	if err := tx.Where(a).First(a); err != nil {
		return false
	}
	return a.Admin || a.ProblemManager
}

func (a *Authority) CanAddAuth(tx *gorm.DB) bool {
	if err := tx.Where(a).First(a); err != nil {
		return false
	}
	return a.Admin
}

func (a *Authority) BeforeCreate(tx *gorm.DB) error {
	if !NewUserWithID(a.UserID).Exist(tx) {
		return ErrInvalidAuthority
	}

	return nil
}

// Add UserID required
func (a *Authority) Add(tx *gorm.DB, auth string) error {
	err := tx.Where(a).First(a).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return ErrInvalidAuthority
	}

	switch auth {
	case "viewer":
		a.Viewer = true
	case "admin":
		a.Admin = true
	case "contest":
		a.ContestManager = true
	case "problem":
		a.ProblemManager = true
	}

	return tx.Save(a).Error
}