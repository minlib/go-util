package slicex

import (
	"encoding/json"
	"fmt"
	"github.com/minlib/go-util/jsonx"
	"reflect"
	"sort"
	"testing"
)

func TestSlice(t *testing.T) {
	fmt.Println(Slice("A1", "A2", "A3", "A4", "A5", "A6"))
}

func TestDelete(t *testing.T) {
	fmt.Println(Delete([]int{1, 2, 3, 4, 5, 6}, 1))
	fmt.Println(Delete([]string{"A1", "A2", "A3", "A4", "A5", "A6"}, 2))
}

func TestSubtract(t *testing.T) {
	intSlice1 := []int{1, 2, 2, 2, 3, 3, 4, 5, 6}
	intSlice2 := []int{2, 1, 4}
	fmt.Println(Subtract(intSlice1, intSlice2))

	stringSlice1 := []string{"A1", "A2", "A3", "A7"}
	stringSlice2 := []string{"A1", "A2", "A4", "A5", "A6"}
	fmt.Println(Subtract(stringSlice1, stringSlice2))
}

func TestSubtractDistinct(t *testing.T) {
	intSlice1 := []int{1, 2, 2, 2, 3, 3, 4, 5, 6}
	intSlice2 := []int{2, 1, 4}
	fmt.Println(SubtractDistinct(intSlice1, intSlice2))

	stringSlice1 := []string{"A1", "A2", "A3", "A7"}
	stringSlice2 := []string{"A1", "A2", "A4", "A5", "A6"}
	fmt.Println(SubtractDistinct(stringSlice1, stringSlice2))
}

func TestIntersect(t *testing.T) {
	intSlice1 := []int{1, 2, 2, 2, 3, 4, 6}
	intSlice2 := []int{2, 1, 4}
	fmt.Println(Intersect(intSlice1, intSlice2))

	stringSlice1 := []string{"A1", "A2", "A3", "A7"}
	stringSlice2 := []string{"A1", "A2", "A4", "A5", "A6"}
	fmt.Println(Intersect(stringSlice1, stringSlice2))
}

func TestUnion(t *testing.T) {
	intSlice1 := []int{1, 2, 2, 2, 6, 4, 8, 3}
	intSlice2 := []int{2, 2, 1, 4}
	fmt.Println(Union(intSlice1, intSlice2))

	stringSlice1 := []string{"A1", "A2", "A3", "A7"}
	stringSlice2 := []string{"A1", "A2", "A4", "A5", "A6"}
	fmt.Println(Union(stringSlice1, stringSlice2))
}

func TestDistinct(t *testing.T) {
	intSlice1 := []int64{1, 2, 3, 4, 4, 1, 5, 6}
	fmt.Println(Distinct(intSlice1))

	intSlice2 := []int32{1, 3, 4, 4, 1, 5, 6}
	fmt.Println(Distinct(intSlice2))

	stringSlice := []string{"A1", "A2", "A3", "A4", "A5", "A6", "A3", "A4", "A5"}
	fmt.Println(Distinct(stringSlice))
}

func TestSum(t *testing.T) {
	intSlice1 := []int64{1, 2, 3}
	fmt.Println(Sum(intSlice1))

	intSlice2 := []int{1, 2, 3}
	fmt.Println(Sum(intSlice2))

	intSlice3 := []int32{1, 2, 3}
	fmt.Println(Sum(intSlice3))

	floatSlice4 := []float64{1.1, 2.2, 3.3}
	fmt.Println(Sum(floatSlice4))

	stringSlice5 := []string{"A1", "A2", "A3"}
	fmt.Println(Sum(stringSlice5))
}

func TestEqual(t *testing.T) {
	fmt.Println(Equal([]int{1, 2, 3}, []int{3, 1, 2}))
	fmt.Println(EqualIgnoreOrder([]int{1, 2, 3}, []int{3, 1, 2}))
	fmt.Println(Equal([]string{"A1", "A2", "A3"}, []string{"A1", "A3", "A2"}))
	fmt.Println(EqualIgnoreOrder([]string{"A1", "A2", "A3"}, []string{"A1", "A3", "A2"}))
}

