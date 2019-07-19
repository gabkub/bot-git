package tableDB

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"strings"
	"time"
)

func CreateTableDB(){
	db, err := bolt.Open("footballTable.db", 0600, nil)

	if err != nil{
		log.Fatal(err)
	}

	createError := db.Update(func(tx *bolt.Tx) error{
		_, err = tx.CreateBucketIfNotExists([]byte(strings.TrimSpace("RESERVATION")))

		if err != nil{
			return fmt.Errorf("create bucket error: %s", err)
		}

		return nil
	})

	if createError != nil{
		//db.Close()
		log.Fatal(createError)
	}

	//defer db.Close()
}
type TimeReservation struct{
	userName	string
	startTime	string
	endTime		string
}

func openConnectionDB(readOnly bool) *bolt.DB{

	if readOnly{
		db, err := bolt.Open("footballTable.db",0600,&bolt.Options{ReadOnly: true})

		if err != nil{
			log.Fatal(err)
		}

		return db
	}

	db, err := bolt.Open("footballTable.db", 0600, nil)

	if err != nil{
		log.Fatal(err)
	}

	return db
}
func closeConnection(database *bolt.DB){
	err := database.Close()

	if err != nil{
		log.Printf("Error occurred during closing database connection: %s", err)
	}
}
func convertOnlyTimeToString(param time.Time) string{

	hours, minutes, seconds := param.Clock()
	convertedTime := fmt.Sprintf("%s:%s:%s",hours,minutes,seconds)

	return convertedTime
}
func convertStringToTime(paramTime string) time.Time{
	resultTime, timeParseError := time.Parse("150405", strings.TrimSpace(paramTime))

	if timeParseError != nil{
		log.Fatal(fmt.Sprintf("Error occured during parsing: %s", timeParseError))
	}

	return resultTime
}

func SetReservation(userName string, startTime time.Time){
	db := openConnectionDB(false)
	defer closeConnection(db)

	tempTime := convertOnlyTimeToString(startTime)

	updateError := db.Update(func(tx *bolt.Tx) error{
		b := tx.Bucket([]byte(strings.TrimSpace("RESERVATION")))
		err := b.Put([]byte(strings.TrimSpace(userName)), []byte(strings.TrimSpace(tempTime)))
		return err
	})

	if updateError != nil{
		log.Fatal(fmt.Sprintf("Unable to update. Error: %s", updateError))
	}
}
func GetReservationByUserName(userName string) TimeReservation{
	db := openConnectionDB(true)

	defer  closeConnection(db)


	var result TimeReservation
	result.userName = userName

	err := db.View(func(tx *bolt.Tx) error{
		b := tx.Bucket([]byte("RESERVATION"))
		v := b.Get([]byte(strings.TrimSpace(userName)))

		if v == nil {
			return fmt.Errorf("value is empty to key: %s", string(v))
		}

		result.startTime = string(v)
		return nil
	})

	if err != nil{
		log.Println("Error occurred during get reservation by username.")
		log.Println(err)

		return TimeReservation{userName,"",""}
	}

	startTime := convertStringToTime(result.startTime)

	result.endTime = convertOnlyTimeToString(startTime.Add(time.Minute * 20))

	return result
}
func GetAllReservationByStartTime(startTime time.Time) []TimeReservation{
	db := openConnectionDB(true)

	defer closeConnection(db)


	var reservations []TimeReservation
	err := db.View(func(tx *bolt.Tx) error{
		bucket := tx.Bucket([]byte("RESERVATION"))

		searchingError := bucket.ForEach(func(username, start []byte) error {

			if strings.Compare(convertOnlyTimeToString(startTime), string(start)) == 1{
				reservation := TimeReservation{
					userName:	string(username),
					startTime:	string(start),
					endTime:	convertOnlyTimeToString(convertStringToTime(string(start)).Add(time.Minute * 20)),
				}
				reservations = append(reservations,reservation)
				return nil
			}
			return nil
		})

		if searchingError != nil{
			return searchingError
		}

		return nil
	})

	if err != nil{
		log.Println(err)
	}

	return reservations
}

