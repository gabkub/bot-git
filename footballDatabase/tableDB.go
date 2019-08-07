package footballDatabase

import (
	"bot-git/logg"
	"fmt"
	bolt "go.etcd.io/bbolt"
	"log"
	"strconv"
	"strings"
	"time"
)

type TimeReservation struct {
	UserName  string
	StartTime string
	EndTime   string
}

func CreateTableDB() {
	db, err := bolt.Open("./footballTable.db", 0600, nil)
	if err != nil {
		logg.WriteToFile("Error opening or creating database.")
		log.Fatal("Error opening or creating database.")
	}

	defer db.Close()

	createError := db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists([]byte(strings.TrimSpace("RESERVATION")))

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
func SetReservation(userName string, startTime time.Time) bool {

	tempTime := TimeToString(roundToMinute(startTime))
	//if TimeToString(roundToMinute(time.Now())) > tempTime{
	//	return
	//}

	db, err := bolt.Open("./footballTable.db", 0600, &bolt.Options{ReadOnly: false, Timeout: 1 * time.Second})

	if err != nil {
		if err.Error() == "timeout" {
			log.Fatal("Database already opened.")
		} else {
			log.Fatal("Error opening the database. " + err.Error())
		}
	}

	defer db.Close()

	canAddUser := true

	viewError := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(strings.TrimSpace("RESERVATION")))
		err := b.ForEach(func(start, user []byte) error {
			year, month, day := convertStringToTime(string(start)).Date()

			if year == startTime.Year() && month == startTime.Month() && day == startTime.Day() {
				if userName == string(user) {
					canAddUser = false
					//limit.AddRequest()
				}
			}
			return nil

		})

		if err != nil {
			log.Println(err)
		}
		return nil
	})

	if viewError != nil {
		log.Println(viewError)
	}

	if canAddUser {
		updateError := db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(strings.TrimSpace("RESERVATION")))

			err := b.Put([]byte(strings.TrimSpace(tempTime)), []byte(strings.TrimSpace(userName)))
			return err
		})

		if updateError != nil {
			log.Fatal(fmt.Sprintf("Unable to update. Error: %s", updateError))
		}
	}
	return canAddUser
}
func GetAllReservationByStartTime(startTime time.Time) []TimeReservation {
	db, _ := bolt.Open("footballTable.db", 0600, &bolt.Options{ReadOnly: true})

	defer db.Close()

	startTime = roundToMinute(startTime)
	startTimeAsString := TimeToString(startTime)
	var reservations []TimeReservation
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("RESERVATION"))

		searchingError := bucket.ForEach(func(start, username []byte) error {

			if strings.Compare(startTimeAsString, string(start)) == -1 ||
				strings.Compare(startTimeAsString, TimeToString(convertStringToTime(string(start)).Add(time.Minute*20))) == -1 {
				reservation := TimeReservation{
					UserName:  string(username),
					StartTime: string(start),
					EndTime:   TimeToString(convertStringToTime(string(start)).Add(time.Minute * 20)),
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

	if err != nil {
		logg.WriteToFile("Error reading the football reservations. " + err.Error())
		log.Println("Error reading the football reservations. " + err.Error())
	}

	return reservations
}
func FreeReservation(paramTime time.Time) time.Time {
	reservations := GetAllReservationByStartTime(paramTime)
	paramTimeAsString := TimeToString(paramTime)

	if len(reservations) == 0 {
		return paramTime
	}

	for _, reservation := range reservations {
		if reservation.StartTime <= paramTimeAsString && paramTimeAsString <= reservation.EndTime {
			continue
		}
		timeStamp := convertStringToTime(reservation.StartTime).Sub(convertStringToTime(TimeToString(paramTime))).Minutes()

		if timeStamp < 20 {
			return time.Time{}
		}
	}

	var lastEndTimeReservation string

	for _, reservation := range reservations {

		if lastEndTimeReservation != "" {
			timeStamp := convertStringToTime(reservation.StartTime).Sub(convertStringToTime(lastEndTimeReservation)).Minutes()

			if timeStamp >= 20 {
				return convertStringToTime(lastEndTimeReservation)
			}
		}
		lastEndTimeReservation = reservation.EndTime
	}

	timeStamp := convertStringToTime(lastEndTimeReservation).Sub(convertStringToTime(paramTimeAsString)).Minutes()

	if timeStamp >= 20 {
		return paramTime
	}

	return convertStringToTime(lastEndTimeReservation)
}

func appendZero(value int) string {
	if len(strconv.Itoa(value)) < 2 {
		return fmt.Sprintf("0%v", value)
	}
	return strconv.Itoa(value)
}

func TimeToString(param time.Time) string {

	//hours, minutes, seconds := param.Clock()

	convertedTime := fmt.Sprintf("%v-%v-%v %v:%v:%v",
		param.Year(),
		appendZero(int(param.Month())),
		appendZero(param.Day()),
		appendZero(param.Hour()),
		appendZero(param.Minute()),
		appendZero(param.Second()))
	return convertedTime
}
func convertStringToTime(paramTime string) time.Time {
	resultTime, timeParseError := time.Parse("2006-01-02 15:04:05", strings.TrimSpace(paramTime))

	if timeParseError != nil {
		logg.WriteToFile(fmt.Sprintf("Error occured during parsing: %s", timeParseError))
	}

	return resultTime
}
func roundToMinute(paramTime time.Time) time.Time {
	hour, minute, _ := paramTime.Clock()

	newTime := time.Date(paramTime.Year(), paramTime.Month(), paramTime.Day(), hour, minute, 0, 0, paramTime.Location())

	return newTime
}
