package goqueu

import "sync"

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
