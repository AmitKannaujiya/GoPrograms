package ratelimiter

import (
	"fmt"
	"testing"
	"time"
)

func TestRateLimiterTokenBucket(t *testing.T) {
	tokenBucket := NewTokenBucket(3, 3)

	for i:=0; i < 15; i++ {
		if tokenBucket.grantRequest() {
			fmt.Printf("request : %d, granted \n", i + 1)
		} else {
			fmt.Printf("request : %d, denied \n", i + 1)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
func TestRateLimiterLeakyBucket(t *testing.T) {
	leakyBucket := NewLeakyBucket(3, 1 * time.Second)

	for i:=0; i < 15; i++ {
		if leakyBucket.grantRequest() {
			fmt.Printf("request : %d, granted \n", i + 1)
		} else {
			fmt.Printf("request : %d, denied \n", i + 1)
		}
		time.Sleep(400 * time.Millisecond)
	}
}

func TestRateLimiterFixedWindowCounter(t *testing.T) {
	fixedWindowCounter := NewFixedWindowCounter(3, 3)
	for i:=0; i < 15; i++ {
		if fixedWindowCounter.grantRequest() {
			fmt.Printf("request : %d, granted \n", i + 1)
		} else {
			fmt.Printf("request : %d, denied \n", i + 1)
		}
		time.Sleep(400 * time.Millisecond)
	}
}

func TestRateLimiterSlidingWindowLog(t *testing.T) {
	slidingWindowLog := NewSlidingWindowLog(3, 2 * time.Second)
	for i:=0; i < 15; i++ {
		if slidingWindowLog.grantRequest() {
			fmt.Printf("request : %d, granted \n", i + 1)
		} else {
			fmt.Printf("request : %d, denied \n", i + 1)
		}
		time.Sleep(400 * time.Millisecond)
	}
}