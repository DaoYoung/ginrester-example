package model

import (
	"github.com/DaoYoung/gorester"
)

type Student struct {
	gorester.Model
	Name          string `json:"name"`
	Hobby         string `json:"hobby"`
	Age           int    `json:"age"`
	Sex           int    `json:"sex"`
	HeadTeacherId int    `json:"head_teacher_id"`
}
