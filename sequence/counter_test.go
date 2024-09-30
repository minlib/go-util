package sequence

import (
	"fmt"
	"testing"
	"time"
)

func TestCounter_NextId(t *testing.T) {
	var counter = new(Counter)
	for i := 1; i <= 10000; i++ {
		counter.NextId()
	}
	fmt.Println(counter.NextId())
}

func TestCounter_NextId2(t *testing.T) {
	var counter = NewCounter(10000)
	for i := 1; i <= 10000; i++ {
		counter.NextId()
	}
	fmt.Println(counter.NextId())
}

func TestCounter_NextId3(t *testing.T) {
	var counter = new(Counter)
	counter.SetValue(20000)
	for i := 1; i <= 10000; i++ {
		counter.NextId()
	}
	fmt.Println(counter.NextId())
}

func TestCounter_NextId4(t *testing.T) {
	var counter = NewCounter(30000)
	for i := 1; i <= 10000; i++ {
		go func() {
			counter.NextId()
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(counter.NextId())
}
