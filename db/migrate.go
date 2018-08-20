package db

import (
	"ginrester-example/model"
	"github.com/DaoYoung/gorester"
)

func InitTables() {
	gorester.Db.AutoMigrate(&model.Student{}, &model.Teacher{})
}