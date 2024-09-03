package queue

import "sync"

type Queue[T any] interface {
	Enqueue(item T)
	Dequeue() (T, bool)
	IsEmpty() bool
}

type ConcurrentQueue[T any] struct {
	items []T
	lock  sync.Mutex
	cond  *sync.Cond
}

func NewConcurrentQueue[T any]() *ConcurrentQueue[T] {
	queue := &ConcurrentQueue[T]{}
	queue.cond = sync.NewCond(&queue.lock)
	return queue
}

func (q *ConcurrentQueue[T]) Enqueue(item T) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.items = append(q.items, item)
	q.cond.Signal()
}

func (q *ConcurrentQueue[T]) Dequeue() (T, bool) {
	q.lock.Lock()
	defer q.lock.Unlock()

	for len(q.items) == 0 {
		q.cond.Wait()
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func (q *ConcurrentQueue[T]) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	return len(q.items) == 0
}
