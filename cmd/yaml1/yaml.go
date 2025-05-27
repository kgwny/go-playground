package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"

	"gopkg.in/yaml.v3"
)

var data = `
a: test
b: 6
`

func parse(data string, dest interface{}) error {
	if reflect.TypeOf(dest).Kind() != reflect.Ptr {
		return errors.New("mapping: dest is not pointer")
	}
	parseError := yaml.Unmarshal([]byte(data), dest)
	return parseError
}

func mapping(source interface{}, copy interface{}) error {
	data, sourceError := yaml.Marshal(source)
	if sourceError != nil {
		return sourceError
	}
	if reflect.TypeOf(copy).Kind() != reflect.Ptr {
		return errors.New("mapping: copy is not pointer")
	}
	destError := yaml.Unmarshal(data, copy)
	return destError
}

func main() {
	var schema struct {
		A string
		B int
	}
	var copy struct {
		A string
	}
	if err := parse(data, &schema); err != nil {
		log.Fatalf("error: %v", err)
	}
	if err := mapping(&schema, &copy); err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(schema)
	fmt.Println(copy)
}
