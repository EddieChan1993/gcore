package utils

import (
	"fmt"
	"time"
)

type tfType string

const Milli int64 = 1000    //秒和毫秒的倍数关系,必须要给个确切类型值，因为通常和int64相乘，所以设定为int64
const DaySecs int64 = 86400 //一天的秒数

var TimeFormat tfType = "2006-01-02,15:04:05"

// TimeNowUnix 当前秒
func TimeNowUnix() int64 {
	return utcNow().Unix()
}

// TimeNowUnixMilli 当前毫秒
func TimeNowUnixMilli() int64 {
	return utcNow().UnixMilli()
}

// TimeNowUnixNano 当前纳秒
func TimeNowUnixNano() int64 {
	return utcNow().UnixNano()
}

// Week 第几周
func Week() int32 {
	year, week := utcNow().ISOWeek()
	return int32(week*10000 + year)
}

// Month 第几月
func Month() string {
	month := utcNow().Month().String()
	return fmt.Sprintf("%s%d", month, utcNow().Year())
}

// ZeroToday utc当天0点（秒）
func ZeroToday() int64 {
	return ZeroOtherDay(0)
}

// ZeroTimeAt 当前时间0点(秒）
func ZeroTimeAt(timeAt int64) int64 {
	t := time.Unix(timeAt, 0).UTC()
	startDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
	return startDay
}

// ZeroOtherDay 其他天（秒）
func ZeroOtherDay(day int) int64 {
	t := utcNow()
	return time.Date(t.Year(), t.Month(), t.Day()+day, 0, 0, 0, 0, t.Location()).Unix()
}

// ZeroToNextDay 指定时间的0点
func ZeroToNextDay(tt time.Time) int64 {
	t := utcNow()
	return time.Date(tt.Year(), tt.Month(), tt.Day()+1, 0, 0, 0, 0, t.Location()).Unix()
}

// DayNow 当前位于开始时间第几天
// start 开始时间毫秒
func DayNow(start int64) int32 {
	now := TimeNowUnixMilli()
	if start <= 0 || start > now {
		return 0
	}
	t := time.UnixMilli(start).UTC()
	startDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).UnixMilli()
	return int32((now-startDay)/(DaySecs*Milli) + 1)
}

// DayBetween 间隔多少天
func DayBetween(start, end int64) int32 {
	now := end
	if start <= 0 || start > now {
		return 0
	}
	t := time.UnixMilli(start).UTC()
	startDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).UnixMilli()
	return int32((now-startDay)/(DaySecs*Milli) + 1)
}

func utcNow() time.Time {
	return time.Now().UTC()
}

// WeekEndAt
// 周末结束时间点（毫秒）
// week为0本周,-1上周，1下周以此类推
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

// WeekDay 星期几
func WeekDay() int32 {
	weekDay := utcNow().Weekday()
	if time.Sunday == weekDay {
		return int32(7)
	}
	return int32(weekDay)
}

// TsWeekDay 时间蹉星期几
func TsWeekDay(ts int64) int32 {
	t := time.Unix(ts, 0)
	weekDay := t.Weekday()
	if time.Sunday == weekDay {
		return int32(7)
	}
	return int32(weekDay)
}

// MonthEndAt 月结束时间
// month 0本月 >0 接下来N月
func MonthEndAt(month int) (endTime int64) {
	now := utcNow()
	thisDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endTime = thisDay.AddDate(0, month+1, -now.Day()+1).UnixMilli()
	return endTime
}

// NextWeekCd 下一周到达cd
func NextWeekCd() int32 {
	now := utcNow()
	endAt := WeekEndAt(0)
	cd := endAt - now.UnixMilli()
	if cd < 0 {
		return 0
	}
	return int32(cd)
}

// NextMonthCd 距离下个月cd
func NextMonthCd() int32 {
	now := utcNow()
	endAt := MonthEndAt(0)
	cd := endAt - now.UnixMilli()
	if cd < 0 {
		return 0
	}
	return int32(cd)
}

// NextDayCd 距离跨天剩余时间
func NextDayCd() int32 {
	zeroOtherDay := ZeroOtherDay(1) * Milli
	cd := zeroOtherDay - TimeNowUnixMilli()
	return int32(cd)
}

// StarToEndCd 间隔cd
func StarToEndCd(star int64, keep int32) int32 {
	now := utcNow().Unix()
	cd := now - star
	if cd < 0 {
		cd = -cd
	}
	if cd > int64(keep) {
		return 0
	} else {
		return int32(int64(keep) - cd)
	}
}

// DayAfter N天后的0点时间戳
func DayAfter(day int, now int64) int64 {
	t := time.Unix(now, 0)
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	tm2 := tm1.AddDate(0, 0, day)
	return tm2.Unix()
}

// TSToDate 时间戳转为日期格式
func TSToDate(ts int64, tf tfType) string {
	return utcNow().Format(string(tf))
}

// DateToTime 日期格式转为时间类型
func DateToTime(ymd string, tf tfType) (time.Time, error) {
	return time.ParseInLocation(string(tf), ymd, utcNow().Location())
}
