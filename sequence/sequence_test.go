package sequence

import (
	"fmt"
	"testing"
)

func TestSequence(t *testing.T) {
	var sequence = New(0, 1)
	for i := 0; i < 1000000; i++ {
		sequence.NextId()
	}
	var sequence1 = New(1, 1)
	fmt.Println(sequence1.NextId())
}
