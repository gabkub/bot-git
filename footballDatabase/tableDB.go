package footballDatabase

import (
	"bot-git/normalizedDate"
	"bot-git/persistence"
	bolt "go.etcd.io/bbolt"
	"log"
	"sort"
	"strings"
	"time"
)

var bucketName = "RESERVATION"

type FootballDb struct {
	per *persistence.Persistence
}

func NewFootballDb(dbPath string) *FootballDb {
	return &FootballDb{per: persistence.NewPersistence(dbPath, bucketName)}
}

func (f *FootballDb) IsBooked(newResUserName string, newResStartTime normalizedDate.NormalizeDate) bool {
	var result bool

	f.per.ReadDb(func(b *bolt.Bucket) error {
		err := b.ForEach(func(start, user []byte) error {
			resStart := normalizedDate.ConvertToTime(start)
			if userAlreadyBookedToday(resStart, newResStartTime, string(user), newResUserName) {
				result = true
			}
			return nil
		})
		if err != nil {
			log.Println(err)
		}
		return nil
	})
	return result
}

func (f *FootballDb) SetReservation(newResUserName string, newResStartTime normalizedDate.NormalizeDate) {
	tempTime := newResStartTime.AsString()
	f.per.WriteDb(func(b *bolt.Bucket) error {
		return addNewReservation(tempTime, newResUserName, b)
	})
}

func userAlreadyBookedToday(start normalizedDate.NormalizeDate, newResStart normalizedDate.NormalizeDate, userName, newResUserName string) bool {
	year, month, day := start.Date()
	return year == newResStart.Year() && month == newResStart.Month() &&
		day == newResStart.Day() && userName == newResUserName
}

func addNewReservation(tempTime []byte, userName string, b *bolt.Bucket) error {
	return b.Put(tempTime, []byte(strings.TrimSpace(userName)))
}
func (f *FootballDb) GetAllReservationByStartTime(startTime normalizedDate.NormalizeDate) TimeReservations {
	var reservations TimeReservations
	f.per.ReadDb(func(b *bolt.Bucket) error {
		searchingError := b.ForEach(func(start, username []byte) error {
			dbResStart := normalizedDate.ConvertToTime(start)
			if isReservationAfterSearchingDate(dbResStart, startTime) {
				reservation := &TimeReservation{
					UserName:  string(username),
					StartTime: dbResStart.Raw(),
				}
				reservations = append(reservations, reservation)
				return nil
			}
			return nil
		})
		if searchingError != nil {
			return searchingError
		}
		return nil
	})
	sort.Sort(reservations)
	return reservations
}
func isReservationAfterSearchingDate(dbResStart, startTime normalizedDate.NormalizeDate) bool {
	endTime := dbResStart.Add(gameDuration)
	return startTime.Equal(dbResStart.Raw()) || dbResStart.After(startTime.Raw()) ||
		endTime.Equal(startTime.Raw()) || endTime.After(startTime.Raw())
}

func (f *FootballDb) IsFree(paramTime normalizedDate.NormalizeDate) bool {
	possibleEndTime := normalizedDate.NewNormalizeDate(paramTime.Add(gameDuration))
	reservations := f.GetAllReservationByStartTime(paramTime)
	for _, reservation := range reservations {
		if reservation.IsBetween(paramTime) || reservation.IsBetween(possibleEndTime) {
			return false
		}
	}
	return true
}

func (f *FootballDb) FirstFreeTimeFor(paramTime normalizedDate.NormalizeDate) time.Time {
	reservations := f.GetAllReservationByStartTime(paramTime)
	if len(reservations) == 0 {
		return paramTime.Raw()
	}
	if len(reservations) == 1 {
		return reservations[0].EndTime()
	}
	for i := 1; i < len(reservations); i++ {
		previousEnd := reservations[i-1].EndTime()
		currentStart := reservations[i].StartTime
		span := currentStart.Sub(previousEnd)
		if span == gameDuration || span > gameDuration {
			return previousEnd
		}
	}
	return reservations[len(reservations)-1].EndTime()
}
