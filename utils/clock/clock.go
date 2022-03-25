package clock

import "time"

func TimeStampToDate(ts int64) string {
	return time.Unix(ts, 0).Format("20060102150405")
}
