package hello

import (
	"sync"
	"fmt"
	"runtime"
)

var (
	counter6 int
	wg6      sync.WaitGroup
	mutex    sync.Mutex
)

func main() {
	wg6.Add(2)

	// TODO goroutine
	go incCounter6("A")
	go incCounter6("B")

	wg6.Wait()
	fmt.Println("counter最后的值：", counter6)
}

func incCounter6(sign string) {
	defer wg6.Done()

	for count := 0; count < 2; count++ {

		mutex.Lock()
		{
			value := counter6
			runtime.Gosched()
			value++
			counter6 = value
		}
		mutex.Unlock()
	}
}
