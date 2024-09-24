package core

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestDateTime_Scan(t1 *testing.T) {
	// 从数据库中获取的 UTC 时间
	loc, _ := time.LoadLocation("Asia/Shanghai")

	utcTime := time.Date(2024, 6, 27, 15, 30, 0, 0, time.UTC)

	// 将 UTC 时间转换为东八区时间
	cstTime := utcTime.In(loc)
	fmt.Println(cstTime)

	x := testModel{
		T: NewDateTime(time.Now()),
	}
	marshal, _ := json.Marshal(&x)
	fmt.Println(string(marshal))

	var y testModel
	_ = json.Unmarshal(marshal, &y)
	fmt.Println(y)
}

type testModel struct {
	T *DateTime
}
