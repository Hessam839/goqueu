package goqueu

import (
	"errors"
	"sync"
)

func NewQueue() *Queue {
	return &Queue{
		head: nil,
		tail: nil,
		len:  0,
		lock: sync.RWMutex{},
	}
}

func (q *Queue) Enqueue(node *Node) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.head == nil {
		q.head = node
		//q.head.ptr = nil
		q.tail = node
		//q.tail.ptr = nil
		q.len++
		return
	}

	q.tail.ptr = node
	q.tail = node
	q.len++
	return
}

func (q *Queue) Dequeue() (interface{}, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.head == nil {
		return nil, errors.New("queue is empty")
	}

	node := q.head
	q.head = q.head.ptr
	q.len--
	node.ptr = nil
	return node.data, nil
}

func (q *Queue) Len() int32 {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.len
}
