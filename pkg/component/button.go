package component

import (
	"context"
	
	"github.com/mitchellh/mapstructure"
	"github.com/veandco/go-sdl2/sdl"
	
	"github.com/maxpower89/guio/pkg/event"
	"github.com/maxpower89/guio/pkg/style"
	"github.com/maxpower89/guio/pkg/theme"
)

type Button struct {
	Id                  string
	Text                string
	Width, Height, X, Y string
	themeLoader         theme.Loader
	mouseHandler        event.MouseHandler
}

func (b *Button) Listen(key interface{}, listener event.Callback) {
	b.mouseHandler.Listen(key, listener)
}

func NewButton(loader theme.Loader, attributes map[string]string) (Component, error) {
	button := &Button{
		themeLoader:  loader,
		mouseHandler: event.NewMouseHandler(),
	}
	mapstructure.WeakDecode(attributes, button)
	return button, nil
}

func (b *Button) Identify() string {
	return b.Id
}

func (b *Button) RegisterWindowHandler(handler event.Handler) {
	b.mouseHandler.Bind(handler)
}

func (b *Button) Update(ctx context.Context) ([]byte, error) {
	b.mouseHandler.UpdateBox(ctx, b.Rect(ctx))
	return b.themeLoader.Load(ctx, "button", map[string]string{
		"width":   style.Width(ctx, b.Width),
		"height":  style.Height(ctx, b.Height),
		"x":       style.Left(ctx, b.X),
		"y":       style.Top(ctx, b.Y),
		"isHover": style.Bool(ctx, b.mouseHandler.IsInBox()),
		"text":    b.Text,
	})
}

func (b *Button) Rect(ctx context.Context) sdl.Rect {
	return style.Rect(ctx, b.X, b.Y, b.Width, b.Height)
}
