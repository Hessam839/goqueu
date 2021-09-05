package goqueu

import (
	"log"
	"math/rand"
	"testing"
	"time"
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
	b.SetParallelism(10000)
	b.ResetTimer()
	//for i:=0 ;i<b.N; i++{
	b.Run("enqueue", func(b *testing.B) {
		b.ReportAllocs()
		b.SetParallelism(10000)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				q.Enqueue("hello world")
			}
		})
	})
	b.Run("dequeue", func(b *testing.B) {
		b.ReportAllocs()
		b.SetParallelism(10000)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_, _ = q.Dequeue()
			}
		})
	})
	//}
}

func Test_Dequeueb(t *testing.T) {
	q := NewQueue()
	rand.Seed(time.Now().UTC().UnixNano())

	log.Println("test started ...")

	go func() {
		//timer := time.NewTicker(time.Second * time.Duration(rand.Intn(3)))
		//for range timer.C{
		//	q.Enqueue(rand.Intn(30))
		//}
		for {
			rnd := rand.Intn(30)
			q.Enqueue(rnd)
			//log.Printf("produce Value %d", rnd)
			time.Sleep(time.Second)
		}
	}()

	log.Println("start consumer ...")
	for i := 0; i < 3; i++ {
		go func(c int) {
			for {
				data := q.DequeueB().(int)
				log.Printf("consumer %d data is: %v", c, data)
				time.Sleep(time.Millisecond * 100)
			}
		}(i)
	}
	select {}
}
