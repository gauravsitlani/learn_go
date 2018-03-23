package main

import (
"fmt"
)

func main(){
	m:= make(map[string]int)
	m["a"]=1
	m["b"]=2
	fmt.Println(m)
	fmt.Println(len(m))
	targetMap := make(map[string]int)

	// Copy from the original map to the target map
	for key,value := range m {
	  targetMap[key] = value
  	}
  	fmt.Println(targetMap)
	_,prs:=m["a"]
	//fmt.Println(val)
	fmt.Println(prs)
 }

