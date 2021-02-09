package render

import (
	"context"
	"fmt"
	
	"github.com/tfriedel6/canvas"
	
	"github.com/maxpower89/guio/pkg/element"
	"github.com/maxpower89/guio/pkg/shape"
	"github.com/maxpower89/guio/pkg/theme"
)

type Renderer interface {
	Render(ctx context.Context, canvas *canvas.Canvas, data []byte) error
	RegisterComponent(key string, handler shape.NewComponent)
	ThemeLoader() theme.Loader
}

type renderer struct {
	components  map[string]shape.NewComponent
	themeLoader theme.Loader
	lastState   []byte
	rootElement *element.Element
	count       int
}

func (r *renderer) ThemeLoader() theme.Loader {
	return r.themeLoader
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

func (r *renderer) Render(ctx context.Context, canvas *canvas.Canvas, data []byte) error {
	if string(data) == string(r.lastState) {
		if r.count > 0 {
			return nil
		}
		r.count++
	} else {
		var err error
		r.rootElement, err = element.FromXml(data, "shape")
		if err != nil {
			return err
		}
		r.lastState = data
		r.count = 0
	}
	
	fmt.Println("Rerender");
	rootComponent, _ := shape.NewShape(map[string]string{})
	
	r.setupComponent(r.rootElement, rootComponent)
	r.renderFunc(ctx, rootComponent, canvas)
	
	return nil
}

func (r *renderer) setupComponent(element *element.Element, component shape.Component) {
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
		postDraw, err := drawEvents.PreDraw(ctx, canvas, r.themeLoader)
		if err != nil {
			defer postDraw()
		}
	}
	component.Draw(ctx, r.renderFunc, canvas, r.themeLoader)
}
