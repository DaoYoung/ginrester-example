package model

import "gorester"

type Student struct {
	gorester.BaseFields
	Name string
	Age int
	Sex int
}
