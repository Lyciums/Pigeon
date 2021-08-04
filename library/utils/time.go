package utils

import (
	"strconv"
	"time"
)

var (
	// 默认日期格式
	defaultDateFormat = "2006-01-02 15:04:05"
	// 默认时区
	defaultDateLocation = "Asia/Shanghai"
	// SecondsPerMinute 定义每分钟的秒数
	SecondsPerMinute = 60
	// SecondsPerHour 定义每小时的秒数
	SecondsPerHour = SecondsPerMinute * 60
	// SecondsPerDay 定义每天的秒数
	SecondsPerDay = SecondsPerHour * 24
	// SecondsPerWeek 每周的秒数
	SecondsPerWeek = SecondsPerDay * 7
)

// SetDefaultDateFormat 设置默认日期格式
func SetDefaultDateFormat(format string) {
	defaultDateFormat = format
}

// SetDefaultDateLocation 设置默认时区
func SetDefaultDateLocation(location string) {
	defaultDateLocation = location
}


// GetCurrentTimestamp 获取当前时间戳
func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}

// FormatCurrentTime 获取当前时间戳
func FormatCurrentTime(format string) string {
	return time.Now().Format(format)
}

// GetCurrentTimestampString 获取当前时间戳字符串
func GetCurrentTimestampString() string {
	return strconv.FormatInt(GetCurrentTimestamp(), 10)
}

// ParseTimestamp 解析时间戳成 time.Time 结构体
func ParseTimestamp(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

// GetFormatDateByTimestamp 时间戳转日期
func GetFormatDateByTimestamp(timestamp int64, format string) string {
	if format == "" {
		format = defaultDateFormat
	}
	return ParseTimestamp(timestamp).Format(format)
}

// TimestampToDate 通过时间戳获取日期
func TimestampToDate(timestamp int64) string {
	return GetFormatDateByTimestamp(timestamp, "")
}

// GetTimestampFromDateByLocation 获取指定日期和时区的时间戳
func GetTimestampFromDateByLocation(date, format, location string) int64 {
	if format == "" {
		format = defaultDateFormat
	}
	if location == "" {
		location = defaultDateLocation
	}
	loc, _ := time.LoadLocation(location)
	t, _ := time.ParseInLocation(format, date, loc)
	return t.Unix()
}

// GetTimestampFromDate 获取指定日期的时间戳
func GetTimestampFromDate(date, format string) int64 {
	return GetTimestampFromDateByLocation(date, format, "")
}

// DateToTimestamp 时间转时间戳
func DateToTimestamp(date string) int64 {
	return GetTimestampFromDate(date, "")
}
