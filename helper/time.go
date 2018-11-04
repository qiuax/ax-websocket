package helper

import (
	"errors"
	"fmt"
	syslog "log"
	"strconv"
	"strings"
	"time"
)

// 获取中国地区
var ChinaLocation *time.Location

func init() {
	var err error

	ChinaLocation, err = time.LoadLocation("Asia/Shanghai") //上海

	if err != nil {
		syslog.Fatalf("解析中国时区地区出错,err: %v", err)
	}
}

// 获取(中国)当前时间
func ChinaTimeNow() time.Time {
	return time.Now().In(ChinaLocation)
}

// 获取unix时间戳
func GetTimestamp() int64 {
	return ChinaTimeNow().Unix()
}

// 获取unix时间戳(纳秒)
func GetTimeNano() int64 {
	return ChinaTimeNow().UnixNano()
}

// 获取当天(中国)的开始时间(时间戳)
func TodayStartTime() int64 {
	now := ChinaTimeNow()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Unix()
}

// 当前小时时间
func CurrentHourStartTime() int64 {
	now := ChinaTimeNow()
	return time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location()).Unix()
}

// 格式化日期 Y-m-d H:i:s (没传第二个参数，默认为当前时间戳)
func FormatDate(format string, timestamp int64) string {
	var t time.Time

	if timestamp > 0 {
		t = time.Unix(timestamp, 0)
	} else {
		t = time.Now()
	}

	t.In(ChinaLocation) // 指定时区

	// 替换年月日
	format = strings.Replace(format, "Y", fmt.Sprintf("%d", t.Year()), -1)
	format = strings.Replace(format, "m", fmt.Sprintf("%02d", int(t.Month())), -1)
	format = strings.Replace(format, "d", fmt.Sprintf("%02d", t.Day()), -1)

	// 替换时分秒
	format = strings.Replace(format, "H", fmt.Sprintf("%02d", t.Hour()), -1)
	format = strings.Replace(format, "i", fmt.Sprintf("%02d", t.Minute()), -1)
	format = strings.Replace(format, "s", fmt.Sprintf("%02d", t.Second()), -1)

	return format
}

// 获取3天的时间长度(秒)
func GetThreeDayTimes() int64 {
	return int64(3 * 86400)
}

// 判断时间戳是否是一天的0点时间
func IsDayStartTimestamp(timestamp int64) bool {
	t := time.Unix(timestamp, int64(0))
	if t.Hour() == 0 && t.Minute() == 0 && t.Second() == 0 {
		return true
	}
	return false
}

// 检测YYYYmmdd日期是否正确, 成功返回对应的时间戳
func CheckDateIsCorrect(dateStr string) (timestamp int64, err error) {
	var year, month, day int

	if year, err = strconv.Atoi(dateStr[0:4]); err != nil {
		return
	}

	if month, err = strconv.Atoi(dateStr[4:6]); err != nil {
		return
	}

	if day, err = strconv.Atoi(dateStr[6:8]); err != nil {
		return
	}

	ChinaLocation, _ := time.LoadLocation("Asia/Shanghai") //上海
	d := time.Date(year, time.Month(month), day, 0, 0, 0, 0, ChinaLocation)

	if year != d.Year() || month != int(d.Month()) || day != d.Day() || d.Unix() < 0 {
		err = errors.New("日期不存在")
		return
	}

	timestamp = d.Unix()

	return
}
