package week_6

import "sync"

type bucket struct {
	windowStart int64
	number      int64
	lock        sync.RWMutex
}

func newBucket(startTime int64) *bucket {
	return &bucket{
		windowStart: startTime,
	}
}

func (b *bucket) add() {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.number += 1
}

func (b *bucket) sub() {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.number -= 1
}