type long int64

func TestIntToString(t *testing.T) {
	fmt.Println(IntToString([]int{1, 2, 3}))
	fmt.Println(IntToString([]int8{4, 5, 6}))
	fmt.Println(IntToString([]int16{7, 8, 9}))
	fmt.Println(IntToString([]int32{11, 12, 13}))
	fmt.Println(IntToString([]int64{14, 15, 16}))
	fmt.Println(IntToString([]uint{21, 22, 23}))
	fmt.Println(IntToString([]uint32{31, 32, 33}))
	fmt.Println(IntToString([]uint64{34, 35, 36}))
	fmt.Println(IntToString([]long{41, 42, 43}))
}

func TestStringToInt(t *testing.T) {
	fmt.Println(StringToInt[int]([]string{"1", "2", "3"}))
	fmt.Println(StringToInt[int8]([]string{"4", "5", "6"}))
	fmt.Println(StringToInt[int16]([]string{"7", "8", "9"}))
	fmt.Println(StringToInt[int32]([]string{"11", "12", "13"}))
	fmt.Println(StringToInt[int64]([]string{"14", "15", "16"}))
	fmt.Println(StringToInt[uint]([]string{"17", "18", "19"}))
}

type Model struct {
	id   int
	name string
}

func TestSliceSortFunc(t *testing.T) {
	var datas []Model
	datas = append(datas, Model{1, "a1"})
	datas = append(datas, Model{4, "a4"})
	datas = append(datas, Model{3, "a3"})
	datas = append(datas, Model{2, "a2"})
	datas = append(datas, Model{2, "a1"})
	sort.Slice(datas, func(i, j int) bool {
		return datas[i].id < datas[j].id
	})
	fmt.Println(datas) // [{1 a1} {2 a2} {3 a3} {4 a4}]
}

func TestFunc(t *testing.T) {
	// stringSlice := []string{"A1", "A3", "A6", "A2", "A5", "A4"}
	// stringSlice2 := []string{"A1", "A3", "A6", "A2", "A4", "A5"}
	// slices.Sort(stringSlice)
	// fmt.Println(stringSlice)

	// float321 := float32(1.0)
	// fmt.Println(float321)
	// str := fmt.Sprintf("%f", float321)
	// fmt.Println(str)

	// fmt.Println(Slice("aaa", "stringSlice", "xx"))
	// var x1 []string = Slice("A")
	// var x2 []string = Slice("A", "B")
	// fmt.Println(x1)
	// fmt.Println(x2)

	// Join(x2)

	// intSlice1 := []int64{1, 2, 3}
	// stringSlice := []string{"A1", "A2", "A3", "A4", "A5", "A6"}

	// fmt.Println(slices.Contains(intSlice1, 2))
	// fmt.Println(slices.Contains(intSlice1, 13))
	// fmt.Println(slices.Contains(stringSlice, "A4"))
	// fmt.Println(slices.Equal(intSlice1, intSlice2))
	// fmt.Println(slices.Equal(intSlice1, intSlice2))
	// intSlice := []int{1, 2, 3, 4, 5, 6}
	// fmt.Println(Delete(intSlice, 1))
	// fmt.Println(stringSlice)
	// fmt.Println(Delete(stringSlice, 2))
	// fmt.Println(Delete(stringSlice, 0))
	// fmt.Println(stringSlice)
}

func TestLongToInt64(t *testing.T) {
	type args struct {
		s []jsonx.Long
	}
	var args1 []jsonx.Long
	json.Unmarshal([]byte(`["",null,"null","0","100"]`), &args1)
	want1 := []int64{0, 100}
	got1 := LongToInt64(args1)
	if !reflect.DeepEqual(got1, want1) {
		t.Errorf("LongToInt64() = %v, want %v", got1, want1)
	}
}
