package Queue

import (
	"fmt"
	"sync"
)

type items []interface {}

type Queue struct {
	items items
	lock  sync.Mutex
}

func NewQueue(size int64) *Queue {
	return &Queue{
		items: make([]interface{}, size),
	}
}

func (q *Queue) Enqueue(item interface{}) error {

	if item == nil {
		return fmt.Errorf("item had no value")
	}

	q.lock.Lock()
	q.items = append(q.items, item)
	q.lock.Unlock()
	return nil
}

func (q *Queue) Dequeue() interface{} {
	var ret interface{}
	if !q.IsEmpty() {
		q.lock.Lock()
		ret = q.items[0]
		q.items = append(q.items[:0], q.items[1:]...)
		q.lock.Unlock()
	}
	return ret
}

func (q *Queue) Peek() interface{} {
	var ret interface{}
	if !q.IsEmpty() {
		q.lock.Lock()
		ret = q.items[0]
		q.lock.Unlock()

	}
	return ret
}

func (q *Queue) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return len(q.items) == 0
}

