package scheduler
import (
	"encoding/json"
	"proj3/png"
	"proj3/concurrent"
	"io"
	"fmt"
	"image"

)

type newTask struct {
	pngImg 	*png.ImageTask
}

func getWork(dec *json.Decoder, size string) *png.ImageTask{
	
	var request request
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

func (newTask *newTask) Run() {
	for i, effect := range newTask.pngImg.Request.Effects {
		bounds := newTask.pngImg.Out.Bounds()
		if i!=0 {
			newTask.pngImg.In = newTask.pngImg.Out
			newTask.pngImg.Out = image.NewRGBA64(bounds)
		}
		switch effect{
		case "S":
			newTask.pngImg.Sharpen(bounds.Min.Y, bounds.Max.Y)
		case "E":
			newTask.pngImg.Edge(bounds.Min.Y, bounds.Max.Y)
		case "B":
			newTask.pngImg.Blur(bounds.Min.Y, bounds.Max.Y)
		case "G":
			newTask.pngImg.Grayscale(bounds.Min.Y, bounds.Max.Y)
		}
	}
	
	//Saves the image to a new file
	OutPath := fmt.Sprint("../data/out/", newTask.pngImg.Request.Size,"_",newTask.pngImg.Request.OutPath)
	err := newTask.pngImg.Save(OutPath)

	//Checks to see if there were any errors when saving.
	if err != nil {
		panic(err)
	}
}

func RunParallel(config Config) {
	dec := read()
	var futures []concurrent.Future
	var executor concurrent.ExecutorService
	if config.Mode == "steal" {
		executor = concurrent.NewWorkStealingExecutor(config.ThreadCount, 10)
	} else if config.Mode == "balance" {
		executor = concurrent.NewWorkBalancingExecutor(config.ThreadCount, 10, 2)
	}

	for {
		task := getWork(dec, config.DataDirs) 
		if task== nil {
			executor.Shutdown()
			break
		} else {
			futures = append(futures, executor.Submit(&newTask{pngImg: task}))
		}
	}

	for _, future := range futures {
		future.Get()
	}

	return

}
