package goqueu

import (
	"log"
	"testing"
)

func Test_Enqueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(&Node{data: "hello world"})
	q.Enqueue(&Node{data: "my name is hessam"})
	q.Enqueue(&Node{data: 12})
	q.Enqueue(&Node{data: "hello world"})
	q.Enqueue(&Node{data: "my name is hessam"})
	q.Enqueue(&Node{data: 12})

	l := q.Len()
	log.Printf("queue length: %d", l)
	for i := int32(0); i < l; i++ {
		node, err := q.Dequeue()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("node data: %v", node)
	}

}
