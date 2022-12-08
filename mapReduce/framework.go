package mapReduce

import(
	"proj3/concurrent"
	//"fmt"
)

type MapReduce struct {
	mapperSupplier	*mapperSupplier
	reducerSupplier	*reducerSupplier
	input 		[]int
	executor	concurrent.ExecutorService
}



func NewMapReduce(capacity int, mode string) *MapReduce{
	var executor concurrent.ExecutorService

	if mode == "steal" {
		executor = concurrent.NewWorkStealingExecutor(capacity, 10)
	} else if mode == "balance" {
		executor = concurrent.NewWorkBalancingExecutor(capacity, 10, 2)
	}
	
	return &MapReduce{executor: executor}
}

func (mapReduce *MapReduce) SetMapperSupplier(mapperSupplier *mapperSupplier) {
	mapReduce.mapperSupplier = mapperSupplier
}

func (mapReduce *MapReduce) SetReducerSupplier(reducerSupplier *reducerSupplier) {
	mapReduce.reducerSupplier = reducerSupplier
}

func (mapReduce *MapReduce) SetInput(inputNodes []int) {
	mapReduce.input = inputNodes
}


func (mapReduce *MapReduce) Call() map[int]float64{
	mapperFutures := make(map[int]concurrent.Future)
	for _, node := range mapReduce.input {   //input is all the correctly reachable nodes distance[node] < Max
		mapper := mapReduce.mapperSupplier.Get(node) 
	
		mapperFutures[node] = mapReduce.executor.Submit(mapper)
	}

	mapResults := make(map[int][]float64)
	for _, future := range mapperFutures {
		mapResult := future.Get().(map[int]float64)
		for targetNode, dist := range mapResult {
			mapResults[targetNode] = append(mapResults[targetNode], dist)
		}
	}

	reducerFutures := make(map[int]concurrent.Future)
	for targetNode, distances := range mapResults{
		reducer := mapReduce.reducerSupplier.Get(targetNode, distances)
		reducerFutures[targetNode] = mapReduce.executor.Submit(reducer)
	}

	result := make(map[int]float64)
	for targetNode, future := range reducerFutures {
		reduceResult := future.Get().(float64)
		result[targetNode] = reduceResult
		mapReduce.executor.Shutdown()
	}
	return result
}


