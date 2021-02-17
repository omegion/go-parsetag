package parsetag_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yitsushi/go-parsetag"
)

const stringType = "string"

func TestParse_simple(t *testing.T) {
	type MyStruct struct {
		Field1 string `mytag:"something"`
		Field2 string `mytag:"something2"`
		Field3 string `mytag:""`
		Field4 string `mytag:"-"`
	}

	data := MyStruct{
		Field1: "value1",
		Field2: "value2",
		Field3: "value3",
		Field4: "value4",
	}
	extracted := map[string]string{}

	parsetag.Parse(data, "mytag", func(tagValue string, fieldType string, fieldValue interface{}) {
		if fieldType == stringType {
			extracted[tagValue] = fieldValue.(string)
		}
	})

	assert.Equal(t, map[string]string{"something": "value1", "something2": "value2"}, extracted)
}

func TestParse_multitype(t *testing.T) {
	type MyStruct struct {
		Field1 string `mytag:"string"`
		Field2 int64  `mytag:"integer"`
		Field3 bool   `mytag:"boolean"`
	}

	data := MyStruct{Field1: "value1", Field2: 99, Field3: true}
	extracted := map[string]string{}

	parsetag.Parse(data, "mytag", func(tagValue string, fieldType string, fieldValue interface{}) {
		switch fieldType {
		case "string":
			extracted[tagValue] = fieldValue.(string)
		case "int64":
			extracted[tagValue] = fmt.Sprintf("%d", fieldValue.(int64))
		case "bool":
			if fieldValue.(bool) {
				extracted[tagValue] = "true"
			} else {
				extracted[tagValue] = "false"
			}
		}
	})

	assert.Equal(t, map[string]string{"string": "value1", "integer": "99", "boolean": "true"}, extracted)
}

func TestParse_complex(t *testing.T) {
	type MyStruct struct {
		Field1 string `mytag:"something"`
		Field2 string `mytag:"something2,case=upper,cut=4"`
	}

	data := MyStruct{Field1: "value1", Field2: "value2"}
	extracted := map[string]string{}

	parsetag.Parse(data, "mytag", func(tagValue string, fieldType string, fieldValue interface{}) {
		if fieldType != "string" {
			return
		}

		var (
			name  string
			value string = fieldValue.(string)
		)

		for _, part := range strings.Split(tagValue, ",") {
			if strings.Contains(part, "=") {
				kv := strings.SplitN(part, "=", 2)

				switch kv[0] {
				case "case":
					value = strings.ToUpper(value)
				case "cut":
					cut, _ := strconv.Atoi(kv[1])
					value = value[0:cut]
				}
			} else {
				name = part
			}
		}

		extracted[name] = value
	})

	assert.Equal(t, map[string]string{"something": "value1", "something2": "VALU"}, extracted)
}

func ExampleParse() {
	type MyStruct struct {
		Field1 string `mytag:"something"`
		Field2 string `mytag:"something2"`
	}

	data := MyStruct{Field1: "asd", Field2: "dsa"}

	extracted := map[string]string{}

	parsetag.Parse(data, "mytag", func(tagValue string, fieldType string, fieldValue interface{}) {
		if fieldType == stringType {
			extracted[tagValue] = fieldValue.(string)
		}
	})

	fmt.Printf("%v\n", data)
	fmt.Printf("%v\n", extracted)
}
