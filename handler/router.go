package handler

import (
	"github.com/gin-gonic/gin"
	"gorester"
	"ginrester-example/handler/endpoint"
)

func Serve(engine *gin.Engine) {
	//engine.POST("/schools", endpoint.UserController{}.Rester().Create)
	school := engine.Group("/school")
	{
		gorester.CreateRestRoutes(school, endpoint.StudentController{}.Rester())
		gorester.CreateRestRoutes(school, endpoint.StudentController{}.Rester(), "list", "create", "update", "delete")

	}
}
