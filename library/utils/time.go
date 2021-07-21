package utils

import (
	"strconv"
	"time"
)

var (
	// 默认日期格式
	defaultDateFormat = "0000-00-00 00:00:00"
	// 默认时区
	defaultDateLocation = "Asia/Shanghai"
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
