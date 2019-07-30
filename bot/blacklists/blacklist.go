package blacklists

import (
	"crypto/sha1"
	"sync"
)

type Blacklist struct {
	Values [][20]byte
	sync.Mutex
}

var BlacklistsMap = map[string]Blacklist{}

func sha(s string) [20]byte {
	return sha1.Sum([]byte(s))
}

func New(name string) {
	for k,_ := range BlacklistsMap {
		if k == name {
			return
		}
	}
	BlacklistsMap[name] = Blacklist{}
}

func (b *Blacklist) AddElement(s string) {
	b.Lock()
	defer b.Unlock()

	if s == "" {
		return
	}

	b.Values = append(b.Values, sha(s))
}

func (b *Blacklist) Contains(s string) bool {
	b.Lock()
	defer b.Unlock()

	shaS := sha(s)
	for _,v := range b.Values {
		if v == shaS {
			return true
		}
	}
	return false
}