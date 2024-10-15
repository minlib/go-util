package taskutil

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestConcurrentExecute(t *testing.T) {
	tasks := []func(){func() {
		time.Sleep(1 * time.Second)
		panic(errors.New("task 1 panic"))
	}, func() {
		time.Sleep(2 * time.Second)
		fmt.Println("task 2 done")
	}}
	ConcurrentExecute(tasks...)
	fmt.Println("所有任务执行完成")
}

func TestConcurrentExecute2(t *testing.T) {
	var tasks []func()
	for i := 0; i < 100; i++ {
		tasks = append(tasks, func() {
			fmt.Println(fmt.Sprintf("i:%d start", i))
			time.Sleep(3 * time.Second)
			fmt.Println(fmt.Sprintf("i:%d end", i))
		})
	}
	ConcurrentExecute(tasks...)
	fmt.Println("所有任务执行完成")
}

func TestTaskStop(t *testing.T) {
	taskMap := make(map[int64]*time.Timer)

	duration := 5 * time.Second

	go func(taskMap map[int64]*time.Timer) {
		fmt.Println("task start")
		timer := time.NewTimer(duration)
		taskMap[123] = timer
		<-timer.C
		fmt.Println("task run")
	}(taskMap)

	go func(taskMap map[int64]*time.Timer) {
		time.Sleep(1 * time.Second)
		fmt.Println("task in 2")
		taskMap[123].Stop()
		fmt.Println("task in 3")
	}(taskMap)

	time.Sleep(10 * time.Second)
	fmt.Println("end")
}

func TestConcurrentExecuteResult(t *testing.T) {
	task1 := func() error { return fmt.Errorf("task1 error") }
	task2 := func() error { return nil }
	task3 := func() error { panic("task3 panicked") }

	results := ConcurrentExecuteResult(task1, task2, task3)
	for _, result := range results {
		if result.Error != nil {
			fmt.Printf("Task %d had an error: %v\n", result.Index, result.Error)
		} else {
			fmt.Printf("Task %d completed successfully.\n", result.Index)
		}
	}
}
