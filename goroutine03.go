package hello

import (
	"sync"
	"runtime"
	"fmt"
)

var (
	counter int

	wg03 sync.WaitGroup
)

func main() {

	fmt.Println("CPU NUM:",runtime.NumCPU())

	wg03.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg03.Wait()
	fmt.Println("最后counter的值：", counter)// TODO 为什么结果 有时是2 有时是4 按理都是2才对？

}

func incCounter(id int) {
	defer wg03.Done()

	for inner := 0; inner < 2; inner++ {
		value := counter
		runtime.Gosched() // TODO 在赋值之前 强制调度器切换goroutine
		value++
		counter = value
	}
}
