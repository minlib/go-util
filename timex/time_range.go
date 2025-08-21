package timex

import (
	"strconv"
	"time"
)

// TimeRange 时间范围，包含起始时间和结束时间
type TimeRange struct {
	StartTime time.Time `json:"startTime" form:"startTime"` // 起始时间
	EndTime   time.Time `json:"endTime" form:"endTime"`     // 结束时间
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

// GetIntervalTime calculates the duration between startTime and endTime. (startTime-endTime)
// It returns the duration as a time.Duration object and a string representation of the duration.
func GetIntervalTime(startTime time.Time, endTime time.Time) (time.Duration, string) {
	duration := endTime.Sub(startTime)
	if duration <= 0 {
		return duration, ""
	}
	days := int(duration.Hours() / 24)
	years := days / 365
	if years > 0 {
		return duration, strconv.Itoa(years) + "年"
	}
	months := days / 30
	if months > 0 {
		return duration, strconv.Itoa(months) + "个月"
	}
	if days > 0 {
		return duration, strconv.Itoa(days) + "天"
	}
	hours := int(duration.Hours())
	if hours > 0 {
		return duration, strconv.Itoa(hours) + "小时"
	}
	minutes := int(duration.Minutes())
	if minutes > 0 {
		return duration, strconv.Itoa(minutes) + "分钟"
	}
	seconds := int(duration.Seconds())
	if seconds > 0 {
		return duration, strconv.Itoa(seconds) + "秒"
	}
	return duration, ""
}

// GetDaysInCurrentMonth 获取当前时间月份的总天数
// 逻辑：下个月的第一天减去1天，即为当月最后一天
func GetDaysInCurrentMonth(t time.Time) int {
	nextMonthFirstDay := time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, t.Location())
	lastDayOfCurrentMonth := nextMonthFirstDay.AddDate(0, 0, -1)
	return lastDayOfCurrentMonth.Day()
}

// GetValidDaysInCurrentMonth 获取当前时间月份的有效天数
func GetValidDaysInCurrentMonth(t time.Time, day int) int {
	lastDaysInCurrentMonth := GetDaysInCurrentMonth(t)
	if day <= 0 || day > lastDaysInCurrentMonth {
		return lastDaysInCurrentMonth
	}
	return day
}
