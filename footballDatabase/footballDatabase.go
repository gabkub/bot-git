package footballDatabase

import "time"

const (
	gameDuration = 20 * time.Minute
)

type NormalizeDate struct {
	time.Time
}

func (d NormalizeDate) Raw() time.Time {
	return d.Time
}

func NewNormalizeDate(t time.Time) NormalizeDate {
	t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
	return NormalizeDate{t}
}

type TimeReservation struct {
	UserName  string
	StartTime time.Time
}

func NewTimeReservation(userName string, startTime time.Time) *TimeReservation {
	return &TimeReservation{UserName: userName, StartTime: startTime}
}

type TimeReservations []*TimeReservation

func (t TimeReservations) Len() int           { return len(t) }
func (t TimeReservations) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t TimeReservations) Less(i, j int) bool { return t[i].StartTime.Before(t[j].StartTime) }

func (t *TimeReservation) EndTime() time.Time {
	return t.StartTime.Add(gameDuration)
}
func (t *TimeReservation) IsBetween(d NormalizeDate) bool {
	return (t.StartTime.Equal(d.Raw()) || t.StartTime.Before(d.Raw())) &&
		(t.EndTime().After(d.Raw()))
}
