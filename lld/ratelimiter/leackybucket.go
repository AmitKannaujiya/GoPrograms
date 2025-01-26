package ratelimiter

import (
	"sync"
	"time"
)

type Queue[T any] struct {
	enqueue func(T)
	dequeue func() T
	isEmpty func() bool
	size func() int
	peek func() T
}

func NewQueue[T any]() *Queue[T] {
	elements := make([]T, 0)
	return &Queue[T]{
		enqueue: func(t T) {
			elements = append(elements, t)
		},
		dequeue: func() T {
			t := elements[0]
			elements = elements[1:]
			return t
		},
		isEmpty: func() bool {
			return len(elements) == 0
		},
		size: func() int {
			return len(elements)
		},
		peek: func() T {
			return elements[0]
		},
	}
}
type leakyBucket struct {
	capacity int
	leakRate time.Duration
	queue    *Queue[int]
	mutex sync.Mutex
	lastUpdated time.Time
}

func NewLeakyBucket(capacity int, leakRate time.Duration) *leakyBucket {
	return &leakyBucket{
		capacity: capacity,
		leakRate: leakRate,
		queue: NewQueue[int](),
		lastUpdated: time.Now(),
	}
}

func(l *leakyBucket) leak() {
	now :=  time.Now()
	elapsed := now.Sub(l.lastUpdated)
	leaked :=  int(elapsed / l.leakRate)
    // Calculate how many requests to remove based on elapsed time
	for leaked > 0 && !l.queue.isEmpty() {
		l.queue.dequeue()
		leaked--
	}
}

func (l *leakyBucket) grantRequest() bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	// leak requests
	l.leak()
	now := time.Now()	
	// if still some space to request
	if l.queue.size() < l.capacity {
		l.lastUpdated = now
		// add request in queue
		l.queue.enqueue(1)
		return true
	}
	// request denied
	return false
}