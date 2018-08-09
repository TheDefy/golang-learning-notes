package main

import (
	"sync"
	"fmt"
	"runtime"
)

var wg sync.WaitGroup

func main() {

	runtime.GOMAXPROCS(1)
	wg.Add(2)
	fmt.Println("goroutine 开始")

	go findSuShu("A")
	go findSuShu("B")
	fmt.Println("goroutine 等待结束")
	wg.Wait()

	fmt.Println("goroutine 结束")

}

func findSuShu(sign string) {

	defer wg.Done()

next:
	for count := 2; count < 5000; count++ {
		for inner := 2; inner < count; inner++ {
			if count%inner == 0 {
				continue next
			}
		}
		fmt.Printf("分组：%s %d\n", sign, count)
	}
	fmt.Println("完成的分组", sign)

}
