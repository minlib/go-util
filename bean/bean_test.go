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
	fruits := getFruits()
	var result []*FruitB
	for _, v := range fruits {
		var item FruitB
		Copy(&item, v)
		result = append(result, &item)
	}
	fmt.Println(result)
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

func getFruits() []*FruitA {
	var data []*FruitA
	for i := 0; i < 10; i++ {
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

func TestCopyList(t *testing.T) {

	//var x float64 = 3.4
	//v := reflect.ValueOf(&x)
	//tx := reflect.TypeOf(&x)
	//fmt.Println(tx)                // *float64
	//fmt.Println(v.Elem().CanSet()) // true

	//var x = new(FruitB)
	//v := reflect.ValueOf(&x)
	//tx := reflect.TypeOf(&x)
	//fmt.Println(tx)                       // **bean.FruitB
	//fmt.Println(v.CanSet())               // false
	//fmt.Println(v.Elem().CanSet())        // true
	//fmt.Println(v.Elem().Elem().CanSet()) // true

	//x := FruitB{}
	//v := reflect.ValueOf(x)
	//tx := reflect.TypeOf(x)
	//fmt.Println(tx)         // bean.FruitB
	//fmt.Println(v.CanSet()) // false

	//var x []*FruitB
	//v := reflect.ValueOf(&x)
	//tx := reflect.TypeOf(&x)
	//fmt.Println(tx)                // *[]*bean.FruitB
	//fmt.Println(v.Elem().CanSet()) // true

	var x []*FruitB
	tx := reflect.TypeOf(x) // []*bean.FruitB
	v := reflect.ValueOf(x) // []
	fmt.Println(tx)         // []*bean.FruitB
	fmt.Println(v)          // []
	fmt.Println(v.CanSet()) // false

	//var x []*FruitB
	//v := reflect.ValueOf(x) // []
	//tx := reflect.TypeOf(x) // []*bean.FruitB
	//fmt.Println(tx)         // []*bean.FruitB
	//fmt.Println(v)          // []
	//fmt.Println(v.CanSet()) // false
	//fmt.Println(v.Elem().CanSet()) // false

	fruits := getFruits()
	var source []*FruitA
	for _, v := range fruits {
		source = append(source, v)
	}
	var source2 []*FruitB

	Copy2(&source, &source2)
	Copy2(source, &source2)

	//var source2 *FruitA
	//sourceType3 := reflect.TypeOf(source2)
	//fmt.Println(sourceType3.Kind())

	//var source2 [5]*FruitA
	//sourceType3 := reflect.TypeOf(source2)
	//fmt.Println(sourceType3.Elem().Elem().Kind())

	sourceType := reflect.TypeOf(source)
	fmt.Println(sourceType)                             // []*bean.FruitA
	fmt.Println(sourceType.Kind())                      // slice
	fmt.Println(sourceType.Elem())                      // *bean.FruitA
	fmt.Println(sourceType.Elem().Kind())               // ptr
	fmt.Println(sourceType.Elem().Elem())               // bean.FruitA
	fmt.Println(sourceType.Elem().Elem().Kind())        // struct
	fmt.Println(sourceType.Elem().Elem().Field(0).Name) // ID

	i1 := &FruitA{}
	i2 := &i1
	fmt.Println(reflect.ValueOf(i2))               // 0xc000006048
	fmt.Println(reflect.ValueOf(i2).Elem())        // &{0  0 0 0001-01-01 00:00:00 +0000 UTC}
	fmt.Println(reflect.ValueOf(i2).Elem().Elem()) // {0  0 0 0001-01-01 00:00:00 +0000 UTC}

	//fmt.Println(sourceType.Elem().Field(0).Name)	//

	sourceValue := reflect.ValueOf(source[0]).Elem()
	sourceType2 := sourceValue.Type()
	fmt.Println(sourceValue)               // {1 名称1 81 0.5372820936538049 2022-05-13 10:46:15.6915219 +0800 CST m=+0.002175701}
	fmt.Println(sourceType2)               // bean.FruitA
	fmt.Println(sourceType2.Field(0).Name) // ID

	//CopyList(source, &bs)

	////先获取reflect.Type
	//rTyp := reflect.TypeOf(source)
	//fmt.Println(rTyp) // []*bean.FruitA

	//sliceType := reflect.SliceOf(rTyp)
	//fmt.Println(sliceType)
	//
	////2.获取到reflect.value
	//rVal := reflect.ValueOf(source)
	//fmt.Println("rVal = ", rVal)
	////fmt.Println("rVal.NumField():", rVal.NumField())
	////fmt.Println("rVal.NumField():", len(rVal))
	//
	//sourceValue := reflect.ValueOf(source[0]).Elem()
	//fmt.Println("sourceValue:", sourceValue)
	//fmt.Println("sourceValue.NumField():", sourceValue.NumField())

	//out := sourceType1.In(0)
	//out := sourceType1.Out(0)
	//fmt.Println(out)

	//sliceType := reflect.SliceOf(sourceType1)
	//fmt.Println(sliceType)

	//of := reflect.ArrayOf(0, sourceType1)
	//fmt.Println(of)
	//
	//fmt.Println(sourceType1)
	//fmt.Println(sourceType1.Field(0).Name)

	//sourceValue := reflect.ValueOf(source[0]).Elem()
	//sourceType := sourceValue.Type()
	//fmt.Println(sourceType)
	//fmt.Println(sourceType.Field(0).Name)

	// begin

	// END

	var list []*FruitA
	if reflect.TypeOf(source).Kind() == reflect.Slice {
		s := reflect.ValueOf(source)
		for i := 0; i < s.Len(); i++ {
			ele := s.Index(i)
			fmt.Println(ele)
			fmt.Println(ele.Interface())
			list = append(list, ele.Interface().(*FruitA))
		}
	}

	if sourceType.Kind() == reflect.Slice {
		var list []*FruitA
		if reflect.TypeOf(source).Kind() == reflect.Slice {
			s := reflect.ValueOf(source)
			for i := 0; i < s.Len(); i++ {
				ele := s.Index(i)
				fmt.Println(ele)
				fmt.Println(ele.Interface())
				list = append(list, ele.Interface().(*FruitA))
			}
		}

		sliceType := reflect.SliceOf(reflect.TypeOf(source))
		slice := reflect.MakeSlice(sliceType, 0, 0)
		slicedata := reflect.New(slice.Type())
		slicedata.Elem().Set(slice)
		data := slicedata.Interface()
		fmt.Println(data)
		fmt.Println(slicedata)

		//for i := 0; i < sourceValue.NumField(); i++ {
		//
		//}
	}

}
