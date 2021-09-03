package goqueu

import (
	"errors"
	"sync"
)

type Node struct {
	data interface{}
	ptr  *Node
}

type Queue struct {
	head *Node
	tail *Node
	len  int32
	lock sync.RWMutex
}

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
		q.tail = node
		q.len++
		return
	}

	q.tail.ptr = node
	q.len++
	return
}

func (q *Queue) Dequeue() (*Node, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.head == nil {
		return nil, errors.New("queue is empty")
	}

	node := q.head
	q.head = q.head.ptr
	return node, nil
}

func (q *Queue) Len() int32 {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.len
}
