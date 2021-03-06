package endpoint

import (
	"github.com/DaoYoung/ginrester"
	"github.com/gin-gonic/gin"
)

type HasManyStudentController struct {
	StudentController
}

func (action HasManyStudentController) Rester() (actionPtr *HasManyStudentController) {
	action.Init(&action)
	return &action
}
func (action *HasManyStudentController) parentController() ginrester.ControllerInterface {
	return TeacherController{}.Rester()
}
func (action *HasManyStudentController) listCondition(c *gin.Context) map[string]interface{} {
	condition := make(map[string]interface{})
	condition["head_teacher_id"] = c.Param("teacher_id")
	return condition
}
