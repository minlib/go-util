package timex

import (
	"strconv"
	"time"
)

// DateString 获取日期字符串
func DateString(t time.Time) string {
	return t.Format("20060102")
}

// TimeString 获取时间字符串
func TimeString(t time.Time) string {
	return t.Format("15:04:05")
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
	return AddDuration(t, seconds, "s")
}

// AddMinute add minutes
func AddMinute(t time.Time, minutes int) time.Time {
	return AddDuration(t, minutes, "m")
}

// AddHour add hours
func AddHour(t time.Time, hours int) time.Time {
	return AddDuration(t, hours, "h")
}

// AddDay add days
func AddDay(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

// AddMonth add months
func AddMonth(t time.Time, months int) time.Time {
	return t.AddDate(0, months, 0)
}

// AddYear add years
func AddYear(t time.Time, years int) time.Time {
	return t.AddDate(years, 0, 0)
}

// StartDateTime 当天的开始时间
func StartDateTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

// EndDateTime 当天的结束时间
func EndDateTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, time.Local)
}

// RangeDateTime 当天的开始与结束时间
func RangeDateTime(t time.Time) (time.Time, time.Time) {
	return StartDateTime(t), EndDateTime(t)
}
