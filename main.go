package main

import (
	"fmt"
	"reflect"

	"./fruit/orange"
)

func main() {
	// apple := apple.Apple{}
	orange := orange.Orange{
		ID:   1,
		Name: "Orange Name",
	}

	test(orange)
}

func test(fruit interface{}) {

	v := reflect.ValueOf(fruit)
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fmt.Println(t.Field(i).Name)
	}

	if c, ok := fruit.(orange.Orange); ok {
		fmt.Println("ID", c.ID)
		fmt.Println("Name", c.Name)
	}
}
