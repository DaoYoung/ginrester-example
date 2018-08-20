package model

import "github.com/DaoYoung/ginrester"

type Teacher struct {
	ginrester.Model
	Name  string `json:"name"`
	Major string `json:"major"`
}
