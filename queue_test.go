package goqueu

import (
	"log"
	"testing"
)

func Test_Enqueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(&Node{data: "hello world", ptr: nil})
	q.Enqueue(&Node{data: "my name is hessam", ptr: nil})
	q.Enqueue(&Node{data: 12, ptr: nil})

	var i int32
	for i = 0; i < q.Len(); i++ {
		node, err := q.Dequeue()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("node data: %v", node)
	}

}
