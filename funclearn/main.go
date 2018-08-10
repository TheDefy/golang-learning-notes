package main

import (
	"fmt"
	"thedefy/LearningNotes/funclearn/functest"
)

func main() {
	result := functest.ApplyFunc(functest.MaxNumber, 3, 4)
	fmt.Println("applyFunc result:", result)
}
