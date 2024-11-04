package main

import (
	"fmt"
	"rate-limiter/rate_limiter"
)

func main() {
	//tokenBucketRateLimiter, err := rate_limiter.NewRateLimiter(rate_limiter.TokenBucket)
	//if err != nil {
	//	fmt.Printf("error creating rate limiter: %s\n", err)
	//}

	url := "www.google.com"
	//for i := 0; i < 15; i++ {
	//	tokenBucketRateLimiter.IsAllowed(url)
	//}

	fixedWindowRateLimiter, err := rate_limiter.NewRateLimiter(rate_limiter.FixedWindow)
	if err != nil {
		fmt.Printf("error creating rate limiter: %s\n", err)
	}

	for i := 0; i < 15; i++ {
		fixedWindowRateLimiter.IsAllowed(url)
	}
}
