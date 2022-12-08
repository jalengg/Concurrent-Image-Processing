package mapReduce

import (
	"math/rand"
	"math"
	"fmt"
	"time"

)


type SSSP struct {
	graph 		map[int][]int  //graph represented by key = nodeID and value = list ofnodes reachable from nodeID
	distances 	map[int]float64 // value (distance) to  nodeID (key) from node 0 
}

//generate a random single source shortest path probelm 
func NewGraph(size int) *SSSP { //size dictates the number of nodes in the graph
	rand.Seed(time.Now().UnixNano())
	graphSize := size
	newGraph := make(map[int][]int)
	for i:=0; i<graphSize; i++ {
		listSize := rand.Intn(graphSize)
		var list []int
		for j:=0; j<listSize; j++ {
			list = append(list, rand.Intn(graphSize)) //there will be duplicate nodes in thelist
		}
		newGraph[i] = list
	}

	newDistances := make(map[int]float64)
	newDistances[0] = 0.0
	for i:=1; i<graphSize; i++ {
		newDistances[i] = math.MaxFloat64
    }

	return &SSSP{graph: newGraph, distances: newDistances}
}

//mapper structs and associated functions
type mapperSupplier struct {
	problem *SSSP
}

type Mapper struct {
	//everything you need to map
	inputNode 		int
	distance 		float64 //distance to origin
	adjacencyList	[]int //list of nodes reachable from inputNode
}

func (problem *SSSP) Mapper() *mapperSupplier{
	return &mapperSupplier{problem: problem}
}

func (supplier *mapperSupplier) Get(inputNode int) *Mapper{
	return &Mapper{inputNode: inputNode, distance: supplier.problem.distances[inputNode], adjacencyList: supplier.problem.graph[inputNode]}
}


func (mapper *Mapper)Call() interface{}{ //returns all the nodes reachable from input node and the possible distance from the source node
	output := make(map[int]float64)
	distance := mapper.distance
	output[mapper.inputNode] = distance
	for _, i := range mapper.adjacencyList {
		output[i] = distance + 1
	}
	return output
}

//reducer and associated types 
type reducerSupplier struct {
	problem *SSSP

}

type Reducer struct {
	//everything you need to reduce
	targetNode 		int
	distances		[]float64 //possible distances to the target from the source node
	curDist			float64 //current distance in the graph
}

func (problem *SSSP) Reducer() *reducerSupplier{
	return &reducerSupplier{problem: problem}
}

func (supplier *reducerSupplier) Get(targetNode int, distances []float64) *Reducer{
	//everything you need to reduce
	return &Reducer{targetNode: targetNode, distances: distances, curDist: supplier.problem.distances[targetNode]}
}


func (reducer *Reducer) Call() interface{} { //returns the minimum distance to targetNode
	min := reducer.curDist
	for _, d := range reducer.distances {
		if d < min {
			min = d
		}
	}
	return min
}

func (problem *SSSP) listOfFiniteDistanceNodes() []int {
	distances := problem.distances
	var output []int
	for key, value := range distances {
		if value < math.MaxFloat64 {
			output = append(output, key)
		}
	}
	return output
}


func (problem *SSSP) updateAndCompare(new map[int]float64, epsilon float64) bool{ 
	//returns true if new distances is within epsilon from oldDistances
	var diff float64
	for key, value := range new {
		old := problem.distances[key]
		if old > value {
			problem.distances[key] = value
			diff = diff + (old-value)
		}
	}
	return diff < epsilon
}

func (problem *SSSP) Display() {
	fmt.Println("Adjacency Matrix:")
	for key, value := range problem.graph {
		fmt.Println(key, ": ", value)
	}

	fmt.Println("Distance to node N from the source")
	for key, value := range problem.distances {
		fmt.Println("Node", key, ": ", value)
	}
}

func RunMap(problem *SSSP, capacity int, mode string) *SSSP{

	
	done := false;
	for !done {
		mapReducer := NewMapReduce(capacity, mode)
		mapReducer.SetMapperSupplier(problem.Mapper())
		mapReducer.SetReducerSupplier(problem.Reducer())
		
		mapReducer.SetInput(problem.listOfFiniteDistanceNodes())
		
		done = problem.updateAndCompare(mapReducer.Call(), 3)
	}
	return problem
}