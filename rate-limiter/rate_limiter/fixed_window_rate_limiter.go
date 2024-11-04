package rate_limiter

import (
	"fmt"
	"sync"
	"time"
)

type FixedWindowRateLimiter struct {
	maxRequests  int
	windowSize   time.Duration
	urlRequests  map[string]int       // track request count per url
	urlResetTime map[string]time.Time // track reset time per url
	mtx          sync.RWMutex
}

func NewFixedWindowRateLimiter() RateLimiter {
	return &FixedWindowRateLimiter{
		maxRequests:  10,
		windowSize:   time.Minute,
		urlRequests:  make(map[string]int),
		urlResetTime: make(map[string]time.Time),
	}
}

func (fw *FixedWindowRateLimiter) IsAllowed(url string) bool {
	fw.mtx.Lock()
	defer fw.mtx.Unlock()

	now := time.Now()

	// check if window for this url has expired
	if resetTime, exists := fw.urlResetTime[url]; !exists || now.After(resetTime) {
		// reset the window for this url
		fmt.Printf("[fixed_window_rate_limiter] [url - %s], resetting window\n", url)
		fw.urlRequests[url] = 0
		fw.urlResetTime[url] = now.Add(fw.windowSize)
	}

	// check if the request count is within the allowed limit
	if fw.urlRequests[url] < fw.maxRequests {
		fmt.Printf("[fixed_window_rate_limiter] [url - %s], allowed\n", url)
		fw.urlRequests[url]++
		return true
	}

	// deny request if limit has been reached
	fmt.Printf("[fixed_window_rate_limiter] [url - %s], blocked\n", url)
	return false
}
