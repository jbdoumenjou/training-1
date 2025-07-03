package main

import (
	"sync"
	"time"
)

type FixedWindowRateLimiter struct {
	limit     int            // Maximum number of requests allowed in the window
	window    time.Duration  // Duration of the rate limit window
	requests  map[string]int // Map to track requests per user
	lastReset time.Time      // Last time the rate limit window was reset
	mu        sync.Mutex     // Mutex to protect concurrent access
}

func NewFixedWindowRateLimiter(limit int, window time.Duration) *FixedWindowRateLimiter {
	return &FixedWindowRateLimiter{
		limit:     limit,
		window:    window,
		requests:  make(map[string]int),
		lastReset: time.Now(),
	}
}

// Very naive approach to rate limiting using fixed windows
// does not support concurrent requests
func (f *FixedWindowRateLimiter) Allow(userID string) bool {
	now := time.Now()
	f.mu.Lock()
	defer f.mu.Unlock()

	// Reset the window if it has expired
	if now.Sub(f.lastReset) >= f.window {
		f.requests = make(map[string]int)
		f.lastReset = now
	}

	// Increment the request count for the user
	f.requests[userID]++

	// Check if the user has exceeded the limit
	if f.requests[userID] > f.limit {
		return false // Deny the request
	}

	return true // Allow the request
}
