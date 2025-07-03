package main

import (
	"sync"
	"time"
)

type UserBucket2 struct {
	tokens     int // Number of tokens available in the bucket
	lastAccess time.Time
}

type TokenBucketRateLimiter struct {
	limit   int                     // Maximum number of tokens in the bucket
	rate    time.Duration           // Time duration to refill the bucket
	mu      sync.Mutex              // Mutex to protect concurrent access
	buckets map[string]*UserBucket2 // Map to track user buckets
}

func NewBucketRateLimiter(limit int, rate time.Duration) *TokenBucketRateLimiter {
	return &TokenBucketRateLimiter{
		limit:   limit,
		rate:    rate,
		buckets: make(map[string]*UserBucket2),
	}
}

func (rl *TokenBucketRateLimiter) Allow(userID string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	bucket, exists := rl.buckets[userID]
	if !exists {
		rl.buckets[userID] = &UserBucket2{tokens: rl.limit - 1, lastAccess: now}
		return true
	}

	// refill tokens
	elapsed := now.Sub(bucket.lastAccess)
	refill := int(elapsed / rl.rate)
	if refill > 0 {
		bucket.tokens += refill
		if bucket.tokens > rl.limit {
			bucket.tokens = rl.limit
		}
		bucket.lastAccess = now
	}

	if bucket.tokens > 0 {
		bucket.tokens--
		return true
	}

	return false
}
