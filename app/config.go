package app

import (
	"github.com/jinzhu/configor"
)

var Config = struct {
	Port    string `default:":8080"`
	PerPage int    `default:"10"`
	Db struct {
		Host     string `default:"localhost"`
		Port     uint64 `default:"3306"`
		Name     string `default:"gorester"`
		User     string `default:"root"`
		Password string `default:"123456"`
	}
}{}

func InitConfig(configPath string) error {
	return configor.Load(&Config, configPath)
}
