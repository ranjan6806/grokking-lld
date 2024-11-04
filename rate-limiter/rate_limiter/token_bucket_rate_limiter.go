package rate_limiter

import (
	"fmt"
	"sync"
	"time"
)

type UrlBucket struct {
	tokens         int
	lastRefillTime time.Time
}

type TokenBucketRateLimiter struct {
	maxTokens  int // maximum tokens a user bucket can hold
	refillRate int // number of tokens added per minute
	urlBuckets map[string]*UrlBucket
	mtx        sync.RWMutex
}

func NewTokenBucketRateLimiter() RateLimiter {
	return &TokenBucketRateLimiter{
		maxTokens:  10,
		refillRate: 5,
		urlBuckets: make(map[string]*UrlBucket),
	}
}

func (tb *TokenBucketRateLimiter) IsAllowed(url string) bool {
	tb.mtx.Lock()
	defer tb.mtx.Unlock()

	//fmt.Printf("[token_bucket_rate_limiter] [url - %s]\n", url)
	userBucket, exists := tb.urlBuckets[url]
	if !exists {

		fmt.Printf("[token_bucket_rate_limiter] [url - %s] initialising bucket\n", url)

		userBucket = &UrlBucket{
			tokens:         tb.maxTokens,
			lastRefillTime: time.Now(),
		}
		tb.urlBuckets[url] = userBucket
	}

	now := time.Now()
	elapsed := now.Sub(userBucket.lastRefillTime).Minutes()
	refillTokens := int(elapsed) * tb.refillRate

	if refillTokens > 0 {
		userBucket.tokens += refillTokens
		if userBucket.tokens >= tb.maxTokens {
			userBucket.tokens = tb.maxTokens
		}
		userBucket.lastRefillTime = now
	}

	if userBucket.tokens > 0 {
		fmt.Printf("[token_bucket_rate_limiter] [url - %s] allowed\n", url)
		userBucket.tokens--
		return true
	}

	fmt.Printf("[token_bucket_rate_limiter] [url - %s] blocked\n", url)

	return false
}
