package footballDatabase

import (
	"bot-git/normalizedDate"
	"time"
)

const (
	gameDuration = 25 * time.Minute
)

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
func (t *TimeReservation) IsBetween(d normalizedDate.NormalizeDate) bool {
	return (t.StartTime.Equal(d.Raw()) || t.StartTime.Before(d.Raw())) &&
		(t.EndTime().After(d.Raw()))
}
