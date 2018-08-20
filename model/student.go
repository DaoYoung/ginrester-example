package model

import (
	"github.com/DaoYoung/ginrester"
)

type Student struct {
	ginrester.Model
	Name          string `json:"name"`
	Hobby         string `json:"hobby"`
	Age           int    `json:"age"`
	Sex           int    `json:"sex"`
	HeadTeacherId int    `json:"head_teacher_id"`
}
