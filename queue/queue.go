package queue

import (
	"sync"
	"time"
)

type Queue struct {
	mu                                  sync.Mutex
	max_size, current_size, write, read int
	queue                               []string
}

func (q *Queue) Enqueue(s string) {

	q.mu.Lock()         // lock mutex
	defer q.mu.Unlock() // unlock mutex after func finishes

	// lock goroutine while queue is full
	for q.current_size == q.max_size {
		q.mu.Unlock()                    // unlock mutex until condition is done
		time.Sleep(1 * time.Microsecond) // sleep for 1 microsecond
		q.mu.Lock()                      // reacquire mutex
	}

	// write string to queue and update indexes
	q.queue[q.write] = s
	q.write++
	q.current_size++
	if q.write == q.max_size {
		q.write = 0
	}
}

func (q *Queue) Dequeue() string {

	q.mu.Lock()         // lock mutex
	defer q.mu.Unlock() // unlock mutex after func finishes

	// lock goroutine while queue is empty
	for q.current_size == 0 {
		q.mu.Unlock()                    // unlock mutex until condition is done
		time.Sleep(1 * time.Microsecond) // sleep for 1 microsecond
		q.mu.Lock()                      // reacquire mutex
	}

	// read string from queue and update indexes
	result := q.queue[q.read]
	q.read++
	if q.read == q.max_size {
		q.read = 0
	}

	return result
}
