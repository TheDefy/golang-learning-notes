package hello

import (
	"math/rand"
	"time"
	"sync"
	"fmt"
)

const (
	workerNub = 4
	sendNub   = 10
)

var wg9 sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	wg9.Add(workerNub)

	chanstrings := make(chan string, sendNub)

	// TODO goroutine
	for count := 1; count <= workerNub; count++ {
		go doWork9(chanstrings, count)
	}

	for count := 1; count <= sendNub; count++ {
		chanstrings <- fmt.Sprintf("任务：%d", count)
	}

	close(chanstrings) //关闭通道 goroutine仍然可以从通道接受数据 但是不能再向通道发送数据

	wg9.Wait()
}

func doWork9(chanstrings chan string, worker int) {

	defer wg9.Done()

	for {
		s, ok := <-chanstrings
		if !ok {
			fmt.Printf("伟大的劳动者：%d 已经完工！！！\n", worker)
			break
		}

		fmt.Printf("伟大的劳动者：%d 开始：%s\n", worker, s)
		intn := rand.Intn(100)
		time.Sleep(time.Duration(intn) * time.Millisecond)

		fmt.Printf("伟大的劳动者：%d 完成：%s\n", worker, s)

	}

}
