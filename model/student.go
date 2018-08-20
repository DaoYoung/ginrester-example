package model

import (
	"gorester"
)

type Student struct {
	gorester.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  int    `json:"sex"`
	HeadTeacherId  int    `json:"head_teacher_id"`
}