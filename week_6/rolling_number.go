package week_6

import (
	"errors"
	"sync"
	"time"
)

type RollingNumber struct {
	time                     time.Time
	timeInMilliseconds       int64
	numberOfBucket           int
	bucketSizeInMilliseconds int64
	// the index of current bucket
	currentBucketIdx int
	// the list of buckets
	buckets       []*bucket
	newBucketLock sync.Mutex
}

func NewRollingNumber(timeInMilliseconds int64, numberOfBucket int) (*RollingNumber, error) {
	if timeInMilliseconds%int64(numberOfBucket) != 0 {
		return nil, errors.New("The timeInMilliseconds must divide equally into numberOfBuckets.")
	}

	n := &RollingNumber{
		time:                     time.Now(),
		timeInMilliseconds:       timeInMilliseconds,
		numberOfBucket:           numberOfBucket,
		bucketSizeInMilliseconds: timeInMilliseconds / int64(numberOfBucket),
		buckets:                  make([]*bucket, numberOfBucket),
	}
	return n, nil
}

// Add will add one to current bucket
func (number *RollingNumber) Add() {
	number.getCurrentBucket().add()
}

// Reset will set all bucket to nil
func (number *RollingNumber) Reset() {
	number.buckets = make([]*bucket, number.numberOfBucket)
}

// GetRollingSum get the sum for all bucket
func (number *RollingNumber) GetRollingSum() int64 {
	var sum int64 = 0
	for _, b := range number.buckets {
		if b == nil {
			continue
		}
		sum += b.number
	}
	return sum
}

func (number *RollingNumber) GetValueOfLatestBucket() int64 {
	b := number.getCurrentBucket()
	if b == nil {
		return 0
	} else {
		return b.number
	}

}

func (number *RollingNumber) GetRollingMaxValue() int64 {
	var m int64 = 0
	for _, b := range number.buckets {
		if b == nil {
			continue
		} else {
			m = max(m, b.number)
		}
	}
	return m
}

func (number *RollingNumber) getCurrentBucket() *bucket {
	now := time.Now().Unix()
	currentBucket := number.buckets[number.currentBucketIdx]
	if currentBucket != nil && now < currentBucket.windowStart+number.bucketSizeInMilliseconds {
		return currentBucket
	}
	// do not find the bucket, we need to create a new bucket here
	number.newBucketLock.Lock()
	defer number.newBucketLock.Unlock()
	lastBucket := number.buckets[number.currentBucketIdx]

	if lastBucket == nil {
		lastBucket = newBucket(now)
		number.currentBucketIdx += 1
		number.buckets[number.currentBucketIdx] = lastBucket
		return lastBucket
	} else {
		// it means list if full, we need to pop out the first one and push this one in
		b := newBucket(now)
		number.buckets = number.buckets[1:]
		number.buckets = append(number.buckets, b)
		return b
	}
}

func max(a, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}
