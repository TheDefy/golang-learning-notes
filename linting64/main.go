package main

import (
	"fmt"
	"thedefy/LearningNotes/linting64/counters"
)

var a int = 1

var b int = 2

var c, d int = 3, 4

func main() {
	//counters := counters.alertCounter(10)
	counters := counters.New(10)

	fmt.Printf("Counter: %d\n", counters)

	fmt.Printf("a: %d, b:%d\n", a, b)
	//changeNub(a, b)
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Printf("a: %d, b:%d\n", a, b)

	dur = duration(10)

	fmt.Printf("duration: %d\n", dur)
}

func changeNub(a int, b int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
}

type duration int

var dur duration
