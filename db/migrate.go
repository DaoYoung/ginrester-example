package db

import (
	"ginrester-example/model"
	"github.com/DaoYoung/ginrester"
)

func InitTables() {
	ginrester.Db.AutoMigrate(&model.Student{}, &model.Teacher{})
}