package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	Field1 string `mytag:"something"`
	Field2 string `mytag:"something2"`
}

func main() {
	s := MyStruct{Field1: "asd", Field2: "dsa"}

	extracted := map[string]string{}

	Parse(s, "mytag", func(tagValue string, fieldType string, fieldValue interface{}) {
		if fieldType == "string" {
			extracted[tagValue] = fieldValue.(string)
		}
	})

	fmt.Printf("%v\n", s)
	fmt.Printf("%v\n", extracted)
}

type Callback func(tagValue string, fieldType string, fieldValue interface{})

// Parse tags on a struct.
func Parse(data interface{}, tagName string, callback Callback) {
	v := reflect.ValueOf(data)
	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get(tagName)
		if tag == "" || tag == "-" {
			continue
		}

		callback(tag, v.Field(i).Type().Name(), v.Field(i).Interface())
	}
}
