package main

import(
	"fmt"
	"math"
)

const s string = "constant_st"

func main(){
	fmt.Println(s)
	//s = "new_string"
	//fmt.Println(s)
	const n = 5000
	//n = n+10
	//fmt.Println(n)
	const v = 4e5/n
	fmt.Println(v)
	fmt.Println(math.Sin(n))
}
