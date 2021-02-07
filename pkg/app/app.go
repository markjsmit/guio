package app

import (
	"context"
	
	"github.com/maxpower89/guio/pkg/component"
	"github.com/maxpower89/guio/pkg/component/container/grid"
	window2 "github.com/maxpower89/guio/pkg/component/container/window"
	"github.com/maxpower89/guio/pkg/component/widget/button"
	"github.com/maxpower89/guio/pkg/component/widget/textbox"
	"github.com/maxpower89/guio/pkg/render"
	"github.com/maxpower89/guio/pkg/shape"
	"github.com/maxpower89/guio/pkg/theme"
	"github.com/maxpower89/guio/pkg/window"
)

type App interface {
	NewWindow(ctx context.Context, name string) (window.Window, error)
	RegisterComponent(key string, handler component.NewComponent)
	RegisterShape(key string, handler shape.NewComponent)
}

type app struct {
	windowLoader window.Loader
	renderer     render.Renderer
}

func (a *app) RegisterComponent(key string, handler component.NewComponent) {
	a.windowLoader.RegisterComponent(key, handler)
}

func NewApp(themeDirectory string, windowDirectory string) App {
	
	themeLoader := theme.NewLoader(themeDirectory)
	renderer := render.NewRenderer(themeLoader)
	windowLoader := window.NewLoader(windowDirectory, themeLoader, renderer)
	
	a := &app{
		windowLoader: windowLoader,
		renderer:     renderer,
	}
	
	a.RegisterComponent(button.Element, button.NewButton)
	a.RegisterComponent(textbox.Element, textbox.NewTextBox)
	a.RegisterComponent(window2.Elem, window2.NewWindow)
	a.RegisterComponent(grid.Elem, grid.NewGrid)
	a.RegisterComponent(grid.RowElem, grid.NewRow)
	a.RegisterComponent(grid.ColumnElem, grid.NewColumn)
	
	a.RegisterShape("rect", shape.NewRect)
	a.RegisterShape("text", shape.NewText)
	a.RegisterShape("shape", shape.NewShape)
	a.RegisterShape("line", shape.NewLine)
	
	return a
}

func (a *app) NewWindow(ctx context.Context, name string) (window.Window, error) {
	return a.windowLoader.Load(ctx, name)
}

func (a *app) RegisterShape(key string, handler shape.NewComponent) {
	a.renderer.RegisterComponent(key, handler)
}
