// Copyright 2010 The draw2d Authors. All rights reserved.
// created: 21/11/2010 by Laurent Le Goff

package main

import (
	"bufio"
	"code.google.com/p/draw2d/draw2d"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

const (
	w, h = 512, 512
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

func saveToPngFile(TestName string, m image.Image) {
	filePath := folder + TestName + ".png"
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

func DrawLine() {
	i, gc := initGc(w, h)
	gc.MoveTo(10, 10)
	gc.LineTo(100, 10)
	gc.FillStroke()
	saveToPngFile("DrawLine", i)
}

func DrawArc() {
	i, gc := initGc(w, h)
	// draw an arc
	xc, yc := 128.0, 128.0
	radiusX, radiusY := 100.0, 100.0
	startAngle := 45.0 * (math.Pi / 180.0) /* angles are specified */
	angle := 135 * (math.Pi / 180.0)       /* in radians           */
	gc.SetLineWidth(10)
	gc.SetLineCap(draw2d.ButtCap)
	gc.SetStrokeColor(image.Black)
	gc.ArcTo(xc, yc, radiusX, radiusY, startAngle, angle)
	gc.Stroke()
	// fill a circle
	gc.SetStrokeColor(color.NRGBA{255, 0x33, 0x33, 0x80})
	gc.SetFillColor(color.NRGBA{255, 0x33, 0x33, 0x80})
	gc.SetLineWidth(6)

	gc.MoveTo(xc, yc)
	gc.LineTo(xc+math.Cos(startAngle)*radiusX, yc+math.Sin(startAngle)*radiusY)
	gc.MoveTo(xc, yc)
	gc.LineTo(xc-radiusX, yc)
	gc.Stroke()

	gc.ArcTo(xc, yc, 10.0, 10.0, 0, 2*math.Pi)
	gc.Fill()
	saveToPngFile("TestDrawArc", i)
}

func main() {
	DrawLine()
}
