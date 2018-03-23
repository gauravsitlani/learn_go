package main

import "fmt"

func main(){
	var a [5]int
	fmt.Println("emp : ",a)
	a[4]=100
	fmt.Println(a)
	fmt.Println(a[4])
	fmt.Println(len(a))
	b:= []int{1,2,3,4,5}
	fmt.Println(len(b))

	var t2D [3][3]int
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
		t2D[i][j]=i+j
		}
	}
	fmt.Println(t2D)

}
