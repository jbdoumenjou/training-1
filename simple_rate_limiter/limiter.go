package main

import (
	"sync"
	"time"
)

type UserBucket struct {
	allowedRequests int       // Number of requests allowed in the current duration
	lastRequestTime time.Time // Time of the last request
}

type MyRateLimiter struct {
	cache        map[string]*UserBucket
	duration     time.Duration
	requestCount int // Maximum number of requests allowed in the duration

	// mu is used to protect the cache from concurrent access
	mu sync.Mutex
}

func NewRateLimiter(requestCount int, duration time.Duration) *MyRateLimiter {
	return &MyRateLimiter{
		cache:        make(map[string]*UserBucket),
		requestCount: requestCount,
		duration:     duration,
	}

}

func (rl *MyRateLimiter) Allow(userID string) bool {
	bucket, ok := rl.cache[userID]
	if !ok {
		// Create a new bucket for the user if it doesn't exist
		bucket = &UserBucket{
			allowedRequests: rl.requestCount,
		}
		rl.cache[userID] = bucket

		return true // Allow the first request
	}
	rl.mu.Lock()
	defer rl.mu.Unlock()
	if bucket.allowedRequests > 0 {
		bucket.allowedRequests-- // Decrement the allowed requests
		return true              // Allow the request
	}

	// If no requests are allowed, check if the duration has passed
	if time.Since(bucket.lastRequestTime) >= rl.duration {
		bucket.allowedRequests = rl.requestCount - 1 // Reset allowed requests, minus the current one
		bucket.lastRequestTime = time.Now()          // Update the last request time
		return true                                  // Allow the request
	}

	return false // Deny the request
}
