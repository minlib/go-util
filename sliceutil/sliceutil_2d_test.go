package sliceutil

import (
	"fmt"
	"testing"
)

func TestDistinct2D(t *testing.T) {
	fmt.Println("--------- Distinct2D[int] ---------")
	fmt.Println(Distinct2D([][]int{{1, 2}, {3, 4}, {5, 6, 7}}))                 // [[1 2] [3 4] [5 6 7]]
	fmt.Println(Distinct2D([][]int{{1, 2}, {1, 2}, {3, 4}, {3, 4}, {5, 6, 7}})) // [[1 2] [3 4] [5 6 7]]
	fmt.Println(Distinct2D([][]int{{2, 1}, {1, 2}, {3, 4}, {3, 4}, {5, 6, 7}})) // [[2 1] [1 2] [3 4] [5 6 7]]

	fmt.Println("--------- Distinct2D[string] ---------")
	fmt.Println(Distinct2D([][]string{{"A", "B"}, {"C", "D"}, {"E", "F", "G"}}))                         // [[A B] [C D] [E F G]]
	fmt.Println(Distinct2D([][]string{{"A", "B"}, {"A", "B"}, {"C", "D"}, {"C", "D"}, {"E", "F", "G"}})) // [[A B] [C D] [E F G]]
	fmt.Println(Distinct2D([][]string{{"B", "A"}, {"A", "B"}, {"C", "D"}, {"C", "D"}, {"E", "F", "G"}})) // [[B A] [A B] [C D] [E F G]]
}

var i1 = [][]int{{1, 2}, {3, 4, 5}, {6, 7, 8, 9}}
var i2 = [][]int{{1, 2}, {3, 4, 5}, {6, 7, 8, 9}}
var i3 = [][]int{{1, 3}, {3, 4, 5}, {6, 7, 8, 9}}
var i4 = [][]int{{2, 1}, {4, 5, 3}, {6, 9, 8, 7}}
var i5 = [][]int{{4, 5, 3}, {2, 1}, {6, 8, 9, 7}}

var s1 = [][]string{{"A", "B"}, {"C", "D", "E"}, {"F", "G", "H", "I"}}
var s2 = [][]string{{"A", "B"}, {"C", "D", "E"}, {"F", "G", "H", "I"}}
var s3 = [][]string{{"A", "C"}, {"C", "D", "E"}, {"F", "G", "H", "I"}}
var s4 = [][]string{{"B", "A"}, {"D", "E", "C"}, {"H", "I", "F", "G"}}
var s5 = [][]string{{"D", "E", "C"}, {"B", "A"}, {"H", "I", "F", "G"}}

func TestEqual2D(t *testing.T) {
	fmt.Println("--------- Equal2D[int] ---------")
	fmt.Println(Equal2D(i1, i2)) // true
	fmt.Println(Equal2D(i1, i3)) // false
	fmt.Println(Equal2D(i1, i4)) // false
	fmt.Println(Equal2D(i1, i5)) // false

	fmt.Println("--------- Equal2D[string] ---------")
	fmt.Println(Equal2D(s1, s2)) // true
	fmt.Println(Equal2D(s1, s3)) // false
	fmt.Println(Equal2D(s1, s4)) // false
	fmt.Println(Equal2D(s1, s5)) // false
}

func TestEqualIgnoreOrder2D(t *testing.T) {
	fmt.Println("--------- EqualIgnoreOrder2D[int] ---------")
	fmt.Println(EqualIgnoreOrder2D(i1, i2)) // true
	fmt.Println(EqualIgnoreOrder2D(i1, i3)) // false
	fmt.Println(EqualIgnoreOrder2D(i1, i4)) // true
	fmt.Println(EqualIgnoreOrder2D(i1, i5)) // true

	fmt.Println("--------- EqualIgnoreOrder2D[string] ---------")
	fmt.Println(EqualIgnoreOrder2D(s1, s2)) // true
	fmt.Println(EqualIgnoreOrder2D(s1, s3)) // false
	fmt.Println(EqualIgnoreOrder2D(s1, s4)) // true
	fmt.Println(EqualIgnoreOrder2D(s1, s5)) // true
}

func TestSubtract2D(t *testing.T) {
	s1 := [][]string{{"A", "B"}, {"C", "D", "E"}}
	s2 := [][]string{{"A", "B"}, {"F", "G"}, {"H", "I"}}
	fmt.Println(Subtract2D(s1, s2)) // [[C D E]]
	fmt.Println(Subtract2D(s2, s1)) // [[F G] [H I]]
}
