package blacklist

import (
	"crypto/md5"
	"io"
	"sync"
	"time"
)

type BlackList struct {
	mtx         sync.Mutex
	list        map[string]time.Time
	daysToClean int
}

func New(daysToClean int) *BlackList {
	return &BlackList{
		mtx:         sync.Mutex{},
		list:        map[string]time.Time{},
		daysToClean: daysToClean,
	}
}
func (b *BlackList) getMd5(content *string) string {
	h := md5.New()
	io.WriteString(h, *content)
	return string(h.Sum(nil))
}

func (b *BlackList) IsFresh(content *string) bool {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	hash := b.getMd5(content)
	_, ok := b.list[hash]
	const wasSent = false
	if ok {
		return wasSent
	}
	b.list[hash] = time.Now()
	return true
}

func (b *BlackList) Clean() {
	b.mtx.Lock()
	defer b.mtx.Unlock()
	monthFromNow := time.Now().AddDate(0, 0, -b.daysToClean)
	for key, val := range b.list {
		if val.Before(monthFromNow) {
			delete(b.list, key)
		}
	}
}
