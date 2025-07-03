package main

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestFixedWindowsRateLimiter_Allow(t *testing.T) {
	limiter := NewRateLimiter(3, time.Second)
	userID := uuid.NewString()

	for range 3 {
		require.True(t, limiter.Allow(userID))
	}

	require.False(t, limiter.Allow(userID))
}

func TestFixedWindowsRateLimiter_Allow_concurrency(t *testing.T) {
	rl := NewFixedWindowRateLimiter(100, time.Second)
	var successCount atomic.Int32
	const total = 1000

	var wg sync.WaitGroup
	for i := 0; i < total; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if rl.Allow("user1") {
				successCount.Add(1)
			}
		}()
	}
	wg.Wait()

	t.Log("Total allowed:", successCount)

	require.LessOrEqual(t, successCount.Load(), int32(100))
}
