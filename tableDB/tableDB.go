package tableDB

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"strings"
	"time"
)

func CreateTableDB(bucketName string){
	db, err := bolt.Open("footballTable.db", 0600, nil)

	if err != nil{
		log.Fatal(err)
	}

	createError := db.Update(func(tx *bolt.Tx) error{
		_, err = tx.CreateBucketIfNotExists([]byte(strings.TrimSpace(bucketName)))

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
func SetReservation(bucketName, userName string, startTime time.Time){

	//db, err := bolt.Open("fotballTable.db", 0600, nil)
	//
	//if err != nil{
	//	log.Fatal(err)
	}

	//hour, minute, second := startTime.Clock()
	//tempTime

	//db.Update(func(tx *bolt.Tx) error{
	//	b := tx.Bucket([]byte(strings.TrimSpace(bucketName)))
	//	err := b.Put([]byte(strings.TrimSpace(userName)), []byte(startTime))
	//})

}
