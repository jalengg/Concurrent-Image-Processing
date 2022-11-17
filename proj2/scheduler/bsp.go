package scheduler
import (
	"encoding/json"
	"proj2/png"
	"io"
	"fmt"
	"image"
	"strings"
	"sync"
)
type bspWorkerContext struct {
	decoder 	*json.Decoder
	mu 			*sync.Mutex
	syncCond	*sync.Cond
	startCond	*sync.Cond
	done 		bool
	start 		bool
	thread 		int
	sizes		[]string
	sizeIndex	int
	sizeLen		int
	currTask 	*png.ImageTask
	effectIndex	int
	effectLen	int
	threadCounter int
}

func NewBSPContext(config Config) *bspWorkerContext {
	//Initialize the context
	//get the reader and lock going

	dec := read()
	var mutex sync.Mutex
	startCond := sync.NewCond(&mutex)
	syncCond := sync.NewCond(&mutex)
	sizes := strings.Split(config.DataDirs, "+")
	sizeLen := len(sizes)
	ctx := &bspWorkerContext{decoder: dec, mu: &mutex, syncCond: syncCond, startCond: startCond, start: false,
		done: false, thread: config.ThreadCount, sizes: sizes, sizeIndex: 0, sizeLen:sizeLen, threadCounter: 0}
	
	getWork(ctx)
	
	return ctx
}

func RunBSPWorker(id int, ctx *bspWorkerContext) {
	for {
		// check done
		if ctx.done {
			return 
		}
		

		//superstep
		effect := ctx.currTask.Request.Effects[ctx.effectIndex]
		
		bounds := ctx.currTask.Out.Bounds()

		yIncrement :=  bounds.Max.Y / ctx.thread 
	
		ctx.currTask.BspMiniWorker(id*yIncrement, (id+1)*yIncrement, effect)

		///// global synchronization aka barrier 
		barrier(ctx)
	}
}


func barrier(ctx *bspWorkerContext) {
	ctx.mu.Lock()

	ctx.threadCounter++

	if ctx.threadCounter == ctx.thread {
		ctx.effectIndex++
		if ctx.effectIndex == ctx.effectLen {
			OutPath := fmt.Sprint("../data/out/", ctx.currTask.Request.Size,"_",ctx.currTask.Request.OutPath)
			err := ctx.currTask.Save(OutPath)
			//Checks to see if there were any errors when saving.
			if err != nil {
				panic(err)
			}
			getWork(ctx)
		} else {
			bounds := ctx.currTask.Out.Bounds()
			ctx.currTask.In = ctx.currTask.Out
			ctx.currTask.Out = image.NewRGBA64(bounds)
		}
		ctx.threadCounter = 0
		ctx.syncCond.Broadcast()
	} else {
		ctx.syncCond.Wait()
	}
	ctx.mu.Unlock()
}

func getWork(ctx *bspWorkerContext) {
	
	if ctx.done {
		return
	} 
	
	var request request
	err := ctx.decoder.Decode(&request)

	if err == io.EOF {
		//
		ctx.sizeIndex++
		if ctx.sizeIndex == ctx.sizeLen {
			ctx.done = true
			return 
		} else {
			ctx.decoder = read()
			ctx.decoder.Decode(&request)
			
		}
	}
	request.Size = ctx.sizes[ctx.sizeIndex]

	filePath := fmt.Sprint("../data/in/", request.Size, "/", request.InPath)
	//Loads the png image and returns the image or an error
	pngImg, err := png.Load(filePath, &request)
	if err != nil {
		panic(err)
	}

	ctx.currTask = pngImg
	ctx.effectIndex = 0
	ctx.effectLen = len(request.Effects)
	
	return 
}
