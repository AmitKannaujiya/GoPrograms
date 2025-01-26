package ratelimiter

import (
	"sync"
	"time"
)

type tokenBucket struct {
	tokens      int
	capacity    int
	fillRate    int // fill rate per seconds
	lastUpdated time.Time
	mutex sync.Mutex
}

func NewTokenBucket(tokens, fillRate int) *tokenBucket {
	return &tokenBucket{
		tokens: tokens,
		fillRate: fillRate,
		capacity: tokens,
		lastUpdated: time.Now(),
	}
}

func(t *tokenBucket) grantRequest() bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.fill()
	if t.tokens > 0 {
		t.lastUpdated = time.Now()
		t.tokens--
		return true
	}
	return false
}

func(t *tokenBucket) fill() {
	now := time.Now()
	elapsed := int(now.Sub(t.lastUpdated).Seconds())
	newTokens := elapsed * t.fillRate
	if newTokens > 0 {
		t.tokens = min(t.capacity, t.tokens + newTokens)
	}
}