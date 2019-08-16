package footballDatabase

import (
	"bot-git/persistence"
	bolt "go.etcd.io/bbolt"
	"log"
	"sort"
	"strconv"
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

func (f *FootballDb) IsBooked(newResUserName string, newResStartTime NormalizeDate) bool {
	var result bool

	f.per.ReadDb(func(b *bolt.Bucket) error {
		err := b.ForEach(func(start, user []byte) error {
			resStart := convertToTime(start)
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

func (f *FootballDb) SetReservation(newResUserName string, newResStartTime NormalizeDate) {
	tempTime := timeToKey(newResStartTime)
	f.per.WriteDb(func(b *bolt.Bucket) error {
		return addNewReservation(tempTime, newResUserName, b)
	})
}

func userAlreadyBookedToday(start NormalizeDate, newResStart NormalizeDate, userName, newResUserName string) bool {
	year, month, day := start.Date()
	return year == newResStart.Year() && month == newResStart.Month() &&
		day == newResStart.Day() && userName == newResUserName
}

func addNewReservation(tempTime, userName string, b *bolt.Bucket) error {
	return b.Put([]byte(strings.TrimSpace(tempTime)), []byte(strings.TrimSpace(userName)))
}
func (f *FootballDb) GetAllReservationByStartTime(startTime NormalizeDate) TimeReservations {
	var reservations TimeReservations
	f.per.ReadDb(func(b *bolt.Bucket) error {
		searchingError := b.ForEach(func(start, username []byte) error {
			dbResStart := convertToTime(start)
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
func isReservationAfterSearchingDate(dbResStart, startTime NormalizeDate) bool {
	endTime := dbResStart.Add(gameDuration)
	return startTime.Equal(dbResStart.Raw()) || dbResStart.After(startTime.Raw()) ||
		endTime.Equal(startTime.Raw()) || endTime.After(startTime.Raw())
}

func (f *FootballDb) IsFree(paramTime NormalizeDate) bool {
	const isFree = true
	possibleEndTime := NewNormalizeDate(paramTime.Add(gameDuration))
	reservations := f.GetAllReservationByStartTime(paramTime)
	for _, reservation := range reservations {
		if reservation.IsBetween(paramTime) || reservation.IsBetween(possibleEndTime) {
			return !isFree
		}
	}
	return isFree
}

func (f *FootballDb) FirstFreeTimeFor(paramTime NormalizeDate) time.Time {
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

func timeToKey(param NormalizeDate) string {
	ticks := param.UTC().Unix()
	return strconv.FormatInt(ticks, 10)
}
func convertToTime(paramTime []byte) NormalizeDate {
	ticks, _ := strconv.ParseInt(string(paramTime), 10, strconv.IntSize)
	t := time.Unix(ticks, 0)
	return NewNormalizeDate(t)
}
