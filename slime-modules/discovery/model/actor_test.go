package model

import (
	"math/rand"
	"sort"
	"testing"
)

func TestMailBoxPut(t *testing.T) {
	r := make([]*Event, 0)
	s := make([]int, 0)
	for i := 0; i < 100; i++ {
		k := rand.Int()
		r = append(r, &Event{Version: int64(k)})
		s = append(s, k)
	}
	sort.Ints(s)
	m := VersionedMailBox{
		queue: make([]*Event, 0),
	}
	for _, e := range r {
		m.Put(e)
	}
	for i := range s {
		v := m.Next().Version
		if  v!= int64(s[i]) {
			t.Fatalf("exceptedP:%d, but got: %d", v,s[i])
		}
	}
}
