package endpoint

import (
	"gorester"
	"ginrester-example/model"
)

type StudentController struct {
	gorester.Controller
}
func (action *StudentController) model() gorester.ResourceInterface {
	return &(model.Student{})
}
func (action *StudentController) modelSlice() interface{} {
	return &[]model.Student{}
}
func (action StudentController) Rester() (actionPtr *StudentController) {
	action.Init(&action)
	return  &action
}