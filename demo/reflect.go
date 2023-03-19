package demo

import (
	"fmt"
	"reflect"
)

func ref(input any) {
	t := reflect.TypeOf(input)
	v := reflect.ValueOf(input)
	fmt.Println(t, v)
}
