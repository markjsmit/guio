package shape

import (
	"context"
	
	"github.com/tfriedel6/canvas"
	
	"github.com/maxpower89/guio/pkg/theme"
)

type NewComponent func(attributes map[string]string) (Component, error)

type Component interface {
	Draw(ctx context.Context, renderFunc RenderFunc, canvas *canvas.Canvas, themeLoader theme.Loader) error
}

type PostDrawFunc func()
type Container interface {
	AddChild(shape Component)
}

type RenderFunc func(ctx context.Context, component Component, canvas *canvas.Canvas)

type DrawEvents interface {
	PreDraw(ctx context.Context, canvas *canvas.Canvas, themeLoader theme.Loader) (PostDrawFunc, error)
}

type TouchOutput struct {
	Touched bool
}

type Touchable interface {
	Touched(canvas *canvas.Canvas, x float64, y float64) TouchOutput
}
