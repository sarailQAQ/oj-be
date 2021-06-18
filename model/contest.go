/**
 * @author: sarail
 * @time: 2021/6/8 22:07
**/

package model

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"sort"
	"time"
)

// 比赛类型常量
const (
	// 对所有人开放
	contestPublic = 1

	// 指定用户可参加
	contestGrant = 2

	// 仅自己
	contestPrivate = 3
)

var (
	ErrInvalidContest      = errors.New("invalid contest")
	ErrInvalidContestQuery = errors.New("invalid contest query")
)

func NewContest() *Contest {
	return &Contest{}
}

func NewContestDefault() *Contest {
	return &Contest{
		Model:       gorm.Model{},
		UserID:      0,
		BeginAt:     time.Time{},
		EndAt:       time.Time{},
		Tittle:      "",
		Description: "",
		Type:        0,
		Problems:    nil,
	}
}

// Contest 比赛实体
type Contest struct {
	gorm.Model

	// 创建者
	UserID uint

	BeginAt time.Time `gorm:"index:idx_begin"`
	EndAt   time.Time

	// 标题和描述
	Tittle      string
	Description string

	Type uint8

	// 关联模式
	Problems     []Problem `gorm:"many2many:contest_problems"`
	Participants []User    `gorm:"many2many:contest_users"`
}

func (c *Contest) AutoMigrate(tx *gorm.DB) {
	err := tx.AutoMigrate(&c)
	if err != nil {
		log.Panicln(err)
	}
}

func (c *Contest) BeforeCreate(tx *gorm.DB) error {
	if c.BeginAt.Before(time.Now()) {
		return ErrInvalidContest
	}

	if !NewUserWithID(c.UserID).Exist(tx) {
		return ErrInvalidContest
	}

	return nil
}

func (c *Contest) Create(tx *gorm.DB) error {
	return tx.Create(c).Error
}

// GetInfo ID required
func (c *Contest) GetInfo(tx *gorm.DB) error {
	return tx.Where(c).First(c).Error
}

func (c *Contest) AddProblem(tx *gorm.DB, problemID uint) error {
	return tx.Model(c).Association("Problem").Append(&Problem{
		Model: Model{ID: problemID},
	})
}

func (c *Contest) DelProblem(tx *gorm.DB, problemID uint) error {
	return tx.Model(c).Association("Problem").Delete(&Problem{
		Model: Model{ID: problemID},
	})
}

func (c *Contest) AddUser(tx *gorm.DB, userID uint) error {
	return tx.Model(c).Association("User").Append(&User{
		Model: Model{ID: userID},
	})
}

func (c *Contest) DelUser(tx *gorm.DB, userID uint) error {
	return tx.Model(c).Association("User").Delete(&User{
		Model: Model{ID: userID},
	})
}

type Rank struct {
	UserID uint

	// Score 过题数
	Score int64

	// Penalty 罚时，单位为秒
	Penalty int
}

type RankList []Rank

func (rl RankList) Len() int {
	return rl.Len()
}

func (rl RankList) Less(i, j int) bool {
	if rl[i].Score != rl[j].Score {
		return rl[i].Score > rl[j].Score
	}
	return rl[i].Penalty < rl[j].Penalty
}

func (rl RankList) Swap(i, j int) {
	rl[i], rl[j] = rl[j], rl[i]
}

func (c *Contest) RankList(tx *gorm.DB) (rankList RankList, err error) {
	err = tx.Find(c).Error
	if err != nil {
		return nil, ErrInvalidContestQuery
	}

	submissionModel := tx.Model(&Submission{}).Where("result = ?", AC)
	scoreQuery := submissionModel.Select("count(*)")
	penaltyQuery := submissionModel.Select("sum(submitted_at)")
	for u := range c.Participants {
		var rank Rank
		err = tx.Select("user_id, (?) as score, (?) as penalty",
			scoreQuery, penaltyQuery).Where("user_id=?", u).Scan(&rank).Error
		if err != nil {
			return nil, err
		}

		rankList = append(rankList, rank)
	}

	sort.Sort(rankList)
	return rankList, nil
}
