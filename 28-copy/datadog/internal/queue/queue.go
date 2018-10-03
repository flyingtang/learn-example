package queue

import (
	"errors"
	"sync"
)

var ErrorQueueClosed = errors.New("queue is closed")

type Queue struct {
	wg *sync.WaitGroup
	closed chan struct{}
	jobs chan func()
}

func New(capacity int, concurrency int) *Queue {
	if capacity < 1 {
		panic("capacity must be greater than or equal to 1")
	}

	if concurrency < 1 {
		panic("concurrency must be greater than or equal to 1")
	}

	jobs := make(chan func(), capacity -1)
	closed := make(chan struct{})
	wg := &sync.WaitGroup{}

	for i := 0; i < concurrency; i++{
		go worker(wg, jobs)
	}

	return  &Queue{wg:wg, closed:closed, jobs:jobs}
}

func worker(wg *sync.WaitGroup, jobs chan func()){
	for job := range jobs {
		job()
		wg.Done()
	}

}

func (q *Queue) Push(fn func()) error {
	select {
	case <-q.closed: return  ErrorQueueClosed
	default:


	}

	q.wg.Add(1)
	q.jobs <- fn
	return nil
}

func (q *Queue)Wait(){
	q.wg.Done()
}

func (q *Queue)Close(){
	close(q.closed)
	q.wg.Wait()
}