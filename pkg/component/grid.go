package component

import (
	"context"
	
	"github.com/mitchellh/mapstructure"
	"github.com/veandco/go-sdl2/sdl"
	
	"github.com/maxpower89/guio/pkg/event"
	"github.com/maxpower89/guio/pkg/style"
	"github.com/maxpower89/guio/pkg/theme"
)

type Grid struct {
	themeLoader theme.Loader
	children    []Component
}

func NewGrid(loader theme.Loader, attributes map[string]string) (Component, error) {
	grid := &Grid{
		themeLoader: loader,
		children:    []Component{},
	}
	mapstructure.WeakDecode(attributes, grid)
	return grid, nil
}

func (g *Grid) RegisterWindowHandler(handler event.Handler) {
	for _, child := range g.children {
		if reactive, ok := child.(Reactive); ok {
			reactive.RegisterWindowHandler(handler)
		}
	}
}

func (g *Grid) AddChild(component Component) {
	if _, ok := component.(*GridRow); ok {
		g.children = append(g.children, component)
	}
}

func (g *Grid) GetChildren() []Component {
	return g.children
}

func (g Grid) Update(ctx context.Context) ([]byte, error) {
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

func (b *Grid) Rect(ctx context.Context) sdl.Rect {
	return style.Rect(ctx, "0", "0", "100%", "100%")
}
