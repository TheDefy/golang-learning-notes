package main

import (
	"fmt"
	"math"
	"thedefy/LearningNotes/fortest/fortest"
)

func main() {
	fmt.Println(
		fortest.ConvertToBin(3),
		fortest.ConvertToBin(5),
	)

	fortest.PrintFile(".gitignore")

	商, 余数 := fortest.Div(13, 3)

	fmt.Println(商, 余数)

	i := fortest.Apply(func(i int, i2 int) int {
		return int(math.Max(float64(i), float64(i2)))
	}, 3, 4)

	fmt.Printf("funclearn apply rusult: %d \n", i)

	fmt.Println("sum result:", fortest.Sum(1, 2, 3))
}
