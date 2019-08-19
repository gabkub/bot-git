package blacklist

import (
	"bot-git/normalizedDate"
	"bot-git/persistence"
	"crypto/md5"
	"go.etcd.io/bbolt"
	"io"
	"log"
	"sync"
	"time"
)

const blackListDbPath = "./blacklist.db"

type BlackList struct {
	mtx         sync.Mutex
	daysToClean int
	name        string
	db          *persistence.Persistence
}

func New(daysToClean int, name string) *BlackList {
	return &BlackList{
		mtx:         sync.Mutex{},
		db:          persistence.NewPersistence(blackListDbPath, name),
		daysToClean: daysToClean,
		name:        name,
	}
}
func (b *BlackList) getMd5(content *string) string {
	h := md5.New()
	_, err := io.WriteString(h, *content)
	if err != nil {
		log.Println(err)
	}
	return string(h.Sum(nil))
}

func (b *BlackList) IsFresh(content *string) bool {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	hash := b.getMd5(content)
	exists := b.db.HasKey(&hash)
	if exists {
		return false
	}
	b.db.WriteDb(func(bucket *bbolt.Bucket) error {
		d := normalizedDate.NewNormalizeDate(time.Now())
		err := bucket.Put([]byte(hash), d.AsString())
		if err != nil {
			log.Println(err)
		}
		return nil
	})
	return true
}

func (b *BlackList) Clean() {
	b.mtx.Lock()
	defer b.mtx.Unlock()
	apartFromNow := time.Now().AddDate(0, 0, -b.daysToClean)
	b.db.WriteDb(func(bucket *bbolt.Bucket) error {
		err := bucket.ForEach(func(k, v []byte) error {
			t := normalizedDate.ConvertToTime(v)
			if t.Before(apartFromNow) {
				err := bucket.Delete(k)
				if err != nil {
					log.Println(err)
				}
			}
			return nil
		})
		if err != nil {
			log.Println(err)
		}
		return nil
	})
}
