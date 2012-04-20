package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

var (
	nWorker      = 4
	nBufferedReq = 50
	workFn       = func() (r int) {
		r = rand.Int()
		fmt.Println("workFn", r)
		return
	}
)

type Request struct {
	fn func() int // The operation to perform
	r  chan int   // The channel to return the result
}

func requester(work chan<- Request) {
	r := make(chan int)
	for {
		work <- Request{workFn, r}
		result := <-r
		furtherProcess(result)
		// fake load
		time.Sleep(time.Duration(rand.Intn(nWorker)) * 2 * time.Second)
	}
}

func furtherProcess(r interface{}) {
	result := r.(int)
	fmt.Println("furtherProcess", result)
}

type Worker struct {
	requests chan Request // work to do (buffered channel)
	pending  int          // count of pending tasks
	index    int          // index in the heap
}

func (w *Worker) work(done chan *Worker) {
	for {
		req := <-w.requests
		req.r <- req.fn()
		done <- w
	}
}

type Pool []*Worker

func (p Pool) Less(i, j int) bool {
	return p[i].pending < p[j].pending
}

func (p Pool) Len() int {
	return len(p)
}

func (p *Pool) Push(i interface{}) {
	a := *p
	n := len(a)
	a = a[0 : n+1]
	a[n] = i.(*Worker)
	a[n].index = n
	*p = a
}

func (p *Pool) Pop() interface{} {
	a := *p
	n := len(a)
	item := a[n-1]
	*p = a[0 : n-1]
	return item
}

func (p *Pool) Swap(i, j int) {
	t := (*p)[i]
	(*p)[i] = (*p)[j]
	(*p)[j] = t
}

type Balancer struct {
	pool Pool
	done chan *Worker
}

func (b *Balancer) balance(work chan Request) {
	for {
		select {
		case req := <-work: // received a Request...
			b.dispatch(req) // ...so send it to a Worker
		case w := <-b.done: // a worker has finished ...
			b.completed(w) // ...so update its info
		}
	}
}

// Send Request to worker
func (b *Balancer) dispatch(req Request) {
	// Grab the least loaded worker...
	w := heap.Pop(&b.pool).(*Worker)
	// ...send it the task.
	w.requests <- req
	// One more in its work queue.
	w.pending++
	// Put it into its place on the heap.
	heap.Push(&b.pool, w)
}

// Job is complete; update heap
func (b *Balancer) completed(w *Worker) {
	// One fewer in the queue.
	w.pending--
	// Remove it from heap.                  
	heap.Remove(&b.pool, w.index)
	// Put it into its place on the heap.
	heap.Push(&b.pool, w)
}

func main() {
	// Init
	fmt.Println("Init...")

	// Worker Pool
	workerpool := make(Pool, nWorker)
	for i := 0; i < nWorker; i++ {
		workerpool[i] = new(Worker)
		workerpool[i].requests = make(chan Request, nBufferedReq/nWorker)
		workerpool[i].pending = 0
		workerpool[i].index = -1
	}
	fmt.Println("Worker Pool:", len(workerpool))

	// Balancer
	balancer := Balancer{workerpool, make(chan *Worker)}
	heap.Init(&balancer.pool)

	for i := 0; i < nWorker; i++ {
		go workerpool[i].work(balancer.done)
	}

	workPipe := make(chan Request, nBufferedReq)

	// Fake requester
	go requester(workPipe)

	// GO! GO! GO!
	balancer.balance(workPipe)

}
