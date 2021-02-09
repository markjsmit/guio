package button

import (
	"context"
	
	"github.com/mitchellh/mapstructure"
	"github.com/veandco/go-sdl2/sdl"
	
	"github.com/maxpower89/guio/pkg/component"
	"github.com/maxpower89/guio/pkg/event"
	"github.com/maxpower89/guio/pkg/style"
	"github.com/maxpower89/guio/pkg/theme"
)

const Element = "button"

type Button struct {
	style           *style.Style
	Style           string
	id              component.Identifier
	Id              string
	Class           string
	Text            string
	themeLoader     theme.Loader
	componentLoader theme.ComponentLoader
	mouseHandler    event.MouseHandler
}

func NewButton(loader theme.Loader, parent component.Component, attributes map[string]string) (component.Component, error) {
	button := &Button{
		themeLoader:  loader,
		mouseHandler: event.NewMouseHandler(),
	}
	mapstructure.WeakDecode(attributes, button)
	button.style = style.FromCss([]byte(button.Style), loader)
	button.componentLoader = theme.NewComponentLoader(loader, "button")
	
	button.id = component.NewIdentifier(button.Class, button.Id, Element, component.StateDefault)
	return button, nil
}

func (b *Button) Listen(key interface{}, listener event.Callback) {
	b.mouseHandler.Listen(key).Callback(listener)
}

func (b *Button) Identify() component.Identifier {
	return b.id
}

func (b *Button) RegisterWindowHandler(handler event.Handler) {
	b.mouseHandler.Bind(handler)
}

func (b *Button) Update(ctx context.Context) ([]byte, error) {
	absoluteStyle := b.style.Build(ctx)
	b.mouseHandler.UpdateBox(ctx, b.Rect(absoluteStyle.Ctx))
	if b.mouseHandler.IsInBox() {
		b.id.ChangeState(component.StateHover)
	} else {
		b.id.ChangeState(component.StateDefault)
	}
	return b.componentLoader.Load(ctx, Tpl{
		Style: absoluteStyle,
		Text:  b.Text,
	})
}

func (b *Button) Rect(ctx context.Context) sdl.FRect {
	absoluteStyle := b.style.Build(ctx)
	return absoluteStyle.OuterBox
}

func (b *Button) Styler() *style.Style {
	return b.style
}
