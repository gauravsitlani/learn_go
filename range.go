package main

import "fmt"

func main(){
	nums := []int{1,2,3}
	sum:=0
	for _,nums:=range nums{
	sum+=nums
	}
	fmt.Println(sum)

}
