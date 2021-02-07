package grid

import (
	"context"
	
	"github.com/mitchellh/mapstructure"
	"github.com/veandco/go-sdl2/sdl"
	
	"github.com/maxpower89/guio/pkg/component"
	"github.com/maxpower89/guio/pkg/style"
	"github.com/maxpower89/guio/pkg/theme"
)

const ColumnElem = "grid-column"

type Column struct {
	style           *style.Style
	Style           string
	id              component.Identifier
	Id              string
	Class           string
	themeLoader     theme.Loader
	children        component.ComponentGroup
	Size            int
	Padding         string
	componentLoader theme.ComponentLoader
}

func NewColumn(loader theme.Loader, parent component.Component, attributes map[string]string) (component.Component, error) {
	column := &Column{
		Size:        1,
		themeLoader: loader,
		children:    component.NewComponentGroup([]component.Component{}),
	}
	column.componentLoader = theme.NewComponentLoader(loader, "container")
	mapstructure.WeakDecode(attributes, column)
	column.id = component.NewIdentifier(column.Class, column.Id, ColumnElem,component.StateDefault)
	column.style = style.FromCss([]byte(column.Style), loader)
	return column, nil
}

func (c *Column) Identify() component.Identifier {
	return c.id
}

func (c *Column) Children() component.ComponentGroup {
	return c.children
}

func (c Column) Update(ctx context.Context) ([]byte, error) {
	absoluteStyle := c.style.Build(ctx)
	children := c.setupChildren(absoluteStyle.InnerCtx)
	
	return c.componentLoader.Load(ctx,  Tpl{
		Style:    absoluteStyle,
		Children: children,
	})
}

func (c Column) setupChildren(ctx context.Context) string {
	children := ""
	c.handleChildren(ctx, func(child component.Component, childCtx context.Context) {
		c, _ := child.Update(childCtx)
		children += string(c)
	})
	return children
}

func (c Column) handleChildren(ctx context.Context, callback func(child component.Component, childCtx context.Context)) {
	c.Children().Each(func(child component.Component) {
		callback(child, ctx)
	})
}

func (c *Column) Rect(ctx context.Context) sdl.FRect {
	absoluteStyle := c.style.Build(ctx)
	rect := style.Rect(absoluteStyle.OuterCtx, "0", "0", "100%", "0")
	var h float32 = 0
	c.handleChildren(absoluteStyle.InnerCtx, func(child component.Component, childCtx context.Context) {
		childRect := child.Rect(childCtx)
		relY := childRect.Y - rect.Y
		relH := childRect.H + relY + absoluteStyle.Margin.Bottom + absoluteStyle.Padding.Bottom + absoluteStyle.Margin.Top
		if h < relH {
			h = relH
		}
	})
	rect.H = h
	return rect
}

func (c *Column) Styler() *style.Style {
	return c.style
}
