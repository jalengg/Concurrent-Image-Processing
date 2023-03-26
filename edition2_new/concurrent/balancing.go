package concurrent

import (
	"math/rand"
	"sync"
	"time"
)

type WBCtx struct { //working balance context
	capacity 	int
	queues 		[]DEQueue
	threshold 	int 
	shutdown	bool //true if the service is to be shut down
}

type shareFuture struct {
	task	interface{}
	wg 		*sync.WaitGroup
	result 	interface{}
}

// NewWorkBalancingExecutor returns an ExecutorService that is implemented using the work-balancing algorithm.
// @param capacity - The number of goroutines in the pool
// @param threshold - The number of items that a goroutine in the pool can
// grab from the executor in one time period. For example, if threshold = 10
// this means that a goroutine can grab 10 items from the executor all at
// once to place into their local queue before grabbing more items. It's
// not required that you use this parameter in your implementation.
// @param thresholdBalance - The threshold used to know when to perform
// balancing. Remember, if two local queues are to be balanced the
// difference in the sizes of the queues must be greater than or equal to
// thresholdBalance. You must use this parameter in your implementation.
func NewWorkBalancingExecutor(capacity, thresholdQueue, thresholdBalance int) ExecutorService {
	var queues []DEQueue
	for i:=0; i < capacity; i++ {
		queues = append(queues, NewUnBoundedDEQueue())
	}

	newCtx := &WBCtx{capacity: capacity, queues: queues, threshold: thresholdBalance, shutdown: false}

	for i:=0; i < capacity; i++ {
		go newCtx.run(i)
	}

	return newCtx
}

func (WBCtx *WBCtx) run(threadID int) {
	for {
		task := WBCtx.queues[threadID].PopBottom()
		if task != nil {
			future := task.(*shareFuture)
			callable, yes := future.task.(Callable)
			if yes {
				future.result = callable.Call()
			} else {
				runnable := future.task.(Runnable)
				runnable.Run() 
			}
			
			future.wg.Done()
		} 

		size := WBCtx.queues[threadID].Size()

		// the chance to balance increases as the queue size decreases
		if (rand.Intn(size+1) == size) {
			rand.Seed(time.Now().UnixNano())
			victim := rand.Intn(WBCtx.capacity)

			var min, max int
			if victim <= threadID {
				min = victim
				max = threadID
			} else {
				min = threadID
				max = victim
			}

			//balancing algorithm
			WBCtx.balance(min, max)
		}

		if WBCtx.queues[threadID].Size() == 0 && WBCtx.shutdown {
			return
		}
	}
}

func (WBCtx *WBCtx) balance(q0 int, q1 int) {
	var qMin, qMax int
	if WBCtx.queues[q0].Size() < WBCtx.queues[q1].Size(){
		qMin = q0
		qMax = q1
	} else {
		qMin = q1
		qMax = q0
	}
	diff := WBCtx.queues[qMax].Size() - WBCtx.queues[qMin].Size()

	if (diff > WBCtx.threshold) {
		for (WBCtx.queues[qMax].Size() > WBCtx.queues[qMax].Size()) {
			WBCtx.queues[qMax].PushBottom(WBCtx.queues[qMax].PopTop())
		}
	}
}


// submit pushes a task to a random thread
func (WBCtx *WBCtx) Submit(task interface{}) Future{
	if WBCtx.shutdown {
		return nil
	}

	rand.Seed(time.Now().UnixNano())
	receiver := rand.Intn(WBCtx.capacity)
	var wg sync.WaitGroup
	wg.Add(1)
	future := &shareFuture{task: task, wg: &wg, result: nil}
	WBCtx.queues[receiver].PushBottom(future)

	return future
}


func (WBCtx *WBCtx) Shutdown() {
	WBCtx.shutdown = true
}


func (future *shareFuture) Get() interface{} {
	future.wg.Wait()
	return future.result
}
