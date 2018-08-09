package runner

import (
	"os"
	"time"
	"os/signal"
	"errors"
)

type Runner struct {

	interrupt chan os.Signal

	complete chan error

	timeout <-chan time.Time

	tasks []func(int)
}

/**
工厂函数返回Runner
 */
func New(t time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(t),
	}
}

// 添加任务
func (r *Runner) Add(tasks ...func(int)) {

	r.tasks = append(r.tasks, tasks...)

}

func (r *Runner) Start() error {
	// 希望接收所有的中断信息
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

var ErrTimeout = errors.New("接收超时")

var ErrInterrupt = errors.New("操作系统中断")

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.isInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}
func (r *Runner) isInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
