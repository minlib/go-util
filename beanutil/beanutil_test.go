package beanutil

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestCopy(t *testing.T) {
	fruits := getFruits()
	var result []*fruitB
	for _, v := range fruits {
		var item fruitB
		Copy(&item, v)
		result = append(result, &item)
	}
	fmt.Println(result)
}

type fruitA struct {
	ID         uint
	Name       string
	Age        int
	Price      float64
	CreateTime time.Time
}

type fruitB struct {
	ID         uint
	Name       string
	Age        int
	Price      float64
	CreateTime time.Time
	UpdateTime time.Time
	Xss        string
}

func getFruits() []*fruitA {
	var data []*fruitA
	for i := 0; i < 10; i++ {
		data = append(data, &fruitA{
			ID:         uint(i + 1),
			Name:       "名称" + strconv.Itoa(i+1),
			Age:        rand.Intn(100),
			Price:      rand.ExpFloat64(),
			CreateTime: time.Now(),
		})
	}
	return data
}
