// Copyright 2024 Minzhan.com Inc. All rights reserved.

package timex

import (
	"strconv"
	"time"
)

// DateString 获取日期字符串
func DateString(t time.Time) string {
	return t.Format("2006-01-02")
}

// TimeString 获取时间字符串
func TimeString(t time.Time) string {
	return t.Format("15:04:05")
}

// DateTimeString 获取日期时间字符串
func DateTimeString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
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
	return t.Add(time.Duration(seconds) * time.Second)
}

// AddMinute add minutes
func AddMinute(t time.Time, minutes int) time.Time {
	return t.Add(time.Duration(minutes) * time.Minute)
}

// AddHour add hours
func AddHour(t time.Time, hours int) time.Time {
	return t.Add(time.Duration(hours) * time.Hour)
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
