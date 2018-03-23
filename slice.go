package main

import (
	"fmt"
)

func main(){
	s:=make([]string,3)
	fmt.Println(s)
	s[0]="1"
	s[1]="2"
	s[2]="a"
	fmt.Println(s)
	fmt.Println(len(s))
	s = append(s,"ss")
	fmt.Println(s)
	//creating a new slice of same length and copying
	c:= make([]string,len(s))
	copy(c,s)
	fmt.Println("copied slice",c)
	//slice init
	
}
