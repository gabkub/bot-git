package persistence

import (
	"fmt"
	bolt "go.etcd.io/bbolt"
	"log"
)

type Persistence struct {
	dbPath     string
	bucketName []byte
}

func NewPersistence(dbPath, bucketName string) *Persistence {
	db := &Persistence{dbPath: dbPath, bucketName: []byte(bucketName)}
	db.createTableDB()
	return db
}

func (p *Persistence) ReadDb(function func(b *bolt.Bucket) error) {
	db := p.openDb(true)
	if db == nil {
		return
	}
	defer db.Close()
	viewError := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(p.bucketName)
		return function(b)
	})
	if viewError != nil {
		log.Println(viewError)
	}
}
func (p *Persistence) WriteDb(function func(b *bolt.Bucket) error) {
	db := p.openDb(false)
	if db == nil {
		return
	}
	defer db.Close()
	updateError := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(p.bucketName)
		return function(b)
	})
	if updateError != nil {
		log.Fatal(fmt.Sprintf("Unable to update. Error: %s", updateError))
	}
}

func (p *Persistence) openDb(readOnly bool) *bolt.DB {
	db, err := bolt.Open(p.dbPath, 0600, &bolt.Options{ReadOnly: readOnly})
	if err != nil {
		if err.Error() == "timeout" {
			log.Fatal("Database already opened.")
		} else {
			log.Fatal("Error opening the database. " + err.Error())
		}
	}
	return db
}

func (p *Persistence) createTableDB() {
	db, err := bolt.Open(p.dbPath, 0600, nil)
	if err != nil {
		log.Fatal("Error opening or creating database.")
	}
	defer db.Close()

	createError := db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists(p.bucketName)
		if err != nil {
			return fmt.Errorf("create bucket error: %s", err)
		}
		return nil
	})
	if createError != nil {
		log.Println("Error creating the table. " + createError.Error())
	}
	log.Printf("Created database table for booking %s\n", p.bucketName)
}

func (p *Persistence) HasKey(key *string) bool {
	var exists bool
	p.ReadDb(func(b *bolt.Bucket) error {
		res := b.Get([]byte(*key))
		if res != nil {
			exists = true
		}
		return nil
	})
	return exists
}
