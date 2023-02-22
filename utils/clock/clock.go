package clock

import "time"

type tfType string

const Milli int64 = 1000    //秒和毫秒的倍数关系,必须要给个确切类型值，因为通常和int64相乘，所以设定为int64
const DaySecs int64 = 86400 //一天的秒数

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

//TimeNowUnix 当前秒
func TimeNowUnix() int64 {
	return utcNow().Unix()
}

//TimeNowUnixMilli 当前毫秒
func TimeNowUnixMilli() int64 {
	return utcNow().UnixMilli()
}

// TimeNowUnixNano 当前纳秒
func TimeNowUnixNano() int64 {
	return utcNow().UnixNano()
}

//Week 第几周
func Week() int32 {
	year, week := time.Now().UTC().ISOWeek()
	return int32(week*10000 + year)
}

//ZeroToday utc当天0点（秒）
func ZeroToday() int64 {
	return ZeroOtherDay(0)
}

//ZeroOtherDay 其他天（秒）
func ZeroOtherDay(day int) int64 {
	t := utcNow()
	return time.Date(t.Year(), t.Month(), t.Day()+day, 0, 0, 0, 0, t.Location()).Unix()
}

//DayNow 当前位于开始时间第几天
//start 开始时间毫秒
func DayNow(start int64) int32 {
	now := TimeNowUnixMilli()
	if start <= 0 || start > now {
		return 0
	}
	t := time.UnixMilli(start).UTC()
	startDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).UnixMilli()
	return int32((now-startDay)/(DaySecs*Milli) + 1)
}

//NextDayCd 距离跨天剩余时间
func NextDayCd() int32 {
	zeroOtherDay := ZeroOtherDay(1) * Milli
	cd := zeroOtherDay - TimeNowUnixMilli()
	return int32(cd)
}

func utcNow() time.Time {
	return time.Now().UTC()
}

//WeekEndAt
//周末结束时间点（毫秒）
//week为0本周,-1上周，1下周以此类推
func WeekEndAt(week int) (endTime int64) {
	now := utcNow()
	offset := int(now.Weekday())
	if offset == 0 {
		//周天默认=0，因此转为8-7=1
		offset = 1
	} else {
		offset = 8 - offset
	}
	thisWeek := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endTime = thisWeek.AddDate(0, 0, offset+7*week).UnixMilli()
	return endTime
}

//NextWeekCd 下一周到达cd
func NextWeekCd() int32 {
	now := utcNow()
	endAt := WeekEndAt(0)
	cd := endAt - now.UnixMilli()
	if cd < 0 {
		return 0
	}
	return int32(cd)
}
