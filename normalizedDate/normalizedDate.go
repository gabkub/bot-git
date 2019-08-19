package normalizedDate

import (
	"strconv"
	"strings"
	"time"
)

type NormalizeDate struct {
	time.Time
}

func (d NormalizeDate) Raw() time.Time {
	return d.Time
}

func NewNormalizeDate(t time.Time) NormalizeDate {
	loc, _ := time.LoadLocation("Local")
	t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, loc).UTC()
	return NormalizeDate{t.In(loc)}
}

func (d NormalizeDate) AsString() []byte {
	ticks := d.UTC().Unix()
	str := strconv.FormatInt(ticks, 10)
	return []byte(strings.TrimSpace(str))

}
func ConvertToTime(paramTime []byte) NormalizeDate {
	ticks, _ := strconv.ParseInt(string(paramTime), 10, strconv.IntSize)
	t := time.Unix(ticks, 0)
	return NewNormalizeDate(t)
}
