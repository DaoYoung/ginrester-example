package endpoint

import (
	"github.com/DaoYoung/ginrester"
	"ginrester-example/model"
	"github.com/gin-gonic/gin"
)

type StudentController struct {
	ginrester.Controller
}
func (action *StudentController) model() ginrester.ResourceInterface {
	return &(model.Student{})
}
func (action *StudentController) modelSlice() interface{} {
	return &[]model.Student{}
}
func (action StudentController) Rester() (actionPtr *StudentController) {
	action.Init(&action)
	return  &action
}
func (action *StudentController) afterCreate(c *gin.Context, m ginrester.ResourceInterface) {
	//todo log something
}