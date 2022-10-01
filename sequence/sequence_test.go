package sequence

import (
	"fmt"
	"testing"
)

func TestSequence(t *testing.T) {
	sequence1 := New(0, 1)
	for i := 0; i < 1000000; i++ {
		sequence1.NextId()
	}
	sequence2 := New(1, 1)
	fmt.Println(sequence2.NextId())
}
