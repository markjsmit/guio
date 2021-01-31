package shape

import (
	"context"
	
	"github.com/mitchellh/mapstructure"
	"github.com/tfriedel6/canvas"
	
	"github.com/maxpower89/guio/pkg/theme"
)

func NewRect(attributes map[string]string) (Component, error) {
	rect := Rect{
		X:      0,
		Y:      0,
		Width:  1,
		Height: 1,
		Fill:   "",
		Stroke: "",
	}
	mapstructure.WeakDecode(attributes, &rect)
	return &rect, nil
}

type Rect struct {
	X, Y, Width, Height float64
	Fill, Stroke        string
}

func (r Rect) Draw(ctx context.Context, renderFunc RenderFunc, canvas *canvas.Canvas, themeLoader theme.Loader) error {
	canvas.BeginPath()
	canvas.Rect(r.X, r.Y, r.Width, r.Height)
	
	if r.Fill != "" {
		canvas.SetFillStyle(r.Fill)
		canvas.Fill()
	}
	
	if r.Stroke != "" {
		canvas.SetStrokeStyle(r.Stroke)
		canvas.Stroke()
	}
	canvas.ClosePath()
	return nil
}
