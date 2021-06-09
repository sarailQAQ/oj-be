
/**
 * @author: sarail
 * @time: 2021/6/8 22:07
**/

package model

import (
	"database/sql"
	"time"
)

type DeletedAt sql.NullTime

// model 重新定义的model，为id添加了索引
type model struct {
	ID        uint `gorm:"primaryKey;index:model"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt DeletedAt `gorm:"index:deleted_at"`
}
