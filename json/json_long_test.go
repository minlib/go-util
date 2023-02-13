package json

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/shopspring/decimal"
// )

// type TestLongStruct2 struct {
// 	Long1 json.Long
// 	Long2 *json.Long
// }

// type TestLongStruct struct {
// 	Long1 json.Long
// 	Long2 *json.Long

// 	DateTime1 DateTime
// 	DateTime2 *DateTime

// 	Decimal1 decimal.Decimal
// 	Decimal2 *decimal.Decimal

// 	DD *int64
// 	EE *int64
// 	FF int64
// }

// func TestLong_MarshalJSON(t *testing.T) {
// 	testLongStruct2 := TestLongStruct2{}
// 	testLongStruct2.Long1 = json.Long{1, true}
// 	testLongStruct2.Long2 = &json.Long{1, true}
// 	marshal := json.Marshal(testLongStruct2)
// 	fmt.Println(string(marshal))
// }

// func TestLong_MarshalJSON3(t *testing.T) {
// 	//as := &Long{0, true}
// 	testLongStruct2 := TestLongStruct2{}
// 	testLongStruct2.Long1 = json.Long{1, true}
// 	testLongStruct2.Long2 = nil
// 	marshal := json.Marshal(testLongStruct2)
// 	fmt.Println(string(marshal))
// }

// func TestLong_MarshalJSON2(t *testing.T) {
// 	//// {"Long3":{"Int64":0},"Long4":null,"Long1":{"Int64":123},"Long2":"234","Decimal1":"1234555","Decimal2":"1234555","DateTime1":"2022-03-21 20:36:38","DateTime2":"2022-03-21 20:36:38","DD":null,"EE":null,"FF":0}
// 	////var a1 = Long(1000)
// 	////var b1 = &a1
// 	//var t2 = TestLongStruct{}
// 	//
// 	//l2 := Long{234, false}
// 	//t2.Long1 = Long{123, true}
// 	//t2.Long2 = &l2
// 	////t2.Long4 = nil
// 	//
// 	//d4 := DateTime{time.Now()}
// 	//t2.DateTime1 = DateTime{time.Now()}
// 	//t2.DateTime2 = &d4
// 	////t2.DateTime4 = nil
// 	//
// 	//d2 := decimal.NewFromInt(1234555)
// 	//t2.Decimal1 = decimal.NewFromInt(1234555)
// 	//t2.Decimal2 = &d2
// 	////t2.Decimal4 = nil
// 	//
// 	//marshal, err := json.Marshal(t2)
// 	//fmt.Println(err)
// 	//fmt.Println(string(marshal))

// 	//jsonString := "{\"Long1\":\"123\",\"Long2\":\"234\",\"Long3\":\"0\",\"Long4\":null,\"Decimal1\":\"1234555\",\"Decimal2\":\"1234555\",\"DateTime1\":\"2022-03-21 20:59:52\",\"DateTime2\":\"2022-03-21 20:59:52\",\"DD\":null,\"EE\":null,\"FF\":0}"
// 	//jsonString := "{\"Long1\":\"\",\"Long2\":\"null\",\"Long3\":\"0\",\"Long4\":null,\"Decimal1\":\"1234555\",\"Decimal2\":\"1234555\",\"DateTime1\":\"2022-03-21 20:59:52\",\"DateTime2\":\"2022-03-21 20:59:52\",\"DD\":null,\"EE\":null,\"FF\":0}"
// 	//jsonString := "{\"Long1\":\"\",\"Long2\":\"\",\"Decimal1\":\"1234555\",\"Decimal2\":\"1234555\",\"DateTime1\":\"2022-03-21 20:59:52\",\"DateTime2\":\"null\",\"DD\":null,\"EE\":null,\"FF\":0}"
// 	//jsonString := "{\"Long1\":null,\"Long2\":\"\"}"
// 	jsonString := "{\"Long1\":null,\"Long2\":null}"
// 	var result TestLongStruct
// 	json.Unmarshal(jsonString, &result)

// 	if result.Long2 != nil {
// 		fmt.Println(result.Long2.Valid)
// 	}
// 	fmt.Println(result)

// }
