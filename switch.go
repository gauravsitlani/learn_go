package main

import (
	"fmt"
	"time"
)

func main(){
	//i:=1
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}
	t := time.Now()
	fmt.Println(t)
	switch{
	case t.Hour() < 12:
		fmt.Println("Before noon")
	default:
		fmt.Println("Afternoon!")
	}
	whatAmI := func(i interface{}) {
		switch t:=i.(type){
		case bool:
			fmt.Println("Boolean")
		
		case int:
			fmt.Println("Interger")
		default:
			fmt.Printf("type of var %T\n",t)
		}
	}
	fmt.Println(whatAmI)
	whatAmI(123.12)
	fmt.Println(whatAmI)
}
