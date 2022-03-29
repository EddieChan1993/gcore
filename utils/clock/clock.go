package clock

import "time"

type tfType string

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

func loc() (*time.Location, error) {
	return time.LoadLocation("Asia/Shanghai")
}
