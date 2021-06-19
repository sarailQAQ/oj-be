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
	Problems     []Problem `gorm:"many2many:contest_problems;foreignKey:ID"`
	Participants []User    `gorm:"many2many:contest_users;foreignKey:ID"`
}

func (c *Contest) AutoMigrate(tx *gorm.DB) {
	err := tx.AutoMigrate(&c)
	if err != nil {
		log.Panicln(err)
	}
}

func (c *Contest) BeforeCreate(tx *gorm.DB) error {
	//if c.BeginAt.Before(time.Now()) {
	//	return ErrInvalidContest
	//}

	if !NewUserWithID(c.UserID).Exist(tx) {
		return ErrInvalidContest
	}

	return nil
}

func (c *Contest) Create(tx *gorm.DB) error {
	return tx.Save(c).Error
}

// GetInfo ID required
func (c *Contest) GetInfo(tx *gorm.DB) error {
	return tx.Where(c).First(c).Error
}

func (c *Contest) AddProblem(tx *gorm.DB, problemID uint) error {
	return tx.Model(c).Association("Problems").Append(&Problem{
		Model: Model{ID: problemID},
	})
}

func (c *Contest) DelProblem(tx *gorm.DB, problemID uint) error {
	return tx.Model(c).Association("Problems").Delete(&Problem{
		Model: Model{ID: problemID},
	})
}

func (c *Contest) AddUser(tx *gorm.DB, userID uint) error {
	return tx.Model(c).Association("Participants").Append(&User{
		Model: Model{ID: userID},
	})
}

func (c *Contest) DelUser(tx *gorm.DB, userID uint) error {
	return tx.Model(c).Association("Participants").Delete(&User{
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
	return len(rl)
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

	tx.Model(c).Association("Participants").Find(&c.Participants)
	tx.Model(c).Association("Problems").Find(&c.Problems)

	var problemIds []uint
	for _, p := range c.Problems {
		problemIds = append(problemIds, p.ID)
	}
	fromSQL := tx.Table("(?) as sub", tx.Model(&Submission{}).Where("?<submitted_at and " +
		"submitted_at<? and problem_id in (?)", c.BeginAt.Unix(), c.EndAt.Unix(), problemIds))

	for _, u := range c.Participants {
		var rank Rank
		err = fromSQL.Session(&gorm.Session{}).Where("user_id=? and result=?",
			u.ID, AC).Select("user_id, (count(*)) as score, (sum(submitted_at)) as penalty",
			).Scan(&rank).Error
		if err != nil {
			return nil, err
		}

		rankList = append(rankList, rank)
	}

	sort.Sort(rankList)
	return rankList, nil
}
