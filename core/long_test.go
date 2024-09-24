package core

import (
	"encoding/json"
	"fmt"
	"testing"
)

type model struct {
	Id Long
}

func TestLong_MarshalJSON(t *testing.T) {
	model1 := &model{
		Id: NewLong(1),
	}
	byte1, _ := json.Marshal(model1)
	fmt.Println(string(byte1))
	if string(byte1) != `{"Id":"1"}` {
		t.Errorf("Marshal got = %v, want %v", string(byte1), `{"Id":"1"}`)
	}

	model2 := &model{
		Id: NewLong(0),
	}
	byte2, _ := json.Marshal(model2)
	fmt.Println(string(byte2))
	if string(byte2) != `{"Id":"0"}` {
		t.Errorf("Marshal got = %v, want %v", string(byte2), `{"Id":"0"}`)
	}

	model3 := &model{}
	byte3, _ := json.Marshal(model3)
	fmt.Println(string(byte3))
	if string(byte3) != `{"Id":null}` {
		t.Errorf("Marshal got = %v, want %v", string(byte3), `{"Id":null}`)
	}

	model4 := model{}
	byte4, _ := json.Marshal(model4)
	fmt.Println(string(byte4))
	if string(byte4) != `{"Id":null}` {
		t.Errorf("Marshal got = %v, want %v", string(byte4), `{"Id":null}`)
	}

	var model5 model
	byte5, _ := json.Marshal(model5)
	fmt.Println(string(byte5))
	if string(byte5) != `{"Id":null}` {
		t.Errorf("Marshal got = %v, want %v", string(byte5), `{"Id":null}`)
	}

}

func TestLong_UnmarshalJSON(t *testing.T) {
	json1 := `{"Id":"0"}`
	model1 := &model{}
	_ = json.Unmarshal([]byte(json1), &model1)
	if model1.Id.Int64 == nil || *model1.Id.Int64 != 0 {
		t.Errorf("Unmarshal got = %v, want %v", *model1.Id.Int64, 0)
	}

	json2 := `{"Id":"2"}`
	model2 := &model{}
	_ = json.Unmarshal([]byte(json2), &model2)
	if model2.Id.Int64 == nil || *model2.Id.Int64 != 2 {
		t.Errorf("Unmarshal got = %v, want %v", *model2.Id.Int64, 2)
	}

	json3 := `{"Id":0}`
	model3 := &model{}
	_ = json.Unmarshal([]byte(json3), &model3)
	if model3.Id.Int64 == nil || *model3.Id.Int64 != 0 {
		t.Errorf("Unmarshal got = %v, want %v", *model3.Id.Int64, 0)
	}

	json4 := `{"Id":4}`
	model4 := &model{}
	_ = json.Unmarshal([]byte(json4), &model4)
	if model4.Id.Int64 == nil || *model4.Id.Int64 != 4 {
		t.Errorf("Unmarshal got = %v, want %v", *model4.Id.Int64, 4)
	}

	json5 := `{"Id":""}`
	model5 := &model{}
	_ = json.Unmarshal([]byte(json5), &model5)
	if model5.Id.Int64 != nil {
		t.Errorf("Unmarshal got = %v, want %v", model5.Id.Int64, nil)
	}

	json6 := `{"Id":null}`
	model6 := &model{}
	_ = json.Unmarshal([]byte(json6), &model6)
	if model6.Id.Int64 != nil {
		t.Errorf("Unmarshal got = %v, want %v", model6.Id.Int64, nil)
	}

	json7 := `{"Id":"null"}`
	model7 := &model{}
	_ = json.Unmarshal([]byte(json7), &model7)
	if model7.Id.Int64 != nil {
		t.Errorf("Unmarshal got = %v, want %v", model7.Id.Int64, nil)
	}

	json8 := `{"Age":123}`
	model8 := &model{}
	_ = json.Unmarshal([]byte(json8), &model8)
	if model8.Id.Int64 != nil {
		t.Errorf("Unmarshal got = %v, want %v", model8.Id.Int64, nil)
	}

}

func TestLongSlice_MarshalJSON(t *testing.T) {
	s1 := []Long{{nil}, NewLong(0), NewLong(11111), NewLong(22222)}
	byte1, _ := json.Marshal(&s1)
	fmt.Println(string(byte1))
	want1 := `[null,"0","11111","22222"]`
	if string(byte1) != want1 {
		t.Errorf("Unmarshal got = %v, want %v", string(byte1), want1)
	}
}

func TestLongSlice_UnmarshalJSON(t *testing.T) {

	json1 := `["",null,"null","0","100"]`
	var s1 []Long
	_ = json.Unmarshal([]byte(json1), &s1)
	fmt.Println(s1)
	if s1[0].Int64 != nil {
		t.Errorf("Unmarshal got = %v, want %v", s1[0].Int64, nil)
	}
	if s1[1].Int64 != nil {
		t.Errorf("Unmarshal got = %v, want %v", s1[1].Int64, nil)
	}
	if s1[2].Int64 != nil {
		t.Errorf("Unmarshal got = %v, want %v", s1[2].Int64, nil)
	}
	if s1[3].Int64 == nil {
		t.Errorf("Unmarshal got = %v, want %v", s1[3].Int64, nil)
	}
	if *s1[3].Int64 != 0 {
		t.Errorf("Unmarshal got = %v, want %v", *s1[3].Int64, 0)
	}
	if s1[4].Int64 == nil {
		t.Errorf("Unmarshal got = %v, want %v", s1[4].Int64, nil)
	}
	if *s1[4].Int64 != 100 {
		t.Errorf("Unmarshal got = %v, want %v", *s1[4].Int64, 100)
	}
	byte1, _ := json.Marshal(&s1)
	fmt.Println(string(byte1))
}
