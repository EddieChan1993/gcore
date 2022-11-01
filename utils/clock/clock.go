package clock

import "time"

type tfType string

const Milli = 1000    //秒和毫秒的倍数关系
const DaySecs = 86400 //一天的秒数

var TimeFormat tfType = "2006-01-02 15:04:05"
var TimeFormat2 tfType = "20060102150405"
var TimeFormat3 tfType = "20060102"

func TimeStampToDate(ts int64, tf tfType) string {
	return time.Unix(ts, 0).Format(string(tf))
}

func DateToTime(ymd string, tf tfType) (time.Time, error) {
	loc, err := loc()
	if err != nil {
		return time.Time{}, err
	}
	return time.ParseInLocation(string(tf), ymd, loc)
}

// TimeToMs returns an integer number, which represents t in milliseconds.
func TimeToMs(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// MsToTime returns the UTC time corresponding to the given Unix time,
// t milliseconds since January 1, 1970 UTC.
func MsToTime(t int64) time.Time {
	return time.Unix(0, t*int64(time.Millisecond)).UTC()
}

func loc() (*time.Location, error) {
	return time.LoadLocation("Asia/Shanghai")
}

//TimeNowUnixMilli 当前毫秒
func TimeNowUnixMilli() int64 {
	return time.Now().UnixMilli()
}

//Week 第几周
func Week() int32 {
	year, week := time.Now().ISOWeek()
	return int32(week*10000 + year)
}

//ZeroToday 当天0点（秒）
func ZeroToday() int64 {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
}

//ZeroOtherDay 其他天（秒）
func ZeroOtherDay(day int) int64 {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day()+day, 0, 0, 0, 0, t.Location()).Unix()
}

//DayNow 当前位于开始时间第几天
//start 开始时间毫秒
func DayNow(start int64) int32 {
	now := TimeNowUnixMilli()
	if start <= 0 || start > now {
		return 0
	}
	t := time.UnixMilli(start)
	startDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).UnixMilli()
	return int32((now-startDay)/(DaySecs*Milli) + 1)
}
