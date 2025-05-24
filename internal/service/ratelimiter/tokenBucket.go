package service

import (
	"sync"
	"time"

	"github.com/labstack/gommon/log"
)

type Bucket struct {
	max           int
	refillRate    int
	currentTokens int
	refillPeriod  time.Duration
	lastRefill    time.Time
	mutex         sync.Mutex
}

func NewTokenBucket(maxTokens, refillRate int, refillPeriod time.Duration) *Bucket {
	return &Bucket{
		max:           maxTokens,
		currentTokens: maxTokens,
		refillRate:    refillRate,
		refillPeriod:  refillPeriod,
		lastRefill:    time.Now(),
	}
}

func (bucket *Bucket) CheckIsAllowed() bool {
	bucket.mutex.Lock()
	defer bucket.mutex.Unlock()

	now := time.Now()
	elapsedTime := now.Sub(bucket.lastRefill)

	log.Info("elapsedTime ", elapsedTime)
	if elapsedTime >= bucket.refillPeriod {
		bucket.refill(elapsedTime)
	}

	if bucket.currentTokens > 0 {
		bucket.currentTokens -= 1
		log.Info("token removed. Left: ", bucket.currentTokens)
		return true
	}

	log.Info("request should be dropped")
	return false
}

func (bucket *Bucket) refill(elapsedTime time.Duration) {

	log.Info("refilling bucket ")
	tokensToAdd := int(elapsedTime/bucket.refillPeriod) * bucket.refillRate
	if bucket.currentTokens+tokensToAdd > bucket.max {
		bucket.currentTokens = bucket.max
	} else {
		bucket.currentTokens += tokensToAdd
	}
	bucket.lastRefill = time.Now()
}
