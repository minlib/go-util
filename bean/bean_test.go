package bean

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestCopyToObject(t *testing.T) {
	source1 := FruitA{
		ID:         100,
		Name:       "名称100",
		Age:        rand.Intn(100),
		Price:      rand.ExpFloat64(),
		CreateTime: time.Now(),
	}
	source2 := &source1

	target1, err := CopyTo[FruitB](source1)
	fmt.Println("target1", target1, err)

	target2, err := CopyTo[*FruitB](source1)
	fmt.Println("target2", target2, err)

	target3, err := CopyTo[FruitB](source2)
	fmt.Println("target3", target3, err)

	target4, err := CopyTo[*FruitB](source2)
	fmt.Println("target4", target4, err)

}

// 测试切片复制
func TestCopyToSlice(t *testing.T) {
	source1 := getFruitsA(10)
	//source1 := getFruitsPointerA(10)

	target1, err := CopyTo[[]FruitB](source1)
	fmt.Println("target1:", len(target1), target1, err)

	target2, err := CopyTo[[]*FruitB](source1)
	fmt.Println("target2:", len(target2), target2, err)

	target3, err := CopyTo[[]**FruitB](source1)
	fmt.Println("target3:", len(target3), target3, err)

}

// 测试对象复制
func TestCopyObject(t *testing.T) {
	source1 := FruitA{
		ID:         100,
		Name:       "名称100",
		Age:        rand.Intn(100),
		Price:      rand.ExpFloat64(),
		CreateTime: time.Now(),
	}
	source2 := &source1

	var target1 FruitB
	err1 := Copy(source1, &target1)
	fmt.Println("target1:", target1, err1)

	var target2 *FruitB
	err2 := Copy(source1, &target2)
	fmt.Println("target2:", target2, err2)

	var target3 FruitB
	err3 := Copy(source2, &target3)
	fmt.Println("target3:", target3, err3)

	var target4 *FruitB
	err4 := Copy(source2, &target4)
	fmt.Println("target4:", target4, err4)

	target5 := FruitB{}
	err5 := Copy(source1, &target5)
	fmt.Println("target5:", target5, err5)

	target6 := &FruitB{}
	err6 := Copy(&source1, &target6)
	fmt.Println("target6:", target6, err6)

	// 空源数据测试
	var obj3 *FruitA
	target10 := &FruitB{}
	err10 := Copy(obj3, &target10)
	fmt.Println("target10:", target10, err10)

	var target21 FruitB
	err21 := copyObj(source1, &target21, "Name", "Age")
	fmt.Println("target21:", target21, err21)

	var target22 *FruitB
	err22 := copyObj(source1, &target22, "Name", "Age")
	fmt.Println("target22:", target22, err22)

	// 没有的属性
	var target23 FruitB
	err23 := copyObj(source1, &target23, "Name2", "Age")
	fmt.Println("target23:", target23, err23)
}

// 测试切片复制
func TestCopySlice(t *testing.T) {
	source1 := getFruitsA(10)
	//source1 := getFruitsPointerA(10)

	var target1 []FruitB
	err1 := Copy(source1, &target1)
	fmt.Println("target1:", len(target1), target1, err1)

	var target2 []*FruitB
	err2 := Copy(source1, &target2)
	fmt.Println("target2:", len(target2), target2, err2)

	var target3 []**FruitB
	err3 := Copy(source1, &target3)
	fmt.Println("target3:", len(target3), target3, err3)

	var target4 = make([]FruitB, 0)
	err4 := Copy(source1, &target4)
	fmt.Println("target4:", len(target4), target4, err4)

	var target5 = make([]*FruitB, 0)
	err5 := Copy(source1, &target5)
	fmt.Println("target5:", len(target5), target5, err5)

	var target6 = make([]**FruitB, 0)
	err6 := Copy(source1, &target6)
	fmt.Println("target6:", len(target6), target6, err6)

	var target7 = make([]**FruitB, 0)
	err7 := Copy(source1, target7)
	fmt.Println("target7:", len(target7), target7, err7) // []

	var target11 = new([]FruitB)
	err11 := Copy(source1, target11)
	fmt.Println("target11:", target11, err11)

	var target12 = new([]FruitB)
	err12 := Copy(source1, &target12)
	fmt.Println("target12:", target12, err12)

	var target21 = make([]*FruitB, 0)
	err21 := Copy(source1, &target21, "Name", "Age")
	fmt.Println("target21:", len(target21), target21, err21)
}

// 测试百万级切片复制
func TestCopySliceTime(t *testing.T) {
	source1 := getFruitsPointerA(3000000)
	//a := getFruitsA(3000000)
	var target1 []*FruitB
	//var b []FruitB
	err := Copy(source1, &target1)
	//err := Copy(a, &b, "Name", "Age")
	fmt.Println(len(target1), target1, err)
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
