package blacklists

import (
	"crypto/sha1"
	"sync"
)

type Blist struct {
	Values [][20]byte
	sync.Mutex
}

var MapBL = map[string]Blist{}

func sha(s string) [20]byte {
	return sha1.Sum([]byte(s))
}

func New(name string) {
	for k,_ := range MapBL {
		if k == name {
			return
		}
	}
	MapBL[name] = Blist{}
}


func (b *Blist) Add(s string) {
	b.Lock()
	defer b.Unlock()

	if s == "" {
		return
	}

	b.Values = append(b.Values, sha(s))
}

func (b *Blist) Contains(s string) bool {
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

func (b *Blist) Remove() {
	b.Lock()
	defer b.Unlock()

	l := len(b.Values)

	b.Values = append(b.Values[:l-1])
}


