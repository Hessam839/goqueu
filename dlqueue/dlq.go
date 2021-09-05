package dlqueue

import (
	"container/list"
	"sync"
)

type dlqueue struct {
	store *list.List
	wait  *sync.Cond
}

func NewDlQueue() *dlqueue {
	mux := &sync.Mutex{}
	return &dlqueue{
		store: list.New(),
		wait:  sync.NewCond(mux),
	}
}

func (q *dlqueue) Put(data interface{}) {
	q.store.PushBack(data)
	q.wait.Signal()
}

func (q *dlqueue) Get() interface{} {
	elem := q.store.Front()
	q.store.Remove(elem)
	q.wait.Signal()

	return elem.Value
}

func (q *dlqueue) TryGet() interface{} {

	q.wait.L.Lock()
	defer q.wait.L.Unlock()
	var elem *list.Element

	if q.store.Len() == 0 {
		q.wait.Wait()
	}
	elem = q.store.Front()
	//if elem == nil {
	//	return nil
	//}
	q.store.Remove(elem)
	q.wait.Signal()
	return elem.Value
}
