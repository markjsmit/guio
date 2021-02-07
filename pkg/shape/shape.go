package shape

import (
	"context"
	
	"github.com/tfriedel6/canvas"
	
	"github.com/maxpower89/guio/pkg/theme"
)

func NewShape(attributes map[string]string) (Component, error) {
	return &Shape{
		children: []Component{},
	}, nil
}

type Shape struct {
	children []Component
}

func (r *Shape) AddChild(shape Component) {
	r.children = append(r.children, shape)
}

func (r *Shape) Draw(ctx context.Context, renderFunc RenderFunc, canvas *canvas.Canvas, themeLoader theme.Loader) error {
	canvas.BeginPath()
	for _, child := range r.children {
		renderFunc(ctx, child, canvas)
	}
	canvas.ClosePath()
	return nil
}
