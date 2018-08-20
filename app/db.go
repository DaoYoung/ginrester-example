package app

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"github.com/DaoYoung/gorester"
)

func InitDb() error {
	var err error
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		Config.Db.User,
		Config.Db.Password,
		Config.Db.Host,
		Config.Db.Port,
		Config.Db.Name,
	)
	if gorester.Db, err = gorm.Open("mysql", dsn); err != nil {
		return err
	}
	gorester.Db.LogMode(true)
	return nil
}