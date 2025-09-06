package core

import (
	"encoding/json"
	"fmt"
	"testing"
)

type intModel struct {
	Id Integer
}

func TestInteger_MarshalJSON(t *testing.T) {
	intModel1 := &intModel{
		Id: NewInteger(1),
	}
	byte1, _ := json.Marshal(intModel1)
	fmt.Println(string(byte1))
	if string(byte1) != `{"Id":1}` {
		t.Errorf("Marshal got = %v, want %v", string(byte1), `{"Id":1}`)
	}

	intModel2 := &intModel{
		Id: NewInteger(0),
	}
	byte2, _ := json.Marshal(intModel2)
	fmt.Println(string(byte2))
	if string(byte2) != `{"Id":0}` {
		t.Errorf("Marshal got = %v, want %v", string(byte2), `{"Id":0}`)
	}

	intModel3 := &intModel{}
	byte3, _ := json.Marshal(intModel3)
	fmt.Println(string(byte3))
	if string(byte3) != `{"Id":null}` {
		t.Errorf("Marshal got = %v, want %v", string(byte3), `{"Id":null}`)
	}

	intModel4 := intModel{}
	byte4, _ := json.Marshal(intModel4)
	fmt.Println(string(byte4))
	if string(byte4) != `{"Id":null}` {
		t.Errorf("Marshal got = %v, want %v", string(byte4), `{"Id":null}`)
	}

	var intModel5 intModel
	byte5, _ := json.Marshal(intModel5)
	fmt.Println(string(byte5))
	if string(byte5) != `{"Id":null}` {
		t.Errorf("Marshal got = %v, want %v", string(byte5), `{"Id":null}`)
	}

}

func TestInteger_UnmarshalJSON(t *testing.T) {
	json1 := `{"Id":"0"}`
	intModel1 := &intModel{}
	json.Unmarshal([]byte(json1), &intModel1)
	if intModel1.Id.Int32 == nil || *intModel1.Id.Int32 != 0 {
		t.Errorf("Unmarshal got = %v, want %v", *intModel1.Id.Int32, 0)
	}

	json2 := `{"Id":"2"}`
	intModel2 := &intModel{}
	json.Unmarshal([]byte(json2), &intModel2)
	if intModel2.Id.Int32 == nil || *intModel2.Id.Int32 != 2 {
		t.Errorf("Unmarshal got = %v, want %v", *intModel2.Id.Int32, 2)
	}

	json3 := `{"Id":0}`
	intModel3 := &intModel{}
	json.Unmarshal([]byte(json3), &intModel3)
	if intModel3.Id.Int32 == nil || *intModel3.Id.Int32 != 0 {
		t.Errorf("Unmarshal got = %v, want %v", *intModel3.Id.Int32, 0)
	}

	json4 := `{"Id":4}`
	intModel4 := &intModel{}
	json.Unmarshal([]byte(json4), &intModel4)
	if intModel4.Id.Int32 == nil || *intModel4.Id.Int32 != 4 {
		t.Errorf("Unmarshal got = %v, want %v", *intModel4.Id.Int32, 4)
	}

	json5 := `{"Id":""}`
	intModel5 := &intModel{}
	json.Unmarshal([]byte(json5), &intModel5)
	if intModel5.Id.Int32 != nil {
		t.Errorf("Unmarshal got = %v, want %v", intModel5.Id.Int32, nil)
	}

	json6 := `{"Id":null}`
	intModel6 := &intModel{}
	json.Unmarshal([]byte(json6), &intModel6)
	if intModel6.Id.Int32 != nil {
		t.Errorf("Unmarshal got = %v, want %v", intModel6.Id.Int32, nil)
	}

	json7 := `{"Id":"null"}`
	intModel7 := &intModel{}
	json.Unmarshal([]byte(json7), &intModel7)
	if intModel7.Id.Int32 != nil {
		t.Errorf("Unmarshal got = %v, want %v", intModel7.Id.Int32, nil)
	}

	json8 := `{"Age":123}`
	intModel8 := &intModel{}
	json.Unmarshal([]byte(json8), &intModel8)
	if intModel8.Id.Int32 != nil {
		t.Errorf("Unmarshal got = %v, want %v", intModel8.Id.Int32, nil)
	}

}

func TestIntegerSlice_MarshalJSON(t *testing.T) {
	s1 := []Integer{{nil}, NewInteger(0), NewInteger(11111), NewInteger(22222)}
	byte1, _ := json.Marshal(&s1)
	fmt.Println(string(byte1))
	want1 := `[null,0,11111,22222]`
	if string(byte1) != want1 {
		t.Errorf("Unmarshal got = %v, want %v", string(byte1), want1)
	}
}

func TestIntegerSlice_UnmarshalJSON(t *testing.T) {
	json1 := `["",null,"null","0","100"]`
	var s1 []Integer
	json.Unmarshal([]byte(json1), &s1)
	fmt.Println(s1)
	if s1[0].Int32 != nil {
		t.Errorf("Unmarshal got = %v, want %v", s1[0].Int32, nil)
	}
	if s1[1].Int32 != nil {
		t.Errorf("Unmarshal got = %v, want %v", s1[1].Int32, nil)
	}
	if s1[2].Int32 != nil {
		t.Errorf("Unmarshal got = %v, want %v", s1[2].Int32, nil)
	}
	if s1[3].Int32 == nil {
		t.Errorf("Unmarshal got = %v, want %v", s1[3].Int32, nil)
	}
	if *s1[3].Int32 != 0 {
		t.Errorf("Unmarshal got = %v, want %v", *s1[3].Int32, 0)
	}
	if s1[4].Int32 == nil {
		t.Errorf("Unmarshal got = %v, want %v", s1[4].Int32, nil)
	}
	if *s1[4].Int32 != 100 {
		t.Errorf("Unmarshal got = %v, want %v", *s1[4].Int32, 100)
	}
	byte1, _ := json.Marshal(&s1)
	fmt.Println(string(byte1))
}
