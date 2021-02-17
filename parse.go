package parsetag

import "reflect"

// Callback is a func that will be called on field match.
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
