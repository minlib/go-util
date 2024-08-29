// Copyright 2024 Minzhan.com Inc. All rights reserved.

package timex

import (
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

// BetweenDays 获取2个日期的涉及天数
func BetweenDays(startTime, endTime time.Time) int {
	if startTime.After(endTime) {
		return 0
	}
	start := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, time.UTC)
	end := time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 0, 0, 0, 0, time.UTC)
	duration := end.Sub(start)
	return int(duration.Hours()/24) + 1
}

// SubDays 获取2个日期的相差天数
func SubDays(startTime, endTime time.Time) int {
	duration := endTime.Sub(startTime)
	if duration < 0 {
		return 0
	}
	return int(duration.Hours() / 24)
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

// GetIntervalString 获取两个时间差字符串
func GetIntervalString(startTime time.Time, endTime time.Time) string {
	duration := endTime.Sub(startTime)
	if duration < 0 {
		return ""
	}
	days := int(duration.Hours() / 24)
	years := days / 365
	if years > 0 {
		return strconv.Itoa(years) + "年"
	}
	months := days / 30
	if months > 0 {
		return strconv.Itoa(months) + "个月"
	}
	if days > 0 {
		return strconv.Itoa(days) + "天"
	}
	hours := int(duration.Hours())
	if hours > 0 {
		return strconv.Itoa(hours) + "小时"
	}
	minutes := int(duration.Minutes())
	if minutes > 0 {
		return strconv.Itoa(minutes) + "分钟"
	}
	seconds := int(duration.Seconds())
	if seconds > 0 {
		return strconv.Itoa(seconds) + "秒"
	}
	return ""
}
