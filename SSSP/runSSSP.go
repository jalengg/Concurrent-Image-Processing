package main

import (
	"fmt"
	"os"
	"proj3/mapReduce"
	"strconv"
	"time"
)


const usage = "Usage: runSSSP.go graph_size mode [number of threads]\n" +
	"graph_size = The number of nodes in a graph.\n" +
	"mode     = (balance) run the balancing mode, (steal) run the stealing executorservice version\n" +
	"number of threads = Runs the parallel version of the program with the specified number of threads.\n"

func main() {

	if len(os.Args) < 3 {
		fmt.Println(usage)
		return
	}
	
	size, _ := strconv.Atoi(os.Args[1])
	mode := os.Args[2]
	threads, _ := strconv.Atoi(os.Args[3])

	problem := mapReduce.NewGraph(size)

	start := time.Now()

	mapReduce.RunMap(problem, threads, mode)

	end := time.Since(start).Seconds()
	fmt.Printf("%.2f\n", end)

}