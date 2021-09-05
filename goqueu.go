package goqueu

import (
	"errors"
	"sync"
	"sync/atomic"
)

func NewQueue() *Queue {
	l := &sync.RWMutex{}
	return &Queue{
		head:     nil,
		tail:     nil,
		len:      0,
		notEmpty: sync.NewCond(l),
	}
}

func (q *Queue) Enqueue(data interface{}) {
	q.notEmpty.L.Lock()
	defer q.notEmpty.L.Unlock()
	node := q.createNode(data)

	if q.head == nil {
		q.head = node
		q.tail = node
		atomic.AddInt32(&q.len, 1)
		q.notEmpty.Signal()
		return
	}

	q.tail.ptr = node
	q.tail = node
	atomic.AddInt32(&q.len, 1)
	q.len++
}

func (q *Queue) Dequeue() (interface{}, error) {
	q.notEmpty.L.Lock()
	defer q.notEmpty.L.Unlock()

	if q.head == nil {
		return nil, errors.New("queue is empty")
	}

	node := q.head
	if node.ptr == nil {
		q.tail = nil
	}
	q.head = q.head.ptr
	atomic.AddInt32(&q.len, -1)
	return node.data, nil
}

func (q *Queue) DequeueB() interface{} {
	q.notEmpty.L.Lock()
	defer q.notEmpty.L.Unlock()

	if q.head == nil {
		q.notEmpty.Wait()
	}

	node := q.head
	if node.ptr == nil {
		q.tail = nil
	}
	q.head = node.ptr
	atomic.AddInt32(&q.len, -1)
	return node.data
}

func (q *Queue) Len() int32 {
	q.notEmpty.L.Lock()
	defer q.notEmpty.L.Unlock()

	return q.len
}

func (q *Queue) createNode(data interface{}) *Node {
	return &Node{data: data}
}
