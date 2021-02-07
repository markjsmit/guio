package shape

import (
	"context"
	
	"github.com/mitchellh/mapstructure"
	"github.com/tfriedel6/canvas"
	
	"github.com/maxpower89/guio/pkg/theme"
)

func NewLine(attributes map[string]string) (Component, error) {
	line := Line{
		X1:     0,
		Y1:     0,
		X2:     1,
		Y2:     1,
		Stroke: "",
	}
	mapstructure.WeakDecode(attributes, &line)
	return &line, nil
}

type Line struct {
	X1, Y1, X2, Y2 float64
	Stroke         string
}

func (r Line) Draw(ctx context.Context, renderFunc RenderFunc, canvas *canvas.Canvas, themeLoader theme.Loader) error {
	canvas.BeginPath()
	canvas.SetStrokeStyle(r.Stroke)
	canvas.MoveTo(r.X1, r.Y1)
	canvas.LineTo(r.X2, r.Y2)
	canvas.Stroke()
	canvas.ClosePath()
	return nil
}
