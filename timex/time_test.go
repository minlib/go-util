package timex

import (
	"fmt"
	"testing"
	"time"
)

func TestTodayAddDay(t *testing.T) {
	now := time.Now()
	fmt.Println(now)                                  // 2024-08-05 22:10:56.2059197 +0800 CST m=+0.001098601
	fmt.Println(FormatDate(now))                      // 2024-08-05
	fmt.Println(FormatTime(now))                      // 22:10:56
	fmt.Println(FormatDateTime(now))                  // 2024-08-05 22:10:56
	fmt.Println(ParseDateTime("2024-08-05 22:10:56")) // 2024-08-05 22:10:56 +0000 UTC <nil>
	fmt.Println(IsSameDay(now, time.Now()))           // true
	// 增加1小时
	fmt.Println(AddDuration(now, 1, "h")) // 2022-10-06 00:35:22.3942125 +0800 CST m=+3600.003905101
	fmt.Println(AddSecond(now, 1))        // 2022-10-05 23:35:23.3942125 +0800 CST m=+1.003905101
	fmt.Println(AddMinute(now, 1))        // 2022-10-05 23:36:22.3942125 +0800 CST m=+60.003905101
	fmt.Println(AddHour(now, 1))          // 2022-10-06 00:35:22.3942125 +0800 CST m=+3600.003905101
	fmt.Println(AddDay(now, 1))           // 2022-10-06 23:35:22.3942125 +0800 CST
	fmt.Println(AddMonth(now, 1))         // 2022-11-05 23:35:22.3942125 +0800 CST
	fmt.Println(AddYear(now, 1))          // 2023-10-05 23:35:22.3942125 +0800 CST
	fmt.Println(StartDateTime(now))       // 2022-10-05 00:00:00 +0800 CST
	fmt.Println(EndDateTime(now))         // 2022-10-05 23:59:59.999999999 +0800 CST
	fmt.Println(RangeDateTime(now))       // 2022-10-05 00:00:00 +0800 CST  2022-10-05 23:59:59.999999999 +0800 CST

	fmt.Println(UnixFormat(now))      // 2023-07-05 02:21:25
	fmt.Println(UnixMilliFormat(now)) // 2023-07-05 02:21:25.622
	fmt.Println(UnixMicroFormat(now)) // 2023-07-05 02:21:25.622508
}

func parseTime(timeString string) time.Time {
	t, _ := ParseDateTime(timeString)
	return t
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

func parseDateTime(value string) time.Time {
	start, _ := ParseInLocation(value, time.Local)
	return start
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
