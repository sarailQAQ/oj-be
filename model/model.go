
/**
 * @author: sarail
 * @time: 2021/6/8 22:07
**/

package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type DeletedAt sql.NullTime

// Model 重新定义的model，为id添加了索引
type Model struct {
	ID        uint `gorm:"primaryKey;index:idx_model"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index:idx_deleted_at"`
}

// Init 显示调用
func Init(db *gorm.DB) {
	NewUser().AutoMigrate(db)
	NewContestProblem().AutoMigrate(db)
	NewContest().AutoMigrate(db)
	NewProblem().AutoMigrate(db)
	NewSubmission().AutoMigrate(db)
}
