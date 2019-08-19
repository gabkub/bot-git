package blacklist_test

import (
	"bot-git/bot/blacklist"
	"bot-git/normalizedDate"
	"bot-git/persistence"
	"crypto/md5"
	"github.com/stretchr/testify/assert"
	"go.etcd.io/bbolt"
	"io"
	"log"
	"os"
	"testing"
	"time"
)

const dbPath = "./blacklist.db"

func TestShouldAddItem(t *testing.T) {
	const name = "test"
	db, f := setup(t, name)
	defer f()

	b := blacklist.New(1, name)

	c := "a"
	hash := getMd5(&c)
	h := db.HasKey(&hash)
	assert.False(t, h)
	r := b.IsFresh(&c)
	assert.True(t, r)

	r = b.IsFresh(&c)
	assert.False(t, r)
	h = db.HasKey(&hash)
	assert.True(t, h)
}

func TestShouldCleanItems(t *testing.T) {
	const name = "test1"
	db, f := setup(t, name)
	defer f()

	b := blacklist.New(1, name)

	c := "a"
	c2 := "a1"
	c3 := "a2"
	b.IsFresh(&c)
	b.IsFresh(&c2)
	b.IsFresh(&c3)

	assert.Equal(t, 3, getCount(db))

	db.WriteDb(func(b *bbolt.Bucket) error {
		err := b.ForEach(func(k, v []byte) error {
			d := normalizedDate.NewNormalizeDate(time.Now().AddDate(0, 0, -2))
			err := b.Put(k, d.AsString())
			if err != nil {
				log.Println(err)
			}
			return nil
		})
		if err != nil {
			log.Println(err)
		}
		return nil
	})

	b.Clean()

	assert.Equal(t, 0, getCount(db))
}

func setup(t *testing.T, name string) (*persistence.Persistence, func()) {
	db := persistence.NewPersistence(dbPath, name)
	return db, func() {
		err := os.Remove(dbPath)
		assert.Nil(t, err)
	}
}

func getMd5(content *string) string {
	h := md5.New()
	_, err := io.WriteString(h, *content)
	if err != nil {
		log.Println(err)
	}
	return string(h.Sum(nil))
}

func getCount(db *persistence.Persistence) int {
	count := 0
	db.ReadDb(func(b *bbolt.Bucket) error {
		err := b.ForEach(func(k, v []byte) error {
			count++
			return nil
		})
		if err != nil {
			log.Println(err)
		}
		return nil
	})
	return count
}
