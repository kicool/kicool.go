// r=a(1-sinθ)
//Idea form http://www.oschina.net/code/snippet_104514_18581 
package main

import (
	"bufio"
	"code.google.com/p/draw2d/draw2d"
	"fmt"
	"image"
	"image/png"
	"log"
	"math"
	"os"
	"image/color"
)

const (
	width, height = 512, 512
)

var (
	lastTime int64
	folder   = "./result/"
)

func initGc(w, h int) (image.Image, draw2d.GraphicContext) {
	i := image.NewRGBA(image.Rect(0, 0, w, h))
	gc := draw2d.NewGraphicContext(i)

	gc.SetStrokeColor(image.Black)
	gc.SetFillColor(image.White)
	// fill the background 
	//gc.Clear()

	return i, gc
}

func saveToPngFile(testName string, m image.Image) {
	filePath := folder + testName + ".png"
	f, err := os.Create(filePath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	b := bufio.NewWriter(f)
	err = png.Encode(b, m)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Wrote %s OK.\n", filePath)
}

func Draw() {
	i, gc := initGc(width, height)

	for i := 0; i < 90; i++ {
		for j := 0; j < 90; j++ {
			r := math.Pi / 45 * float64(i) * (1 - math.Sin(math.Pi / 45 * float64(j))) * 19
			x := r * math.Cos(math.Pi / 45 * float64(j)) * math.Sin(math.Pi / 45 * float64(i)) + width / 2
			y := -r * math.Sin(math.Pi / 45 * float64(j)) + height / 4

			gc.SetStrokeColor(color.RGBA{0xFF, 0x33, 0x33, 0xff})
			gc.MoveTo(x, y)
			gc.LineTo(x+1, y+1)
			gc.Stroke()
		}
	}

	saveToPngFile("Draw", i)
}

func main() {
	Draw()
}
