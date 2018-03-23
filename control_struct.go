package main

import "fmt"

func main(){
	i:=1
	for i<=4{
	fmt.Println(i)
	i++
	}
	for{
	fmt.Println("loop")
	break
	}
	if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }

   // fmt.Println(num)

}
