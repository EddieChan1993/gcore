package clock

import (
	"log"
	"testing"
)

func TestTimeStampToDate(t *testing.T) {
	log.Println(TimeStampToDate(1666281600, TimeFormat3))
}

func TestDateToTime(t *testing.T) {
	time, _ := DateToTime("20221021", TimeFormat3)
	log.Println(time.Unix())
}
