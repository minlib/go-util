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

	fmt.Println(UnixFormat(now))      // 2023-07-05 02:21:25
	fmt.Println(UnixMilliFormat(now)) // 2023-07-05 02:21:25.622
	fmt.Println(UnixMicroFormat(now)) // 2023-07-05 02:21:25.622508
}

func parseTime(timeString string) time.Time {
	t, _ := ParseDateTime(timeString)
	return t
}

func parseDateTime(value string) time.Time {
	start, _ := ParseInLocation(value, time.Local)
	return start
}
