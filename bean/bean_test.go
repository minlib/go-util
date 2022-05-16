package bean

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

// 测试对象复制
func TestCopyObject(t *testing.T) {
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
	b10 := &FruitB{}
	err10 := Copy(a3, &b10)
	fmt.Println(err10)
	fmt.Println(b10)

	var b21 FruitB
	copyObj(a1, &b21, "Name", "Age")
	fmt.Println(b21)

	var b22 *FruitB
	copyObj(a1, &b22, "Name", "Age")
	fmt.Println(b22)

	// 没有的属性
	var b23 FruitB
	err23 := copyObj(a1, &b23, "Name2", "Age")
	fmt.Println(err23)
	fmt.Println(b23)
}

// 测试切片复制
func TestCopySlice(t *testing.T) {
	a1 := getFruitsA(10)
	//a1 := getFruitsPointerA(10)

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

	var b21 = make([]*FruitB, 0, 0)
	Copy(a1, &b21, "Name", "Age")
	fmt.Println("---------------")
	fmt.Println("b21:", len(b21), b21)
}

// 测试百万级切片复制
func TestCopySliceTime(t *testing.T) {
	a := getFruitsPointerA(1000000)
	//a := getFruitsA(3000000)
	var b []*FruitB
	//var b []FruitB
	err := Copy(a, &b)
	//err := Copy(a, &b, "Name", "Age")
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
	Address    string
	CreateTime time.Time
	UpdateTime time.Time
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
