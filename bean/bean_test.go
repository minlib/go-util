package bean

import (
	"fmt"
	"math/rand"
	"reflect"
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
	Copy(a, b)
	fmt.Println(b)

	b = FruitB{}
	Copy(&a, b)
	fmt.Println(b)

	b = FruitB{}
	Copy(a, &b)
	fmt.Println(b)

	b = FruitB{}
	Copy(&a, &b)
	fmt.Println(b)

	var b2 FruitB
	fmt.Printf("%p\n", &b)
	copyObj(a, &b2)
	fmt.Println(a)
	fmt.Println(b2)
}

func TestCopyNil(t *testing.T) {
	//var a *FruitA
	a := &FruitA{
		ID:         100,
		Name:       "名称100",
		Age:        rand.Intn(100),
		Price:      rand.ExpFloat64(),
		CreateTime: time.Now(),
	}
	var b FruitB
	//var b = FruitB{}
	fmt.Printf("%p\n", &b)
	copyObj(a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
func TestCopyObj(t *testing.T) {
	a := &FruitB{
		ID:         100,
		Name:       "名称100",
		Age:        rand.Intn(100),
		Price:      rand.ExpFloat64(),
		CreateTime: time.Now(),
	}

	var b *FruitB
	fmt.Println(b)
	fmt.Println(&b)
	copyObjTest(a, &b)
	fmt.Println(b)
}

func Test222(t *testing.T) {
	var slice = []string{"one", "tow", "three"}
	fmt.Println(reflect.ValueOf(slice))      // [one tow three]
	fmt.Println(reflect.TypeOf(slice))       // []string
	fmt.Printf("slice的内存地址 = %p\n", &slice)  //这个是变量t1的地址
	fmt.Printf("slice指向的内存地址 = %p\n", slice) //这个是变量t1对应的切片的地址

	var a *int     // 存储的是int的指针，目前为空
	var b int = 4  // 存储的是int的值
	fmt.Println(a) // <nil>
	a = &b         // a 指向 b 的地址

	fmt.Printf("a的内存地址 = %p\n", &a)  //这个是变量t1的地址
	fmt.Printf("a指向的内存地址 = %p\n", a) //这个是变量t1对应的切片的地址

	//a = b            // a 无法等于 b，会报错，a是指针，b是值，存储的类型不同
	fmt.Println(a)   // a:0xc00000a090(返回了地址)
	fmt.Println(*a)  // *a:4(返回了值)
	fmt.Println(*&a) // *抵消了&，返回了0xc00000a090本身
	*a = 5           // 改变 a 的地址的值
	fmt.Println(b)   // b:5，改变后 b 同样受到改变，因为 a 的地址是指向 b 的

}

func Test333(t *testing.T) {
	b := 255
	var a *int = &b
	fmt.Printf("%T\n", a)  // *int
	fmt.Printf("%p\n", a)  // 0xc00000a300
	fmt.Printf("%p\n", &a) // 指针类型的变量也是有地址的
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
