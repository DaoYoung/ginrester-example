package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/DaoYoung/ginrester"
	"ginrester-example/handler/endpoint"
)

func Serve(engine *gin.Engine) {
	school := engine.Group("/school")
	{
		ginrester.CreateRoutes(school, endpoint.StudentController{}.Rester())
		ginrester.CreateRoutes(school, endpoint.TeacherController{}.Rester())
		ginrester.CreateRoutes(school, endpoint.HasManyStudentController{}.Rester(), "list", "info")
	}
}
