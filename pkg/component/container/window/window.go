package window

import (
	"context"
	"fmt"
	"strconv"
	
	"github.com/mitchellh/mapstructure"
	"github.com/veandco/go-sdl2/sdl"
	
	"github.com/maxpower89/guio/pkg/component"
	"github.com/maxpower89/guio/pkg/event"
	"github.com/maxpower89/guio/pkg/style"
	"github.com/maxpower89/guio/pkg/theme"
)

const Elem = "window"

type Window struct {
	style           *style.Style
	Style           string
	id              component.Identifier
	Id              string
	Class           string
	Title           string
	themeLoader     theme.Loader
	children        component.ComponentGroup
	eventHandler    event.Handler
	componentLoader theme.ComponentLoader
}

func (w Window) Meta() component.WindowMeta {
	width, _ := strconv.Atoi(w.style.Width)
	height, _ := strconv.Atoi(w.style.Height)
	return component.WindowMeta{
		Width:  width,
		Height: height,
		Title:  w.Title,
	}
}

func (w *Window) UpdateMeta(meta component.WindowMeta) {
	w.style.Width = fmt.Sprint(meta.Width)
	w.style.Height = fmt.Sprint(meta.Height)
	w.Title = fmt.Sprint(meta.Title)
}

type WindowAttributes struct {
}

func NewWindow(loader theme.Loader, parent component.Component, attributes map[string]string) (component.Component, error) {
	window := &Window{
		Title:       "",
		themeLoader: loader,
		children:    component.NewComponentGroup([]component.Component{}),
	}
	window.componentLoader = theme.NewComponentLoader(loader, "window")
	
	mapstructure.WeakDecode(attributes, window)
	
	window.id = component.NewIdentifier(window.Class, window.Id, Elem, component.StateDefault)
	window.style = style.FromCss([]byte(window.Style), loader)
	if w, ok := attributes["width"]; ok {
		window.style.Width = w
	}
	if h, ok := attributes["height"]; ok {
		window.style.Height = h
	}
	
	return window, nil
}

func (w Window) Update(ctx context.Context) ([]byte, error) {
	children := ""
	ctx = style.BoxContext(ctx, "0", "0", w.style.Width, w.style.Height)
	absoluteStyle := w.style.Build(ctx)
	
	w.Children().Each(func(child component.Component) {
		c, _ := child.Update(absoluteStyle.InnerCtx)
		children += string(c)
	})
	
	return w.componentLoader.Load(ctx, Tpl{
		Style:    absoluteStyle,
		Children: children,
	})
}

func (w *Window) Identify() component.Identifier {
	return w.id
}

func (w *Window) Children() component.ComponentGroup {
	return w.children
}

func (w *Window) Rect(ctx context.Context) sdl.FRect {
	ctx = style.BoxContext(ctx, "0", "0", w.style.Width, w.style.Height)
	return style.Rect(ctx, "0", "0", "100%", "100%")
}

func (w *Window) Styler() *style.Style {
	return w.style
}
