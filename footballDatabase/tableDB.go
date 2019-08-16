package footballDatabase

import (
	"bot-git/logg"
	"fmt"
	bolt "go.etcd.io/bbolt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

var bucketName = []byte("RESERVATION")

type FootballDb struct {
	dbPath string
}

func NewFootballDb(dbPath string) *FootballDb {
	db := &FootballDb{dbPath: dbPath}
	db.createTableDB()
	return db
}

func (f *FootballDb) readDb(function func(b *bolt.Bucket) error) {
	db := f.openDb(true)
	if db == nil {
		return
	}
	defer db.Close()
	viewError := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		return function(b)
	})
	if viewError != nil {
		log.Println(viewError)
	}
}
func (f *FootballDb) writeDb(function func(b *bolt.Bucket) error) {
	db := f.openDb(false)
	if db == nil {
		return
	}
	defer db.Close()
	updateError := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		return function(b)
	})
	if updateError != nil {
		log.Fatal(fmt.Sprintf("Unable to update. Error: %s", updateError))
	}
}

func (f *FootballDb) openDb(readOnly bool) *bolt.DB {
	db, err := bolt.Open(f.dbPath, 0600, &bolt.Options{ReadOnly: readOnly})
	if err != nil {
		if err.Error() == "timeout" {
			log.Fatal("Database already opened.")
		} else {
			log.Fatal("Error opening the database. " + err.Error())
		}
	}
	return db
}

func (f *FootballDb) createTableDB() {
	db, err := bolt.Open(f.dbPath, 0600, nil)
	if err != nil {
		logg.WriteToFile("Error opening or creating database.")
		log.Fatal("Error opening or creating database.")
	}
	defer db.Close()

	createError := db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return fmt.Errorf("create bucket error: %s", err)
		}
		return nil
	})
	if createError != nil {
		logg.WriteToFile("Error creating the table. " + createError.Error())
		log.Println("Error creating the table. " + createError.Error())
	}
	logg.WriteToFile("Created database table for booking.")
}

func (f *FootballDb) IsBooked(newResUserName string, newResStartTime NormalizeDate) bool {
	var result bool

	f.readDb(func(b *bolt.Bucket) error {
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
	f.writeDb(func(b *bolt.Bucket) error {
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
	f.readDb(func(b *bolt.Bucket) error {
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
