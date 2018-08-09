package hello

import (
	"sync"
	"fmt"
	"time"
	"math/rand"
)

var wg8 sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	wg8.Add(1)

	chanints := make(chan int)

	// TODO 第一个跑步者持有接力棒
	go Runner(chanints)

	chanints <- 1

	wg8.Wait()
}

func Runner(chanints chan int) {
	var newRunner int

	runner := <-chanints

	fmt.Printf("跑步者%d :开始绕着跑道跑步\n", runner)

	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("下一跑步者%d :准备接力跑步\n", newRunner)

		go Runner(chanints)

	}

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	if runner == 4 {
		fmt.Printf("比赛结束\n")
		wg8.Done()
		close(chanints)
		return
	}

	fmt.Printf("接力者:%d 接过跑步者: %d手中的棒\n", newRunner, runner)

	chanints <- newRunner
}
