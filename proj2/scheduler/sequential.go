package scheduler
import (
	"encoding/json"
	"strings"
	"proj2/png"
	"fmt"
	"os"
	"image"
	"io"
)


type request = png.Request

func read() *json.Decoder{
	effectsPathFile := fmt.Sprintf("../data/effects.txt")
	effectsFile, _ := os.Open(effectsPathFile)
	reader := json.NewDecoder(effectsFile)
	return reader
}



func process(request *request) {
	filePath := fmt.Sprint("../data/in/", request.Size, "/", request.InPath)
	
	//Loads the png image and returns the image or an error
	pngImg, err := png.Load(filePath, request)
	if err != nil {
		panic(err)
	}

	for i, effect := range request.Effects {
		bounds := pngImg.Out.Bounds()
		if i!=0 {
			pngImg.In = pngImg.Out
			pngImg.Out = image.NewRGBA64(bounds)
		}
		switch effect{
		case "S":
			pngImg.Sharpen(bounds.Min.Y, bounds.Max.Y)
		case "E":
			pngImg.Edge(bounds.Min.Y, bounds.Max.Y)
		case "B":
			pngImg.Blur(bounds.Min.Y, bounds.Max.Y)
		case "G":
			pngImg.Grayscale(bounds.Min.Y, bounds.Max.Y)
		}
	}
	
	//Saves the image to a new file
	OutPath := fmt.Sprint("../data/out/", request.Size,"_",request.OutPath)
	err = pngImg.Save(OutPath)

	//Checks to see if there were any errors when saving.
	if err != nil {
		panic(err)
	}

}

func RunSequential(config Config) {
	
	sizes := strings.Split(config.DataDirs, "+")

	for _, size := range sizes {
		dec := read()
		for true {
			var request request
			err := dec.Decode(&request)
			if err == io.EOF {
				break
			}
			request.Size = size
			process(&request)
		}
	}
}
