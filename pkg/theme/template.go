package theme

import (
	"fmt"
	"strconv"
	"text/template"
	
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/backend/softwarebackend"
)

var canvasSimulator = canvas.New(softwarebackend.New(1000, 1000))

func TemplateFuncs(loader Loader) template.FuncMap {
	return template.FuncMap{
		"add":      sum,
		"textPosY": sum,
		"cursorX":  cursorX(loader),
	}
}

func sum(args ...interface{}) float64 {
	var cur float64 = 0
	for _, arg := range args {
		val, _ := strconv.ParseFloat(fmt.Sprint(arg), 64)
		cur += val
	}
	return cur
}

func cursorX(loader Loader) func(x float32, cursorPos int, fontSize float64, fontFamily string, text string) float32 {
	return func(x float32, cursorPos int, fontSize float64, fontFamily string, text string) float32 {
		
		usedText := text[0:cursorPos]
		canvasSimulator.BeginPath()
		canvasSimulator.SetFont(loader.Path("font", fontFamily), fontSize)
		offset := canvasSimulator.MeasureText(usedText).Width
		canvasSimulator.ClosePath()
		return x + float32(offset) +1
	}
}
