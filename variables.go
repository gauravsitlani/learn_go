package main

import (
	"fmt"
	"reflect"
)

func main(){
	var a string = "string_value"
	fmt.Println(a)
	var e int
	fmt.Println(e)
	f:="any_value"
	fmt.Println(reflect.TypeOf(f))
}
