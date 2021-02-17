# Go ParseTag

[![Coverage Status](https://coveralls.io/repos/github/yitsushi/go-parsetag/badge.svg?branch=main)](https://coveralls.io/github/yitsushi/go-parsetag?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/yitsushi/go-parsetag)](https://goreportcard.com/report/github.com/yitsushi/go-parsetag)
[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/yitsushi/go-parsetag) [![Join the chat at https://gitter.im/yitsushi/go-parsetag](https://badges.gitter.im/yitsushi/go-parsetag.svg)](https://gitter.im/yitsushi/go-parsetag?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Parse tags from structs with a compact callback function.

## Example

```go
package main

import (
	"fmt"
	"github.com/yitsushi/go-parsetag"
)

func main() {
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
```

Check `parse_test.go` for more examples.
