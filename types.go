package goqueu

import "sync"

type Node struct {
	data interface{}
	ptr  *Node
}

type Queue struct {
	poll     *sync.Pool
	head     *Node
	tail     *Node
	len      int32
	notEmpty *sync.Cond
}
