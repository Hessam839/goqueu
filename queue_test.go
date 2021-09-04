package goqueu

import (
	"log"
	"testing"
)

func Test_Enqueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue("hello world")
	q.Enqueue("my name is hessam")
	q.Enqueue(12)
	q.Enqueue("hello world")
	q.Enqueue("my name is hessam")
	q.Enqueue(12)

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

func BenchmarkEnqueue(b *testing.B) {
	q := NewQueue()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Enqueue("hello world")
	}
}

func BenchmarkEnqueue_Parallel(b *testing.B) {
	q := NewQueue()
	b.ReportAllocs()
	b.SetParallelism(1000)
	b.ResetTimer()
	//for i:=0 ;i<b.N; i++{
	b.Run("enqueue", func(b *testing.B) {
		b.ReportAllocs()
		b.SetParallelism(1000)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				q.Enqueue("hello world")
			}
		})
	})
	b.Run("dequeue", func(b *testing.B) {
		b.ReportAllocs()
		b.SetParallelism(1000)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_, _ = q.Dequeue()
			}
		})
	})
	//}
}
