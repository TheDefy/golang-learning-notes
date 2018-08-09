package main

import (
	"time"
	"log"
	"os"
	"thedefy/LearningNotes/runnertest/runner"
)

const timeout = 3 * time.Second

func main() {
	log.Println("开始任务....")

	r := runner.New(timeout)
	r.Add(createTask(), createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println(err)
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println(err)
			os.Exit(2)
		}
	}

	log.Println("任务结束....")
}

func createTask() func(int) {
	return func(i int) {
		log.Printf("正在执行的任务%d", i)
		i2 := time.Duration(i)
		log.Printf("sleep time %d", i2)
		time.Sleep(time.Second * i2)
	}

}
