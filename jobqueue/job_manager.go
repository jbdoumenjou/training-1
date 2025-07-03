package main

import "sync"

type TaskStatus string

const (
	StatusPending TaskStatus = "pending"
	StatusRunning TaskStatus = "running"
	StatusFailed  TaskStatus = "failed"
	StatusDone    TaskStatus = "done"
)

type TaskQueue interface {
	Enqueue(id string, fn func() error)
	Status(id string) TaskStatus
	Start()
	Stop()
}

type Task struct {
	id       string
	fn       func() error
	attempts int
}

type Queue struct {
	workerCount int
	maxAttempts int

	task     chan Task
	statuses map[string]TaskStatus
	statusMu sync.Mutex

	wg     sync.WaitGroup
	closed chan struct{}
}

func NewTaskQueue(workers int, maxAttempts int) Queue {
	return Queue{
		workerCount: workers,
		maxAttempts: maxAttempts,
		task:        make(chan Task),
		statuses:    make(map[string]TaskStatus),
		closed:      make(chan struct{}),
	}
}

func (q *Queue) Enqueue(id string, fn func() error) {
	q.statusMu.Lock()
	q.statuses[id] = StatusPending
	q.statusMu.Unlock()

	q.task <- Task{id: id, fn: fn}
}

func (q *Queue) Start() {
	for range q.workerCount {
		q.wg.Add(1)
		go q.worker()
	}
}

func (q *Queue) worker() {
	defer q.wg.Done()

	for {
		select {
		case <-q.closed:
			return
		case task := <-q.task:
			q.runTask(task)
		}
	}
}

func (q *Queue) setStatus(id string, status TaskStatus) {
	q.statusMu.Lock()
	defer q.statusMu.Unlock()
	q.statuses[id] = status
}

func (q *Queue) runTask(t Task) {
	q.setStatus(t.id, StatusRunning)

	for attempt := 0; attempt <= q.maxAttempts; attempt++ {
		err := t.fn()
		if err == nil {
			q.setStatus(t.id, StatusDone)
			return
		}
	}

	q.setStatus(t.id, StatusFailed)
}

func (q *Queue) Stop() {
	close(q.closed)
	q.wg.Wait()
	close(q.task)
}

func (q *Queue) Status(id string) TaskStatus {
	q.statusMu.Lock()
	defer q.statusMu.Unlock()

	return q.statuses[id]
}
