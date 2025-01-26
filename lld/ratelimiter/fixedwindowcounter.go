package ratelimiter

import (
	"sync"
	"time"
)

type fixedWindowCounter struct {
	windowDuration  int   // duration of window
	maxReqInWindow  int  // limit no of request per window
	windowStartTime time.Time // start time of window
	requestCount int  // no of req per window
	mutex sync.Mutex
}

func NewFixedWindowCounter(windowDuration, maxReqInWindow int) *fixedWindowCounter {
	return &fixedWindowCounter{
		windowDuration: windowDuration,
		maxReqInWindow: maxReqInWindow,
		windowStartTime: time.Now(),
		requestCount: 0,
	}
}

func (f *fixedWindowCounter) grantRequest() bool {
	f.mutex.Lock()
	defer f.mutex.Unlock()
    now := time.Now()
	// check if current window is expired
	if now.Sub(f.windowStartTime).Seconds() >= float64(f.windowDuration) {
		// reset the window
		f.windowStartTime = now
		f.requestCount = 0
	}
	// if current req can come in window
	if f.requestCount < f.maxReqInWindow {
		f.requestCount++
		return true
	}
	return false
}