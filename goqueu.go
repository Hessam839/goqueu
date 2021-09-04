package goqueu

import (
	"errors"
	"sync"
	"sync/atomic"
)

func NewQueue() *Queue {
	return &Queue{
		head: nil,
		tail: nil,
		len:  0,
		lock: sync.RWMutex{},
	}
}

func (q *Queue) Enqueue(data interface{}) {
	//q.lock.Lock()
	//defer q.lock.Unlock()
	node := &Node{data: data}
	var check int64 = 0
	if atomic.CompareAndSwapInt64(&check, 0, 1) {
		defer atomic.StoreInt64(&check, 0)
	} else {
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
	}
}

func (q *Queue) Dequeue() (interface{}, error) {
	//q.lock.Lock()
	//defer q.lock.Unlock()
	var check int64 = 0
	if atomic.CompareAndSwapInt64(&check, 0, 1) {
		defer atomic.StoreInt64(&check, 0)
	} else {
		if q.head == nil {
			return nil, errors.New("queue is empty")
		}

		node := q.head
		q.head = q.head.ptr
		q.len--
		node.ptr = nil
		return node.data, nil
	}
	return nil, nil
}

func (q *Queue) Len() int32 {
	//q.lock.Lock()
	//defer q.lock.Unlock()
	var check int64 = 0
	if atomic.CompareAndSwapInt64(&check, 0, 1) {
		defer atomic.StoreInt64(&check, 0)
	} else {
		return q.len
	}
	return 0
}
