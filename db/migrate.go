package db

import (
	"ginrester-example/model"
	"gorester"
)

func InitTables() {
	gorester.Db.AutoMigrate(&model.Student{}, &model.Teacher{})
}