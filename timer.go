package main

import (
"fmt"
"time"
)

func main(){
	t1:=time.NewTimer(time.Second)
	t1.Start()
	fmt.Println(t1)
}
