package main

import "fmt"

func loop(sign string) {
	for count := 0; count < 5; count++ {
		fmt.Printf("loop %s :%d \n", sign, count)
	}

	if sign == "go" {
		<-ints
	}

}

var ints chan int = make(chan int)

func main() {

	//ints := make(chan int)
	loop("main")

	go loop("go")

	ints <- 0

}
