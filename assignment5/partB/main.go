package main

import (
	"fmt"
)

var global int
var done = make(chan bool)

func count() {
	for i := 0; i < 10000; i++ {
		global++
	}
	done <- true
}

func main() {
	go count()
	<-done
	go count()
	<-done
	fmt.Println(global)
}