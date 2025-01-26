package ratelimiter

import "time"

type slidingWindowLog struct {
	maxReqInWindow  int    // max request allowed in window
	windowSize time.Duration    // size of sliding window in seconds
	requestLog      *Queue[time.Time]  // log of req timestamp
}

func NewSlidingWindowLog(maxReqInWindow int, windowSizeInSec time.Duration) *slidingWindowLog {
	return &slidingWindowLog{
		maxReqInWindow: maxReqInWindow,
		windowSize: windowSizeInSec,
		requestLog: NewQueue[time.Time](),
	}
}

func (s *slidingWindowLog) grantRequest() bool {
	now := time.Now()
	// get the time which are for the previous window
	cutoffTime := now.Add(-s.windowSize)
	// remove the timestamps out side the window
	for !s.requestLog.isEmpty() && s.requestLog.peek().Before(cutoffTime) {
		s.requestLog.dequeue()
	}
	// check request can be allowed
	if s.requestLog.size() < s.maxReqInWindow {
		s.requestLog.enqueue(now)
		return true
	}
	return false
}