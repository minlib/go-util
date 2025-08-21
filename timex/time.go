// Copyright 2024 Minzhan.com Inc. All rights reserved.

package timex

import (
	"github.com/minlib/go-util/core"
	"strconv"
	"time"
)

// FormatDate 获取日期字符串
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// FormatTime 获取时间字符串
func FormatTime(t time.Time) string {
	return t.Format("15:04:05")
}

// FormatDateTime 获取日期时间字符串
func FormatDateTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// ParseDateTime 时间字符串转换为日期时间
func ParseDateTime(value string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", value)
}

// ParseInLocation is like Parse but differs in two important ways.
func ParseInLocation(value string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", value, loc)
}

// IsSameDay 判断两个日期是同一天
func IsSameDay(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// AddDuration 增加时间
func AddDuration(t time.Time, interval int, unit string) time.Time {
	duration, err := time.ParseDuration(strconv.Itoa(interval) + unit)
	if err != nil {
		panic(err)
	}
	return t.Add(duration)
}

// AddSecond add seconds
func AddSecond(t time.Time, seconds int) time.Time {
	if seconds == 0 {
		return t
	}
	return t.Add(time.Duration(seconds) * time.Second)
}

// AddMinute add minutes
func AddMinute(t time.Time, minutes int) time.Time {
	if minutes == 0 {
		return t
	}
	return t.Add(time.Duration(minutes) * time.Minute)
}

// AddHour add hours
func AddHour(t time.Time, hours int) time.Time {
	if hours == 0 {
		return t
	}
	return t.Add(time.Duration(hours) * time.Hour)
}

// AddDay add days
func AddDay(t time.Time, days int) time.Time {
	if days == 0 {
		return t
	}
	return t.AddDate(0, 0, days)
}

// AddMonth add months
func AddMonth(t time.Time, months int) time.Time {
	if months == 0 {
		return t
	}
	return t.AddDate(0, months, 0)
}

// AddYear add years
func AddYear(t time.Time, years int) time.Time {
	if years == 0 {
		return t
	}
	return t.AddDate(years, 0, 0)
}

func UnixFormat(t time.Time) string {
	return time.Unix(t.Unix(), 0).Format("2006-01-02 15:04:05")
}

func UnixMilliFormat(t time.Time) string {
	return time.UnixMilli(t.UnixMilli()).Format("2006-01-02 15:04:05.000")
}

func UnixMicroFormat(t time.Time) string {
	return time.UnixMicro(t.UnixMicro()).Format("2006-01-02 15:04:05.000000")
}

// ToBeijingTime 转成北京时间
func ToBeijingTime(t time.Time) time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return t.In(loc)
}

// ToBeijingZone 转成北京时区，时间字符串保持不变
func ToBeijingZone(t time.Time) time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), loc)
}

// ToDateTime converts a *time.Time pointer to a *core.DateTime pointer.
// If the input time pointer is nil, it returns nil.
// Otherwise, it creates and returns a new *core.DateTime based on the provided time.
func ToDateTime(t *time.Time) *core.DateTime {
	if t == nil {
		return nil
	}
	return core.NewDateTime(*t)
}
