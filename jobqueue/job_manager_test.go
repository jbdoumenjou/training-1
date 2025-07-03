package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTask_Enqueue(t *testing.T) {
	q := NewTaskQueue(2, 3) // 2 workerCount, 3 retries
	defer q.Stop()

	done := make(chan struct{})

	q.Enqueue("task1", func() error {
		close(done)
		return nil
	})

	q.Start()

	select {
	case <-done:
	case <-time.After(1 * time.Second):
		t.Fatal("Task did not complete")
	}

	assert.Equal(t, StatusDone, q.Status("task1"))

}
