package lib

import (
	"fmt"
	"reflect"
)

type TwoDimensionalArray []interface{}

func (tda TwoDimensionalArray) String() string {
	var s string
	for _, array := range tda {
		s += fmt.Sprintln(array)
	}

	return s
}

func isSlice(arg interface{}) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)

	if val.Kind() == reflect.Slice {
		ok = true
	}

	return
}

func Slice(slice interface{}) []interface{} {
	val, ok := isSlice(slice)
	if !ok {
		return nil
	}

	length := val.Len()

	out := make([]interface{}, length)
	for i := 0; i < length; i++ {
		out[i] = val.Index(i).Interface()
	}

	return out
}
