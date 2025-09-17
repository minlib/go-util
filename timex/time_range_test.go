package timex

import (
	"fmt"
	"testing"
	"time"
)

func TestRangeDateTime(t *testing.T) {
	dt, _ := ParseDateTime("2024-10-05 22:10:56")
	fmt.Println(StartDateTime(dt))   // 2022-10-05 00:00:00 +0800 CST
	fmt.Println(EndDateTime(dt))     // 2022-10-05 23:59:59.999999999 +0800 CST
	fmt.Println(StartNanoSecond(dt)) // 2024-10-05 22:10:56 +0800 CST
	fmt.Println(EndNanoSecond(dt))   // 2024-10-05 22:10:56.999999999 +0800 CST
	fmt.Println(RangeDateTime(dt))   // 2022-10-05 00:00:00 +0800 CST  2022-10-05 23:59:59.999999999 +0800 CST
}

func TestGetIntervalString(t *testing.T) {
	fmt.Println(GetIntervalTime(parseTime("2024-08-05 22:00:00"), parseTime("2026-08-15 22:00:00"))) // 17760h0m0s 2年
	fmt.Println(GetIntervalTime(parseTime("2024-08-05 22:00:00"), parseTime("2025-08-05 22:00:00"))) // 8760h0m0s 1年
	fmt.Println(GetIntervalTime(parseTime("2024-08-05 22:00:00"), parseTime("2025-08-04 22:00:00"))) // 8736h0m0s 12个月
	fmt.Println(GetIntervalTime(parseTime("2024-08-05 22:00:00"), parseTime("2025-03-03 21:00:00"))) // 5039h0m0s 6个月
	fmt.Println(GetIntervalTime(parseTime("2024-08-05 22:00:00"), parseTime("2024-08-25 21:00:00"))) // 479h0m0s 19天
	fmt.Println(GetIntervalTime(parseTime("2024-08-05 22:00:00"), parseTime("2024-08-06 21:00:00"))) // 23h0m0s 23小时
	fmt.Println(GetIntervalTime(parseTime("2024-08-05 22:00:00"), parseTime("2024-08-05 22:03:11"))) // 3m11s 3分钟
	fmt.Println(GetIntervalTime(parseTime("2024-08-05 22:00:00"), parseTime("2024-08-05 22:00:11"))) // 11s 11秒
	fmt.Println(GetIntervalTime(parseTime("2024-08-05 22:00:00"), parseTime("2024-08-05 22:00:00"))) // 0s ""
	fmt.Println(GetIntervalTime(parseTime("2024-08-05 22:00:00"), parseTime("2024-08-04 22:00:11"))) // -23h59m49s ""
}

func TestBetweenDays(t *testing.T) {
	type args struct {
		startTime time.Time
		endTime   time.Time
	}
	tests := []struct {
		name string
		args args
		want int
	}{{
		name: "test1",
		args: args{
			startTime: parseDateTime("2024-08-06 12:10:56"),
			endTime:   parseDateTime("2024-08-05 03:10:56"),
		},
		want: 0,
	}, {
		name: "test2",
		args: args{
			startTime: parseDateTime("2024-08-05 12:10:56"),
			endTime:   parseDateTime("2024-08-05 03:10:56"),
		},
		want: 0,
	}, {
		name: "test3",
		args: args{
			startTime: parseDateTime("2024-08-05 12:10:56"),
			endTime:   parseDateTime("2024-08-05 23:10:56"),
		},
		want: 1,
	}, {
		name: "test4",
		args: args{
			startTime: parseDateTime("2024-08-05 12:10:56"),
			endTime:   parseDateTime("2024-08-06 11:10:56"),
		},
		want: 2,
	}, {
		name: "test5",
		args: args{
			startTime: parseDateTime("2024-08-05 12:10:56"),
			endTime:   parseDateTime("2024-08-09 11:10:56"),
		},
		want: 5,
	}, {
		name: "test6",
		args: args{
			startTime: parseDateTime("2024-08-05 12:10:56"),
			endTime:   parseDateTime("2024-09-05 11:10:56"),
		},
		want: 32,
	}, {
		name: "test7",
		args: args{
			startTime: parseDateTime("2024-08-05 12:10:56"),
			endTime:   parseDateTime("2025-08-05 11:10:56"),
		},
		want: 366,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BetweenDays(tt.args.startTime, tt.args.endTime); got != tt.want {
				t.Errorf("name %v, BetweenDays() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestSubDays(t *testing.T) {
	type args struct {
		startTime time.Time
		endTime   time.Time
	}
	tests := []struct {
		name string
		args args
		want int
	}{{
		name: "test1",
		args: args{
			startTime: parseDateTime("2024-08-06 12:10:56"),
			endTime:   parseDateTime("2024-08-05 03:10:56"),
		},
		want: 0,
	}, {
		name: "test2",
		args: args{
			startTime: parseDateTime("2024-08-05 12:10:56"),
			endTime:   parseDateTime("2024-08-05 03:10:56"),
		},
		want: 0,
	}, {
		name: "test3",
		args: args{
			startTime: parseDateTime("2024-08-05 12:10:56"),
			endTime:   parseDateTime("2024-08-05 23:10:56"),
		},
		want: 0,
	}, {
		name: "test4",
		args: args{
			startTime: parseDateTime("2024-08-05 12:10:56"),
			endTime:   parseDateTime("2024-08-06 12:10:56"),
		},
		want: 1,
	}, {
		name: "test4",
		args: args{
			startTime: parseDateTime("2024-08-05 12:10:56"),
			endTime:   parseDateTime("2024-08-06 13:10:56"),
		},
		want: 1,
	}, {
		name: "test5",
		args: args{
			startTime: parseDateTime("2024-08-05 12:10:56"),
			endTime:   parseDateTime("2024-08-09 11:10:56"),
		},
		want: 3,
	}, {
		name: "test6",
		args: args{
			startTime: parseDateTime("2024-08-05 12:10:56"),
			endTime:   parseDateTime("2024-09-05 11:10:56"),
		},
		want: 30,
	}, {
		name: "test7",
		args: args{
			startTime: parseDateTime("2024-08-05 12:10:56"),
			endTime:   parseDateTime("2025-08-05 11:10:56"),
		},
		want: 364,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubDays(tt.args.startTime, tt.args.endTime); got != tt.want {
				t.Errorf("name %v, BetweenDays() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestGetDaysInCurrentMonth(t *testing.T) {
	now := time.Now()
	for i := 0; i < 10; i++ {
		ts := AddMonth(now, i)
		fmt.Println(ts, GetDaysInCurrentMonth(ts))
	}
}

func TestGetValidNextMonthDay(t *testing.T) {
	now := time.Now()
	for i := 0; i < 10; i++ {
		ts := AddMonth(now, i)
		fmt.Println(ts, GetValidDaysInCurrentMonth(ts, 0))
	}
}
