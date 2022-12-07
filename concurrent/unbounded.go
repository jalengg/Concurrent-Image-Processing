package concurrent

import (
	
	"sync"
)
/**** YOU CANNOT MODIFY ANY OF THE FOLLOWING INTERFACES/TYPES ********/
type Task interface{}

type DEQueue interface {
	PushBottom(task Task)
	IsEmpty() bool //returns whether the queue is empty
	PopTop() Task
	PopBottom() Task
	Size()	int
}

/******** DO NOT MODIFY ANY OF THE ABOVE INTERFACES/TYPES *********************/

type node struct {
	task 	Task
	prev 	*node
	next 	*node
	
}

type queue struct {
	top 	*node
	bottom 	*node
	size 	int
	mu 		*sync.Mutex
}


// NewUnBoundedDEQueue returns an empty UnBoundedDEQueue
func NewUnBoundedDEQueue() DEQueue {
	var lock sync.Mutex
	return &queue{top: nil, bottom: nil, size: 0, mu: &lock}
}

func newNode(task Task) *node {
	
	return &node{task: task, prev: nil, next: nil}
}

func (queue *queue) PushBottom(task Task) {
	queue.mu.Lock() 
	defer queue.mu.Unlock()

	new := newNode(task)

	if queue.bottom == nil {
		queue.bottom = new
		queue.top = new
	} else {
		new.prev = queue.bottom
		queue.bottom.next = new
		queue.bottom = new
	}

	queue.size++
}

func (queue *queue) IsEmpty() bool{
	return (queue.bottom == nil || queue.top == nil || queue.Size() == 0)
}

func (queue *queue) PopTop() Task{

	

	if (queue.IsEmpty()) {
		return nil  //check again to make sure no one beat us
	} else {
		queue.mu.Lock() 
		defer queue.mu.Unlock()

		queue.size--
		temp := queue.top
		queue.top = queue.top.next
		
		if (queue.top==nil) {
			queue.bottom = nil
		} else {
			queue.top.prev = nil
		}

		return temp.task
	}
}

func (queue *queue) PopBottom() Task {
	if (queue.IsEmpty()) {
		return nil
	} else {
		queue.mu.Lock() 
		defer queue.mu.Unlock()

		queue.size--
		temp := queue.bottom
		queue.bottom = queue.bottom.prev
		
		if (queue.bottom==nil) {
			queue.top = nil
		} else {
			queue.bottom.next = nil
		}

		return temp.task
	}
}

func (queue *queue) Size() int{
	return queue.size
}
