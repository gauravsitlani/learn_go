package main

import "fmt"

func main() {
	mess := make(chan string, 2)
	mess <- "buffered"
	mess <- " channel"
	//mess <- " go lang "
	for i := 0; i < 2; i++ {
		fmt.Print(<-mess)
	}
}
