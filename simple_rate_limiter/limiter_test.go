package main

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestMyRateLimiter_Allow(t *testing.T) {

	testCases := []struct {
		Name                string
		UserID              string
		ExpectedAllowed     bool
		RequestCount        int
		RequestCountPerUser int
		Duration            time.Duration
	}{

		{
			Name:                "user1",
			UserID:              uuid.NewString(),
			ExpectedAllowed:     true,
			RequestCount:        5,
			RequestCountPerUser: 5,
			Duration:            1 * time.Second,
		},
		{
			Name:                "user1",
			UserID:              uuid.NewString(),
			ExpectedAllowed:     false,
			RequestCount:        5,
			RequestCountPerUser: 10,
			Duration:            1 * time.Second,
		},
		{
			Name:                "user1",
			UserID:              uuid.NewString(),
			ExpectedAllowed:     false,
			RequestCount:        0,
			RequestCountPerUser: 1,
			Duration:            1 * time.Second,
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {

			rl := NewRateLimiter(test.RequestCount, test.Duration)
			for i := 0; i < test.RequestCountPerUser; i++ {
				allowed := rl.Allow(test.UserID)
				if !allowed {
					t.Errorf("Expected request %d for user %s to be allowed, but it was denied", i+1, test.UserID)
				}
			}

		})
	}
}

func TestMyRateLimiter_Allow_in_time(t *testing.T) {
	limiter := NewRateLimiter(3, time.Second)
	userID := uuid.NewString()

	for range 3 {
		require.True(t, limiter.Allow(userID))
	}

	require.False(t, limiter.Allow(userID))
}
