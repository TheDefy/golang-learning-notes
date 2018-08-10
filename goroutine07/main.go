package main

import (
	"sync"
	"math/rand"
	"time"
	"fmt"
)

var wg7 sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	court := make(chan int) // 创建一个无缓存的channel

	wg7.Add(2)

	// TODO 启动两个选手

	go player("cl", court)
	go player("yx", court)

	court <- 1

	wg7.Wait()

}

func player(name string, court chan int) {

	defer wg7.Done()

	for {
		ball, ok := <-court
		if !ok { // 如果通道被关闭 我们就赢了
			fmt.Printf("玩家%s赢了\n", name)
			return
		}

		n := rand.Intn(100)
		fmt.Println("rand num:", n)
		if n%13 == 0 {
			fmt.Printf("玩家%s丢球了\n", name)
			close(court)
			return
		}

		fmt.Printf("玩家%s接住了球并击打回去%d\n", name, ball)
		ball++

		court <- ball
	}

}
