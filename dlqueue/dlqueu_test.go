package dlqueue

import (
	"log"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"
)

func TestDlQueue(t *testing.T) {
	q := NewDlQueue()

	q.Put("hello world")

	log.Printf("data is: %v", q.Get())
}

func Test_Blocking(t *testing.T) {
	q := NewDlQueue()

	go func() {
		timer := time.NewTicker(time.Second)

		for range timer.C {
			q.Put(rand.Intn(30))
		}
	}()

	time.Sleep(time.Millisecond)

	for i := 0; i < 3; i++ {
		go func(id int) {
			for {
				data := q.TryGet()
				log.Printf("goroutine %d  data is: %+v", id, data)
			}
		}(i)
	}
	select {}
}

func Test_AtomicPush(t *testing.T) {
	var a int64 = 1

	if atomic.CompareAndSwapInt64(&a, 0, 1) {
		defer atomic.StoreInt64(&a, 0)
		log.Println("hello world")
	}
	//log.Printf("swaped: %+v", ok)
}
