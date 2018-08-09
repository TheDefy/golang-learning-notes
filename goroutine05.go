package hello

import (
	"sync"
	"time"
	"fmt"
	"sync/atomic"
)

var (
	shutDown int64
	wg05     sync.WaitGroup
)

func main() {
	wg05.Add(2)

	go doWork("A")
	go doWork("B")

	time.Sleep(1 * time.Second)

	fmt.Println("该停止工作了")
	atomic.StoreInt64(&shutDown, 2)
	wg05.Wait()
}

func doWork(name string) {

	defer wg05.Done()

	for {
		fmt.Printf("%s 正在工作...\n", name)

		time.Sleep(200 * time.Millisecond)

		if atomic.LoadInt64(&shutDown) == 2 {
			fmt.Printf("%s 工作停止\n", name)
			break
		}
	}
}
