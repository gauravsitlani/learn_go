package main

import "fmt"

func val(ival int){
	ival=0
}

func optr(iptr *int){
	*iptr=0
}

func main(){
	i:=1
	fmt.Println("initial value of i",i)
	optr(&i)
	fmt.Println(i)
	fmt.Println(&i)
	fmt.Println(i)
	a:=[3]int{1,2,3}
	b:=&a[0]
	fmt.Println("Value of b",*b)
}
