package scheduler
import (
	"encoding/JSON"
	"proj2/png"
	"fmt"
	"os"
	"image"
	"io"
)

func read() *json.Decoder{
	effectsPathFile := fmt.Sprintf("../data/effects.txt")
	effectsFile, _ := os.Open(effectsPathFile)
	reader := json.NewDecoder(effectsFile)
	return reader
}

type request struct {
	InPath 		string 
	OutPath		string 
	Effects 	[]string 
	size		string	
}

func process(request *request) {
	filePath := fmt.Sprint("../data/in/", request.size, "/", request.InPath)
	
	//Loads the png image and returns the image or an error
	pngImg, err := png.Load(filePath)
	if err != nil {
		panic(err)
	}

	for i, effect := range request.Effects {
		if i!=0 {
			pngImg.In = pngImg.Out
			pngImg.Out = image.NewRGBA64(pngImg.Out.Bounds())
		}
		switch effect{
		case "S":
			pngImg.Sharpen()
		case "E":
			pngImg.Edge()
		case "B":
			pngImg.Blur()
		}
	}
	
	//Saves the image to a new file
	OutPath := fmt.Sprint("../data/out/", request.InPath)
	err = pngImg.Save(OutPath)

	//Checks to see if there were any errors when saving.
	if err != nil {
		panic(err)
	}

}

func RunSequential(config Config) {
	dec := read()
	for true {
		var request request
		
		err := dec.Decode(&request)
		if err == io.EOF {
			break
		}

		request.size = config.DataDirs

		process(&request)
	}
}
