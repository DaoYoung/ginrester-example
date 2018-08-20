package endpoint

import (
	"github.com/DaoYoung/gorester"
	"ginrester-example/model"
)

type TeacherController struct {
	gorester.Controller
}
func (action *TeacherController) IsRestRoutePk() bool {
	return true
}
func (action *TeacherController) model() gorester.ResourceInterface {
	return &(model.Teacher{})
}
func (action *TeacherController) modelSlice() interface{} {
	return &[]model.Teacher{}
}
func (action TeacherController) Rester() (actionPtr *TeacherController) {
	action.Init(&action)
	return  &action
}