package concurrent

import (
	"math/rand"
	"sync"
	"time"
)

type SCtx struct { //stealing context
	capacity 	int
	queues 		[]DEQueue
	shutdown	bool //true if the service is to be shut down
}

// NewWorkStealingExecutor returns an ExecutorService that is implemented using the work-stealing algorithm.
// @param capacity - The number of goroutines in the pool
// @param threshold - The number of items that a goroutine in the pool can
// grab from the executor in one time period. For example, if threshold = 10
// this means that a goroutine can grab 10 items from the executor all at
// once to place into their local queue before grabbing more items. It's
// not required that you use this parameter in your implementation.
func NewWorkStealingExecutor(capacity, threshold int) ExecutorService {
	var queues []DEQueue
	for i:=0; i < capacity; i++ {
		queues = append(queues, NewUnBoundedDEQueue())
	}

	newCtx := &SCtx{capacity: capacity, queues: queues, shutdown: false}

	for i:=0; i < capacity; i++ {
		go newCtx.run(i)
	}

	return newCtx
}

func (SCtx *SCtx) run(threadID int) {
	
	task := SCtx.queues[threadID].PopBottom()

	for {
		for task != nil {
			future := task.(*shareFuture)
			callable, yes := future.task.(Callable)
			if yes {
				future.result = callable.Call()
			} else {
				runnable := future.task.(Runnable)
				runnable.Run() 
			}
			future.wg.Done()

			task = SCtx.queues[threadID].PopBottom()
		} 

		for task == nil {
			if SCtx.queues[threadID].Size() == 0 && SCtx.shutdown {
				return
			}
			
			rand.Seed(time.Now().UnixNano())
			victim := rand.Intn(SCtx.capacity) 
			if !SCtx.queues[victim].IsEmpty() {
				task = SCtx.queues[victim].PopTop()
			}
		}
	}

}


// submit pushes a task to a random thread
func (SCtx *SCtx) Submit(task interface{}) Future{
	if SCtx.shutdown {
		return nil
	}

	rand.Seed(time.Now().UnixNano())
	receiver := rand.Intn(SCtx.capacity)
	var wg sync.WaitGroup
	wg.Add(1)
	future := &shareFuture{task: task, wg: &wg, result: nil}
	SCtx.queues[receiver].PushBottom(future)

	return future
}


func (SCtx *SCtx) Shutdown() {
	SCtx.shutdown = true
}
