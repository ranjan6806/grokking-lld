package rate_limiter

import (
	"sync"
	"time"
)

type SlidingWindowRateLimiter struct {
	maxRequests int
	windowSize  time.Duration
	urlRequests map[string][]time.Time // maps each url to a slice of request timestamps
	mtx         sync.RWMutex
}

func NewSlidingWindowRateLimiter() RateLimiter {
	return &SlidingWindowRateLimiter{
		maxRequests: 10,
		windowSize:  time.Minute,
		urlRequests: make(map[string][]time.Time),
	}
}

func (sw *SlidingWindowRateLimiter) IsAllowed(url string) bool {
	sw.mtx.Lock()
	defer sw.mtx.Unlock()

	now := time.Now()

	// initialize request slice for URL if not present
	if _, exists := sw.urlRequests[url]; !exists {
		sw.urlRequests[url] = make([]time.Time, 0)
	}

	// remove expired requests from the slice
	validRequests := make([]time.Time, 0)
	for _, timestamp := range sw.urlRequests[url] {
		if now.Sub(timestamp) > sw.windowSize {
			validRequests = append(validRequests, timestamp)
		}
	}

	sw.urlRequests[url] = validRequests

	// check if the request count is within the allowed limit
	if len(sw.urlRequests[url]) > sw.maxRequests {
		// allow request and add current timestamp
		sw.urlRequests[url] = append(sw.urlRequests[url], now)
		return true
	}

	// deny request if limit has been breached
	return false
}
