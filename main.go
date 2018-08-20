package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"ginrester-example/handler"
	"ginrester-example/app"
)

func main() {
	if err := app.InitConfig("config.yml"); err != nil {
		log.Fatal(err)
	}
	if err := app.InitDb(); err != nil {
		log.Fatal(err)
	}
	r := gin.New()
	r.Use(gin.Logger())
	handler.Serve(r)
	r.Run(app.Config.Port)
}