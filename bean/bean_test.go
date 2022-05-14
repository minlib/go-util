package bean

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestCopy(t *testing.T) {
	a := FruitA{
		ID:         100,
		Name:       "名称100",
		Age:        rand.Intn(100),
		Price:      rand.ExpFloat64(),
		CreateTime: time.Now(),
	}
	//var b = new(FruitB)
	var b = FruitB{}
	copyObj(a, b)
	fmt.Println(b)

	b = FruitB{}
	copyObj(&a, b)
	fmt.Println(b)

	b = FruitB{}
	copyObj(a, &b)
	fmt.Println(b)

	b = FruitB{}
	copyObj(&a, &b)
	fmt.Println(b)

}

func TestCopySlice(t *testing.T) {
	a1 := getFruitsA(10)
	a2 := getFruitsPointerA(10)

	var b1 []FruitB
	Copy(a1, &b1)
	fmt.Println("b1", b1)

	var b2 []*FruitB
	Copy(a1, b2)
	fmt.Println("b2", b2)

	var b3 = new([]FruitB)
	Copy(a1, &b3)
	fmt.Println("b3", b3)

	var b4 []FruitB
	Copy(a2, &b4)
	fmt.Println("b4", b4)

	var b5 []*FruitB
	Copy(a2, b5)
	fmt.Println("b5", b5)

	var b6 = make([]FruitB, 0, 0)
	Copy(a2, &b6)
	fmt.Println("b6", b6)

}

/**
300万条切片中值与地址传递的耗时对比
[]FruitsA
12.29s
12.46s
12.18s
11.64s

[]*FruitsA
12.11s
12.88s
12.67s
11.92s
*/
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
