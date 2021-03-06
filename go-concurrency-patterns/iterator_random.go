// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	iterator("iterator!")
}

// START OMIT
func iterator(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

// STOP OMIT
