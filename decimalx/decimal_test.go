package decimalx

import (
	"github.com/shopspring/decimal"
	"testing"
)

func NewFromString(s string) decimal.Decimal {
	d, _ := decimal.NewFromString(s)
	return d
}

func TestGetPlaces(t *testing.T) {
	tests := []struct {
		name  string
		value decimal.Decimal
		want  int32
	}{
		{
			value: NewFromString("0"),
			want:  0,
		},
		{
			value: NewFromString("0.00"),
			want:  0,
		},
		{
			value: NewFromString("5"),
			want:  0,
		},
		{
			value: NewFromString("5."),
			want:  0,
		},
		{
			value: NewFromString("5.0000"),
			want:  0,
		},
		{
			value: NewFromString("0.4"),
			want:  1,
		},
		{
			value: NewFromString("0.05"),
			want:  2,
		},
		{
			value: NewFromString("0.001"),
			want:  3,
		},
		{
			value: NewFromString("0.0006"),
			want:  4,
		},
		{
			value: NewFromString("5.000000001"),
			want:  9,
		},
		{
			value: NewFromString("-5.0000"),
			want:  0,
		},
		{
			value: NewFromString("-0.4"),
			want:  1,
		},
		{
			value: NewFromString("-0.05"),
			want:  2,
		},
		{
			value: NewFromString("-0.001"),
			want:  3,
		},
		{
			value: NewFromString("-0.0006"),
			want:  4,
		},
		{
			value: NewFromString("-5.000000001"),
			want:  9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPlaces(tt.value); got != tt.want {
				t.Errorf("GetPlaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPlaces(t *testing.T) {
	tests := []struct {
		name   string
		value  decimal.Decimal
		places int32
		want   bool
	}{
		{
			value:  NewFromString("0"),
			places: 0,
			want:   true,
		},
		{
			value:  NewFromString("0.00"),
			places: 0,
			want:   true,
		},
		{
			value:  NewFromString("5"),
			places: 0,
			want:   true,
		},
		{
			value:  NewFromString("5."),
			places: 0,
			want:   true,
		},
		{
			value:  NewFromString("5.0000"),
			places: 0,
			want:   true,
		},
		{
			value:  NewFromString("0.4"),
			places: 1,
			want:   true,
		},
		{
			value:  NewFromString("0.05"),
			places: 2,
			want:   true,
		},
		{
			value:  NewFromString("0.001"),
			places: 3,
			want:   true,
		},
		{
			value:  NewFromString("0.0006"),
			places: 4,
			want:   true,
		},
		{
			value:  NewFromString("5.000000001"),
			places: 9,
			want:   true,
		},
		{
			value:  NewFromString("-5.0000"),
			places: 0,
			want:   true,
		},
		{
			value:  NewFromString("-0.4"),
			places: 1,
			want:   true,
		},
		{
			value:  NewFromString("-0.05"),
			places: 2,
			want:   true,
		},
		{
			value:  NewFromString("-0.001"),
			places: 3,
			want:   true,
		},
		{
			value:  NewFromString("-0.0006"),
			places: 4,
			want:   true,
		},
		{
			value:  NewFromString("-5.000000001"),
			places: 9,
			want:   true,
		},
		{
			value:  NewFromString("-5.000000001"),
			places: 10,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPlaces(tt.value, tt.places); got != tt.want {
				t.Errorf("GetPlaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
