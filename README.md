# Build api with gorester
This Example is about student and teacher at school.
* new student will go to school, like create a student
* you can find students who like basketball, and look for someone's information
* when someone transferred, it will be delete name at student book
* so as teachers
* Teacher has many students

## Create model
Embed `gorester.Model` in you code, like this

```
type Student struct {
	gorester.Model
	Name          string `json:"name"`
	Hobby         string `json:"hobby"`
	Age           int    `json:"age"`
	Sex           int    `json:"sex"`
	HeadTeacherId int    `json:"head_teacher_id"` // set up a one-to-one connection with teacher
}
```

## Create controller
Embed `gorester.Controller`, and declare 3 func: model(), modelSlice(), Rester()
```
type StudentController struct {
	gorester.Controller
}
//set up resource
func (action *StudentController) model() gorester.ResourceInterface {
	return &(model.Student{})
}
//for list result
func (action *StudentController) modelSlice() interface{} {
	return &[]model.Student{}
}
//controller factory
func (action StudentController) Rester() (actionPtr *StudentController) {
	action.Init(&action)
	return  &action
}
```
## Create route
simple and easy to create restful route.
```
gorester.CreateRoutes(school, endpoint.StudentController{}.Rester())
```
Append arguments at last, if you want personal actions.
```
gorester.CreateRoutes(school, endpoint.StudentController{}.Rester(), "create", "list", "info")
```

## Go build main.go
* you can create/update/list/findOne/delete student.
* you can insert some code before/after create/update/delete student.
```
func (action *StudentController) beforeCreate(c *gin.Context, m gorester.ResourceInterface) {
	//valid data
}
func (action *StudentController) afterCreate(c *gin.Context, m gorester.ResourceInterface) {
	//log something
}
```
* field condition in url will work, if you want filter list. Like this: <http://localhost:8080/school/students?hobby=basketball>
![filter_list.png](filter_list.png)

##So as teacher.
Create model, controller, route etc at the same way

## Teacher has many students

### Create HasManyStudentController
* It's contain `StudentController`, and declare `parentController func`.
* The resource teacher connect to student by field 'head_teacher_id', so we need declare `listCondition func`

### Duplicate resource ID in route
Default key is ":id", so it will confuse to get key's value at two level.
![route_key_duplicate.png](route_key_duplicate.png)
Declare func `IsRestRoutePk` will help to create unique key
![route_key.png](route_key.png)
Now, you can get students from teacher.