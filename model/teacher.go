package model

import "github.com/DaoYoung/gorester"

type Teacher struct {
	gorester.Model
	Name  string `json:"name"`
	Major string `json:"major"`
}
