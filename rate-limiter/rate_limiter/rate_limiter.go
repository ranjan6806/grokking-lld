package rate_limiter

import "fmt"

type RateLimiter interface {
	IsAllowed(url string) bool
}

type RateLimiterType string

const (
	TokenBucket   RateLimiterType = "token_bucket"
	FixedWindow   RateLimiterType = "fixed_window"
	SlidingWindow RateLimiterType = "sliding_window"
)

func NewRateLimiter(rateLimiterType RateLimiterType) (RateLimiter, error) {
	switch rateLimiterType {
	case TokenBucket:
		return NewTokenBucketRateLimiter(), nil
	case FixedWindow:
		return NewFixedWindowRateLimiter(), nil
	default:
		return nil, fmt.Errorf("invalid rate limiter: %s", rateLimiterType)
	}
}
