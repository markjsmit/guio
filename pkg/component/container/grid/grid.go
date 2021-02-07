package grid

import (
	"context"
	"fmt"
	
	"github.com/mitchellh/mapstructure"
	"github.com/veandco/go-sdl2/sdl"
	
	"github.com/maxpower89/guio/pkg/component"
	"github.com/maxpower89/guio/pkg/style"
	"github.com/maxpower89/guio/pkg/theme"
)

const Elem = "grid"

type Grid struct {
	style           *style.Style
	Style           string
	id              component.Identifier
	Id              string
	Class           string
	themeLoader     theme.Loader
	children        component.ComponentGroup
	componentLoader theme.ComponentLoader
}

func NewGrid(loader theme.Loader, parent component.Component, attributes map[string]string) (component.Component, error) {
	grid := &Grid{
		themeLoader: loader,
		children:    component.NewComponentGroup([]component.Component{}),
	}
	grid.componentLoader = theme.NewComponentLoader(loader, "container")
	mapstructure.WeakDecode(attributes, grid)
	grid.id = component.NewIdentifier(grid.Class, grid.Id, Elem,component.StateDefault)
	grid.style = style.FromCss([]byte(grid.Style), loader)
	return grid, nil
}
func (r *Grid) Identify() component.Identifier {
	return r.id
}

func (w *Grid) Children() component.ComponentGroup {
	return w.children
}

func (g Grid) Update(ctx context.Context) ([]byte, error) {
	absoluteStyle := g.style.Build(ctx)
	children := g.updateChildren(absoluteStyle.InnerCtx)
	return g.componentLoader.Load(absoluteStyle.Ctx, Tpl{
		Style:    absoluteStyle,
		Children: children,
	})
}

func (g Grid) updateChildren(ctx context.Context) string {
	children := ""
	totalOffset := float32(0)
	g.Children().Each(func(child component.Component) {
		childCtx := style.BoxContext(ctx, "0", fmt.Sprint(totalOffset), "100%", "100%")
		c, _ := child.Update(childCtx)
		rect := child.Rect(childCtx)
		children += string(c)
		totalOffset += rect.H
	})
	return children
}

func (b *Grid) Rect(ctx context.Context) sdl.FRect {
	absoluteStyle := b.style.Build(ctx)
	return style.Rect(absoluteStyle.OuterCtx, "0", "0", "100%", "100%")
}

func (g *Grid) Styler() *style.Style {
	return g.style
}
