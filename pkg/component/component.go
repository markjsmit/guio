package component

import (
	"context"
	
	"github.com/veandco/go-sdl2/sdl"
	
	"github.com/maxpower89/guio/pkg/event"
	"github.com/maxpower89/guio/pkg/style"
	"github.com/maxpower89/guio/pkg/theme"
)

type NewComponent func(loader theme.Loader, parent Component, attributes map[string]string) (Component, error)

type Component interface {
	Update(ctx context.Context) ([]byte, error)
	Rect(ctx context.Context) sdl.FRect
	Identify() Identifier
}

type Reactive interface {
	RegisterWindowHandler(handler event.Handler)
}

type Container interface {
	Children() ComponentGroup
}

type Stylable interface {
	Styler() *style.Style
}

type WindowMeta struct {
	Width  int
	Height int
	Title  string
}

type RootComponent interface {
	Component
	Container
	UpdateMeta(meta WindowMeta)
	Meta() WindowMeta
}

