package component

import (
	"context"
	"fmt"
	"strconv"
	
	"github.com/mitchellh/mapstructure"
	"github.com/veandco/go-sdl2/sdl"
	
	"github.com/maxpower89/guio/pkg/event"
	"github.com/maxpower89/guio/pkg/style"
	"github.com/maxpower89/guio/pkg/theme"
)

type Window struct {
	Padding      string
	Width        string
	Height       string
	Title        string
	themeLoader  theme.Loader
	children     []Component
	eventHandler event.Handler
}

func (w *Window) RegisterWindowHandler(handler event.Handler) {
	for _, child := range w.children {
		if reactive, ok := child.(Reactive); ok {
			reactive.RegisterWindowHandler(handler)
		}
	}
}

func (w Window) Meta() WindowMeta {
	width, _ := strconv.Atoi(w.Width)
	height, _ := strconv.Atoi(w.Height)
	return WindowMeta{
		Width:  width,
		Height: height,
		Title:  w.Title,
	}
}

func (w *Window) UpdateMeta(meta WindowMeta) {
	w.Width = fmt.Sprint(meta.Width)
	w.Height = fmt.Sprint(meta.Height)
	w.Title = fmt.Sprint(meta.Title)
}

type WindowAttributes struct {
}

func NewWindow(loader theme.Loader, attributes map[string]string) (Component, error) {
	window := &Window{
		Width:       "800",
		Height:      "600",
		Padding:     "10",
		Title:       "",
		themeLoader: loader,
		children:    []Component{},
	}
	
	mapstructure.WeakDecode(attributes, window)
	return window, nil
}

func (w Window) Update(ctx context.Context) ([]byte, error) {
	children := ""
	
	ctx = style.BoxContext(ctx, "0", "0", w.Width, w.Height)
	
	pt, pr, pb, pl := style.PaddingString(w.Padding)
	childCtx := style.PadBoxContext(ctx, pt, pr, pb, pl)
	for _, child := range w.children {
		c, _ := child.Update(childCtx)
		children += string(c)
	}
	
	return w.themeLoader.Load(ctx, "window", map[string]string{
		"width":    style.Width(ctx, "100%"),
		"height":   style.Height(ctx, "100%"),
		"children": children,
	})
}

func (w *Window) AddChild(component Component) {
	w.children = append(w.children, component)
}

func (w *Window) GetChildren() []Component {
	return w.children
}

func (b *Window) Rect(ctx context.Context) sdl.Rect {
	return style.Rect(ctx, "0", "0", "100%", "100%")
}
