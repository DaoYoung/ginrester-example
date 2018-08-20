package handler

import (
	"github.com/gin-gonic/gin"
	"gorester"
	"ginrester-example/handler/endpoint"
)

func Serve(engine *gin.Engine) {
	school := engine.Group("/school")
	{
		gorester.CreateRoutes(school, endpoint.StudentController{}.Rester())
		gorester.CreateRoutes(school, endpoint.TeacherController{}.Rester())
		gorester.CreateRoutes(school, endpoint.HasManyStudentController{}.Rester(), "list")
	}
}
