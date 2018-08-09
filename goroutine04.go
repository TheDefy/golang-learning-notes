package hello

import (
	"sync"
	"fmt"
	"sync/atomic"
	"runtime"
)

var (
	counter04 int64

	wg04 sync.WaitGroup
)

func main() {
	wg04.Add(2)

	go incCounter04(1)
	go incCounter04(2)

	wg04.Wait()
	fmt.Println("最后的counter值：", counter04)
}

func incCounter04(id int) {
	defer wg04.Done()

	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter04, 1)

		runtime.Gosched() // 从线程退出，并放回到队列 调度器切换goroutine
	}
}
