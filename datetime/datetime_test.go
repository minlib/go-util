package datetime

import (
	"fmt"
	"testing"
	"time"
)

func TestTodayAddDay(t *testing.T) {
	now := time.Now()
	fmt.Println(now)             // 2022-10-05 23:35:22.3942125 +0800 CST m=+0.003905101
	fmt.Println(DateString(now)) // 20221005
	fmt.Println(TimeString(now)) // 23:35:22
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
}
