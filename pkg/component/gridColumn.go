package component

import (
	"context"
	
	"github.com/mitchellh/mapstructure"
	"github.com/veandco/go-sdl2/sdl"
	
	"github.com/maxpower89/guio/pkg/event"
	"github.com/maxpower89/guio/pkg/style"
	"github.com/maxpower89/guio/pkg/theme"
)

type GridColumn struct {
	themeLoader theme.Loader
	children    []Component
}

func NewGridColumn(loader theme.Loader, attributes map[string]string) (Component, error) {
	column := &GridColumn{
		
		themeLoader: loader,
		children:    []Component{},
	}
	mapstructure.WeakDecode(attributes, column)
	return column, nil
}

func (g *GridColumn) RegisterWindowHandler(handler event.Handler) {
	for _, child := range g.children {
		if reactive, ok := child.(Reactive); ok {
			reactive.RegisterWindowHandler(handler)
		}
	}
}

func (g *GridColumn) AddChild(component Component) {
	g.children = append(g.children, component)
}

func (g *GridColumn) GetChildren() []Component {
	return g.children
}

func (g GridColumn) Update(ctx context.Context) ([]byte, error) {
	children := ""
	
	for _, child := range g.children {
		c, _ := child.Update(ctx)
		children += string(c)
	}
	
	return g.themeLoader.Load(ctx, "container", map[string]string{
		"width":    style.Width(ctx, "100%"),
		"height":   style.Height(ctx, "100%"),
		"children": children,
	})
}

func (b *GridColumn) Rect(ctx context.Context) sdl.Rect {
	return style.Rect(ctx, "0", "0", "100%", "100%")
}

