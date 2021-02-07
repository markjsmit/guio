package grid

import (
	"context"
	"fmt"
	"strconv"
	
	"github.com/mitchellh/mapstructure"
	"github.com/veandco/go-sdl2/sdl"
	
	"github.com/maxpower89/guio/pkg/component"
	"github.com/maxpower89/guio/pkg/style"
	"github.com/maxpower89/guio/pkg/theme"
)
const RowElem = "grid-row"


type Row struct {
	style           *style.Style
	Style           string
	id              component.Identifier
	Id              string
	Class           string
	themeLoader     theme.Loader
	children        component.ComponentGroup
	componentLoader theme.ComponentLoader
}

func NewRow(loader theme.Loader, parent component.Component, attributes map[string]string) (component.Component, error) {
	row := &Row{
		themeLoader: loader,
		children:    component.NewComponentGroup([]component.Component{}),
	}
	row.componentLoader = theme.NewComponentLoader(loader, "container")
	mapstructure.WeakDecode(attributes, row)
	row.id = component.NewIdentifier(row.Class, row.Id, RowElem, component.StateDefault)
	row.style = style.FromCss([]byte(row.Style), loader)
	return row, nil
}

func (r *Row) Identify() component.Identifier {
	return r.id
}

func (w *Row) Children() component.ComponentGroup {
	return w.children
}

func (g *Row) Update(ctx context.Context) ([]byte, error) {
	absoluteStyle := g.style.Build(ctx)
	children := g.setupChildren(absoluteStyle.InnerCtx)
	
	return g.componentLoader.Load(ctx, Tpl{
		Style:    absoluteStyle,
		Children: children,
	})
}

func (g *Row) setupChildren(ctx context.Context) string {
	children := ""
	
	g.handleChildren(ctx, func(child *Column, childCtx context.Context) {
		c, _ := child.Update(childCtx)
		children += string(c)
	})
	return children
}

func (g *Row) handleChildren(ctx context.Context, callback func(child *Column, childCtx context.Context)) {
	totalSize := g.GetTotalSize()
	totalOffset := 0.0
	g.Children().Each(func(child component.Component) {
		widthP := 100.0 / totalSize * child.(*Column).Size
		width, _ := strconv.ParseFloat(style.Width(ctx, fmt.Sprint(widthP, "%")), 64)
		childCtx := style.BoxContext(ctx, fmt.Sprint(totalOffset), "0", fmt.Sprint(width), "100%")
		callback(child.(*Column), childCtx)
		totalOffset += width
	})
}

func (g *Row) GetTotalSize() int {
	size := 0
	g.Children().Each(func(child component.Component) {
		size += child.(*Column).Size
	})
	return size
}

func (g *Row) Rect(ctx context.Context) sdl.FRect {
	absoluteStyle := g.style.Build(ctx)
	rect := style.Rect(absoluteStyle.OuterCtx, "0", "0", "100%", "0")
	var h float32 = 0
	g.handleChildren(ctx, func(child *Column, childCtx context.Context) {
		childRect := child.Rect(childCtx)
		if childRect.H > h {
			h = childRect.H
		}
	})
	rect.H = h
	return rect
}

func (r *Row) Styler() *style.Style {
	return r.style
}
