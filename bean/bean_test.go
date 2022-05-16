package bean

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestCopy(t *testing.T) {
	a1 := FruitA{
		ID:         100,
		Name:       "名称100",
		Age:        rand.Intn(100),
		Price:      rand.ExpFloat64(),
		CreateTime: time.Now(),
	}
	a2 := &a1

	var b1 FruitB
	Copy(a1, &b1)
	fmt.Println(b1)

	var b2 *FruitB
	Copy(a1, &b2)
	fmt.Println(b2)

	var b3 FruitB
	Copy(a2, &b3)
	fmt.Println(b3)

	var b4 *FruitB
	Copy(a2, &b4)
	fmt.Println(b4)

	b5 := FruitB{}
	Copy(a1, &b5)
	fmt.Println(b5)

	b6 := &FruitB{}
	Copy(&a1, &b6)
	fmt.Println(b6)

	// 空源数据测试
	var a3 *FruitA
	b7 := &FruitB{}
	err := Copy(a3, &b7)
	println(err)
	fmt.Println(b7)
}

func Test_copyField(t *testing.T) {
	a1 := FruitA{
		ID:         100,
		Name:       "名称100",
		Age:        rand.Intn(100),
		Price:      rand.ExpFloat64(),
		CreateTime: time.Now(),
	}
	//a2 := &a1
	var b1 FruitB
	copyField(a1, &b1)
	copyField(a1, &b1, "Name", "Age")

	fmt.Println(b1)
}

func TestCopySliceBySourceValue(t *testing.T) {
	a1 := getFruitsA(10)

	var b1 []FruitB
	Copy(a1, &b1)
	fmt.Println("---------------")
	fmt.Println("b1:", len(b1), b1)

	var b2 []*FruitB
	Copy(a1, &b2)
	fmt.Println("---------------")
	fmt.Println("b2:", len(b2), b2)

	var b3 []**FruitB
	Copy(a1, &b3)
	fmt.Println("---------------")
	fmt.Println("b3:", len(b3), b3)

	var b4 = make([]FruitB, 0, 0)
	Copy(a1, &b4)
	fmt.Println("---------------")
	fmt.Println("b4:", len(b4), b4)

	var b5 = make([]*FruitB, 0, 0)
	Copy(a1, &b5)
	fmt.Println("---------------")
	fmt.Println("b5:", len(b5), b5)

	var b6 = make([]**FruitB, 0, 0)
	Copy(a1, &b6)
	fmt.Println("---------------")
	fmt.Println("b6:", len(b6), b6)

	var b7 = make([]**FruitB, 0, 0)
	Copy(a1, b7)
	fmt.Println("---------------")
	fmt.Println("b7:", len(b7), b7) // []

	var b11 = new([]FruitB)
	Copy(a1, b11)
	fmt.Println("---------------")
	fmt.Println("b11:", b11)

	var b12 = new([]FruitB)
	Copy(a1, &b12)
	fmt.Println("---------------")
	fmt.Println("b12:", b12)
}

func TestCopySliceByPointer(t *testing.T) {
	a1 := getFruitsPointerA(10)

	var b1 []FruitB
	Copy(a1, &b1)
	fmt.Println("---------------")
	fmt.Println("b1:", len(b1), b1)

	var b2 []*FruitB
	Copy(a1, &b2)
	fmt.Println("---------------")
	fmt.Println("b2:", len(b2), b2)

	var b3 []**FruitB
	Copy(a1, &b3)
	fmt.Println("---------------")
	fmt.Println("b3:", len(b3), b3)

	var b4 = make([]FruitB, 0, 0)
	Copy(a1, &b4)
	fmt.Println("---------------")
	fmt.Println("b4:", len(b4), b4)

	var b5 = make([]*FruitB, 0, 0)
	Copy(a1, &b5)
	fmt.Println("---------------")
	fmt.Println("b5:", len(b5), b5)

	var b6 = make([]**FruitB, 0, 0)
	Copy(a1, &b6)
	fmt.Println("---------------")
	fmt.Println("b6:", len(b6), b6)

	var b7 = make([]**FruitB, 0, 0)
	Copy(a1, b7)
	fmt.Println("---------------")
	fmt.Println("b7:", len(b7), b7) // []

	var b11 = new([]FruitB)
	Copy(a1, b11)
	fmt.Println("---------------")
	fmt.Println("b11:", b11)

	var b12 = new([]FruitB)
	Copy(a1, &b12)
	fmt.Println("---------------")
	fmt.Println("b12:", b12)
}

// TestCopySliceTime 性能测试
func TestCopySliceTime(t *testing.T) {
	//a := getFruitsPointerA(3000000)
	//var b []*FruitB
	a := getFruitsA(3000000)
	var b []FruitB
	err := Copy(a, &b)
	fmt.Println(err)
	fmt.Println(len(b))
}

type FruitA struct {
	ID         uint
	Name       string
	Age        int
	Price      float64
	CreateTime time.Time
}

type FruitB struct {
	ID         uint
	Name       string
	Age        int
	Price      float64
	CreateTime time.Time
	UpdateTime time.Time
	Xss        string
}

func getFruitsPointerA(len int) []*FruitA {
	var data []*FruitA
	for i := 0; i < len; i++ {
		data = append(data, &FruitA{
			ID:         uint(i + 1),
			Name:       "名称" + strconv.Itoa(i+1),
			Age:        rand.Intn(100),
			Price:      rand.ExpFloat64(),
			CreateTime: time.Now(),
		})
	}
	return data
}

func getFruitsA(len int) []FruitA {
	var data []FruitA
	for i := 0; i < len; i++ {
		data = append(data, FruitA{
			ID:         uint(i + 101),
			Name:       "名称" + strconv.Itoa(i+1),
			Age:        rand.Intn(100),
			Price:      rand.ExpFloat64(),
			CreateTime: time.Now(),
		})
	}
	return data
}
