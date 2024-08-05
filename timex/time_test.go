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
	fmt.Println(GetIntervalString(parseTime("2024-08-05 22:00:00"), parseTime("2026-08-15 22:00:00"))) // 2
	fmt.Println(GetIntervalString(parseTime("2024-08-05 22:00:00"), parseTime("2025-08-05 22:00:00"))) // 1
	fmt.Println(GetIntervalString(parseTime("2024-08-05 22:00:00"), parseTime("2025-08-04 22:00:00"))) // 12个月
	fmt.Println(GetIntervalString(parseTime("2024-08-05 22:00:00"), parseTime("2025-03-03 21:00:00"))) // 6个月
	fmt.Println(GetIntervalString(parseTime("2024-08-05 22:00:00"), parseTime("2024-08-25 21:00:00"))) // 19天
	fmt.Println(GetIntervalString(parseTime("2024-08-05 22:00:00"), parseTime("2024-08-06 21:00:00"))) // 23小时
	fmt.Println(GetIntervalString(parseTime("2024-08-05 22:00:00"), parseTime("2024-08-05 22:03:11"))) // 3分钟
	fmt.Println(GetIntervalString(parseTime("2024-08-05 22:00:00"), parseTime("2024-08-05 22:00:11"))) // 11秒
}
