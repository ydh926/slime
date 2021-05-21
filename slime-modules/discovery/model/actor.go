package model

import (
	"sync"
)

type Actor interface {
	Start()
	Stop()
	Send(event *Event) bool
}

type MailBox interface {
	Put(event *Event)
	Next() *Event
	IsOutDate(event *Event) bool
	IsEmpty() bool
}

// 按版本排序的事件处理队列
type VersionedMailBox struct {
	CurVersion int64
	queue      []*Event
	queueLock  sync.Mutex
}

func NewVersionedMailBox() *VersionedMailBox {
	return &VersionedMailBox{
		queue:     make([]*Event, 0),
		queueLock: sync.Mutex{},
	}
}

func (v *VersionedMailBox) IsEmpty() bool {
	return len(v.queue) == 0
}

func (v *VersionedMailBox) IsOutDate(event *Event) bool {
	return event.Version < v.CurVersion
}

func (v *VersionedMailBox) Put(event *Event) {
	v.queueLock.Lock()
	defer v.queueLock.Unlock()
	v.queue = append(v.queue, event)
	v.siftup(len(v.queue) - 1)
}

func (v *VersionedMailBox) Next() *Event {
	v.queueLock.Lock()
	defer v.queueLock.Unlock()
	if len(v.queue) == 0 {
		return nil
	}
	next := v.queue[0]
	v.queue[0], v.queue[len(v.queue)-1] = v.queue[len(v.queue)-1], v.queue[0]
	v.queue = v.queue[:len(v.queue)-1]
	v.siftdown(0)
	v.CurVersion = next.Version
	return next
}

func (v *VersionedMailBox) siftup(i int) {
	p := (i - 1) / 2
	for i > 0 {
		if v.queue[p].Version > v.queue[i].Version {
			v.queue[i], v.queue[p] = v.queue[p], v.queue[i]
			i = p
			p = (i - 1) / 2
		} else {
			break
		}
	}
}

func (v *VersionedMailBox) siftdown(i int) {
	l, r, min := 2*i+1, 2*i+2, i
	for i < len(v.queue) {
		if l < len(v.queue) && (v.queue[l].Version < v.queue[min].Version) {
			min = l
		}
		if r < len(v.queue) && (v.queue[r].Version < v.queue[min].Version) {
			min = r
		}
		if min != i {
			v.queue[min], v.queue[i] = v.queue[i], v.queue[min]
			i = min
			l, r, min = 2*i+1, 2*i+2, i
		} else {
			break
		}
	}
}
