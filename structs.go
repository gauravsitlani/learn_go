package main

import "fmt"

type animal struct{
	name string
	age int
}

//playing with structures

func main(){
	s := animal{name:"Tiger",age:10}
	fmt.Println(s)
	new_s:=&s
	new_s.age=99
	fmt.Println(new_s)
	fmt.Println(s)
}
