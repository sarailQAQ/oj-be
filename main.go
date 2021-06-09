package main

import (
	"oj-be/dao"
	"oj-be/model"
)

func main() {
	model.Init(dao.OriginDB())
}
