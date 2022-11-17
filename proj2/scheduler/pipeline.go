package scheduler
import (
	"encoding/json"
	"proj2/png"
	"io"
	"fmt"
	"image"
	"strings"
)

func getTask(dec *json.Decoder, size string) *png.ImageTask{
	var request png.Request
	err := dec.Decode(&request)
	if err == io.EOF {
		return nil
	}
	request.Size = size

	filePath := fmt.Sprint("../data/in/", request.Size, "/", request.InPath)
	
	//Loads the png image and returns the image or an error
	pngImg, err := png.Load(filePath, &request)
	if err != nil {
		panic(err)
	}
	return pngImg
}


func getTaskStream(taskStream chan *png.ImageTask, size string) {

	dec := read()
	for { 
		filter := getTask(dec, size)
		if filter == nil {
			return
		} else {
			taskStream <- filter
		}
	}
}

func taskBuffer(taskStream chan *png.ImageTask, sizes []string) {
	defer close(taskStream)
	for _, size := range sizes {
		getTaskStream(taskStream, size)
	}
	return
}

func generator(sizes []string) <-chan *png.ImageTask {
	
	taskStream := make(chan *png.ImageTask)
	go taskBuffer(taskStream, sizes)
	return taskStream
}

func worker(numThreads int, taskStream <-chan *png.ImageTask, resultStream chan<- *png.ImageTask) {
	defer close(resultStream)
	for pngImg := range taskStream {
		if pngImg == nil {
			continue
		}
		for i, effect := range pngImg.Request.Effects {
			bounds := pngImg.Out.Bounds()
			if i!=0 {
				pngImg.In = pngImg.Out
				pngImg.Out = image.NewRGBA64(bounds)
			}
	
			done := make(chan bool, numThreads)
	
			yIncrement :=  bounds.Max.Y / numThreads
			for j:=0; j < numThreads; j++ {
				go pngImg.MiniWorker(j*yIncrement, (j+1)*yIncrement, done, effect)
			}
	
			for k:=0; k < numThreads; k++ {
				<-done
			}
		}
		resultStream <- pngImg
	}
}



func fanOut(numThreads int, taskStream <-chan *png.ImageTask) <-chan *png.ImageTask{
	//fan out
	resultStream := make(chan *png.ImageTask)
	go worker(numThreads, taskStream, resultStream)
	return resultStream
}

func workerListener(c <-chan *png.ImageTask, done chan<- bool, aggregated chan<- *png.ImageTask) {
	for i := range c {
		aggregated <- i
	}
	done <- true
}


func aggregator(channels []<-chan *png.ImageTask, numThreads int) {
	//fan in 

	aggregated := make(chan *png.ImageTask)
	done := make(chan bool, numThreads)
	
	for _, c := range channels {
		go workerListener(c, done, aggregated)
	}

	// Wait for all the reads to complete
	go func() {
		for k:=0; k < numThreads; k++ {
			<-done
		}	
		close(aggregated)
	}()

	//Saves the image to a new file
	for result := range aggregated {
		OutPath := fmt.Sprint("../data/out/", result.Request.Size,"_",result.Request.OutPath)
		err := result.Save(OutPath)
	
		//Checks to see if there were any errors when saving.
		if err != nil {
			panic(err)
		}
	}
}

func RunPipeline(config Config) {
	sizes := strings.Split(config.DataDirs, "+")

	taskStream := generator(sizes)
	
	numThreads := config.ThreadCount
	taskResults := make([]<-chan *png.ImageTask, numThreads)

	for i := 0; i < numThreads; i++ {
		taskResults[i] = fanOut(numThreads, taskStream)
	}

	aggregator(taskResults, numThreads)
}
