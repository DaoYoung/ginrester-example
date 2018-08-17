package gorester

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/model"
	"strconv"
	"net/http"
	"apibuilder-server/helper"
	"apibuilder-server/app"
)

type ModelController struct {
	Controller
}
func (action *ModelController) model() DaoInterface {
	return &(ApiModel{})
}
func (action *ModelController) modelSlice() interface{} {
	return &[]ApiModel{}
}
func (action ModelController) Rester() (actionPtr *ModelController) {
	action.init(&action)
	return  &action
}
func (action *ModelController) beforeCreate(c *gin.Context, m DaoInterface) {
	user := GetUserFromToken(c)
	m.(*ApiModel).AuthorId = user.ID
}
func (action *ModelController) IsRestRoutePk() bool{
	return true
}
func (action *ModelController) RouteName() string {
	return "model"
}
func NoteModel(c *gin.Context) {
	var jsonForm ApiModelNote
	var info interface{}
	err := c.BindJSON(&jsonForm)
	if err != nil {
		panic(JsonTypeError(err))
	}
	jsonForm.ModelId, _ = strconv.Atoi(c.Param("id"))
	dbNote := ApiModelNote{ModelId: jsonForm.ModelId, ParentId: jsonForm.ParentId, ModelKey: jsonForm.ModelKey}
	ExsitAndFirst(&dbNote)
	if dbNote.ID >0 {
		Delete(dbNote, dbNote.ID)
	}
	info = Create(&jsonForm)
	helper.ReturnSuccess(c, http.StatusOK, info)
}

func NoteModelDetail(c *gin.Context) {
	condition := make(map[string]interface{})
	id, _ := strconv.Atoi(c.Param("id"))
	condition["model_id"] = id
	modelNotes := &([]ApiModelNote{})
	FindListWhereMap(modelNotes, condition, "", 1, app.Config.PerPage)
	helper.ReturnSuccess(c, http.StatusOK, modelNotes)
}

type ModelMapController struct {
	Controller
}
func (action ModelMapController) Rester() ControllerInterface {
	action.Controller.Rester = &action
	action.Controller.RestModel = func() DaoInterface { return &(ApiModelMap{}) }
	action.Controller.RestModelSlice = func() interface{} { return &[]ApiModelMap{} }
	action.Controller.ParentController = ModelController{}.Rester()
	return  &action
}
func (action *ModelMapController) RouteName() string {
	return "map"
}
func (action *ModelMapController) BeforeRest(c *gin.Context, m DaoInterface) {
	user := GetUserFromToken(c)
	m.(*ApiModelMap).AuthorId = user.ID
}
