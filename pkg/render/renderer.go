package render

import (
	"context"
	
	"github.com/tfriedel6/canvas"
	
	"github.com/maxpower89/guio/pkg/element"
	"github.com/maxpower89/guio/pkg/shape"
	"github.com/maxpower89/guio/pkg/theme"
)

type Renderer interface {
	Render(ctx context.Context, canvas *canvas.Canvas, data []byte) error
	RegisterComponent(key string, handler shape.NewComponent)
}

type renderer struct {
	components    map[string]shape.NewComponent
	rootComponent shape.Component
	themeLoader   theme.Loader
}

func NewRenderer(loader theme.Loader) Renderer {
	renderer := &renderer{
		themeLoader: loader,
		components:  map[string]shape.NewComponent{},
	}
	return renderer
}

func (r *renderer) RegisterComponent(key string, handler shape.NewComponent) {
	r.components[key] = handler
}

func (r renderer) Render(ctx context.Context, canvas *canvas.Canvas, data []byte) error {
	rootElement, err := element.FromXml(data, "shape")
	if err != nil {
		return err
	}
	
	rootComponent, _ := shape.NewShape(map[string]string{})
	r.setupComponent(rootElement, rootComponent)
	
	r.renderFunc(ctx, rootComponent, canvas)
	
	return nil
}

func (r *renderer) setupComponent(element element.Element, component shape.Component) {
	if container, ok := component.(shape.Container); ok {
		for _, child := range element.Children {
			if constructor, ok := r.components[child.Tag]; ok {
				childComponent, _ := constructor(child.Attr)
				r.setupComponent(child, childComponent)
				container.AddChild(childComponent)
			}
		}
	}
}

func (r *renderer) renderFunc(ctx context.Context, component shape.Component, canvas *canvas.Canvas) {
	if drawEvents, ok := component.(shape.DrawEvents); ok {
		postDraw, err := drawEvents.PreDraw(ctx, canvas,r.themeLoader)
		if err != nil {
			defer postDraw()
		}
	}
	component.Draw(ctx, r.renderFunc, canvas,r.themeLoader)
}
