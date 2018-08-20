package endpoint

import (
	"github.com/DaoYoung/ginrester"
	"ginrester-example/model"
)

type TeacherController struct {
	ginrester.Controller
}
func (action *TeacherController) IsRestRoutePk() bool {
	return true
}
func (action *TeacherController) model() ginrester.ResourceInterface {
	return &(model.Teacher{})
}
func (action *TeacherController) modelSlice() interface{} {
	return &[]model.Teacher{}
}
func (action TeacherController) Rester() (actionPtr *TeacherController) {
	action.Init(&action)
	return  &action
}