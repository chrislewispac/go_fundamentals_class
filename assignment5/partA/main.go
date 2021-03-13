package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	for i := 0; i < 5; i++ {
		go func() {
			c <- i
		}()
		fmt.Println(<-c)
	}
}