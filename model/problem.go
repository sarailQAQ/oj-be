
/**
 * @author: sarail
 * @time: 2021/6/8 22:07
**/

package model

// 问题的状态常量
const (
	// 正常状态，所有用户都可见
	problemNormal = 1

	// 仅自己可见
	problemOwnerOnly = 2

	// 比赛题目，比赛开始前
	problemBeforeContest = 3

	// 比赛题目，且比赛正在进行中
	problemInContest = 4
)

type Problem struct {
	model

	UserID uint

	// status: 题目的状态，有以下几种
	Status uint8

	// 标题、题目描述、输入描述、输出描述等
	Tittle            string
	Description       string
	InputDescription  string
	OutputDescription string
	Tip               string
}
