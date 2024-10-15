package taskutil

import (
	"fmt"
	"sync"
)

type TaskError struct {
	Index int
	Error error
}

// ConcurrentExecute 并发执行任务的函数
func ConcurrentExecute(tasks ...func()) {
	var wg sync.WaitGroup
	wg.Add(len(tasks))

	for _, task := range tasks {
		go func(t func()) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("recovered from panic:", r)
				}
				wg.Done()
			}()
			t()
		}(task)
	}

	wg.Wait()
}

// ConcurrentExecuteResult 并发执行任务的函数
func ConcurrentExecuteResult(tasks ...func() error) []TaskError {
	var wg sync.WaitGroup
	wg.Add(len(tasks))

	results := make([]TaskError, len(tasks))
	for i, task := range tasks {
		go func(index int, t func() error) {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					results[index] = TaskError{Index: index, Error: fmt.Errorf("recovered from panic: %v", r)}
				}
			}()
			err := t()
			if err != nil {
				results[index] = TaskError{Index: index, Error: err}
			} else {
				results[index] = TaskError{Index: index, Error: nil}
			}
		}(i, task)
	}

	wg.Wait()
	return results
}
