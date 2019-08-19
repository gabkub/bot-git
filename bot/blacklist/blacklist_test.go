package blacklist_test

import (
	"bot-git/bot/blacklist"
	"bot-git/normalizedDate"
	"bot-git/persistence"
	"crypto/md5"
	"github.com/stretchr/testify/assert"
	"go.etcd.io/bbolt"
	"io"
	"os"
	"testing"
	"time"
)

func TestShouldAddItem(t *testing.T) {
	const dbPath = "./blacklist.db"
	const name = "test"
	db := persistence.NewPersistence(dbPath, name)
	defer os.Remove(dbPath)
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
	const dbPath = "./blacklist.db"
	const name = "test1"
	db := persistence.NewPersistence(dbPath, name)
	defer os.Remove(dbPath)
	b := blacklist.New(1, name)

	c := "a"
	c2 := "a1"
	c3 := "a2"
	b.IsFresh(&c)
	b.IsFresh(&c2)
	b.IsFresh(&c3)

	assert.Equal(t, 3, getCount(db))

	db.WriteDb(func(b *bbolt.Bucket) error {
		b.ForEach(func(k, v []byte) error {
			d := normalizedDate.NewNormalizeDate(time.Now().AddDate(0, 0, -2))
			b.Put(k, d.AsString())
			return nil
		})
		return nil
	})

	b.Clean()

	assert.Equal(t, 0, getCount(db))
}

func getMd5(content *string) string {
	h := md5.New()
	io.WriteString(h, *content)
	return string(h.Sum(nil))
}

func getCount(db *persistence.Persistence) int {
	count := 0
	db.ReadDb(func(b *bbolt.Bucket) error {
		b.ForEach(func(k, v []byte) error {
			count++
			return nil
		})
		return nil
	})
	return count
}
