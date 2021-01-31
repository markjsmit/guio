package component

import (
	"context"
	
	"github.com/veandco/go-sdl2/sdl"
	
	"github.com/maxpower89/guio/pkg/event"
	"github.com/maxpower89/guio/pkg/theme"
)

type NewComponent func(loader theme.Loader, attributes map[string]string) (Component, error)

type Component interface {
	Update(ctx context.Context) ([]byte, error)
	Rect(ctx context.Context) sdl.Rect
}

type Reactive interface {
	RegisterWindowHandler(handler event.Handler)
}
type Identifiable interface {
	Identify() string
}

type Container interface {
	AddChild(component Component)
	GetChildren() []Component
}

type WindowMeta struct {
	Width  int
	Height int
	Title  string
}

type RootComponent interface {
	Component
	Container
	Reactive
	UpdateMeta(meta WindowMeta)
	Meta() WindowMeta
}
