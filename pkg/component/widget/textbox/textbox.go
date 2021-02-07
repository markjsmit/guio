package textbox

import (
	"context"
	
	"github.com/mitchellh/mapstructure"
	"github.com/veandco/go-sdl2/sdl"
	
	"github.com/maxpower89/guio/pkg/component"
	"github.com/maxpower89/guio/pkg/event"
	"github.com/maxpower89/guio/pkg/style"
	"github.com/maxpower89/guio/pkg/theme"
)

const Element = "textbox"

type Textbox struct {
	Style string
	id    component.Identifier
	Id    string
	Class string
	Text  string
	
	style        *style.Style
	active       bool
	cursorPos    int
	blinkCounter int
	blinkState   bool
	
	themeLoader     theme.Loader
	componentLoader theme.ComponentLoader
	mouseHandler    event.MouseHandler
	keyboardHandler event.KeyboardHandler
}

func NewTextBox(loader theme.Loader, parent component.Component, attributes map[string]string) (component.Component, error) {
	textbox := &Textbox{
		themeLoader:     loader,
		mouseHandler:    event.NewMouseHandler(),
		keyboardHandler: event.NewKeyboardHandler(),
	}
	
	mapstructure.WeakDecode(attributes, textbox)
	textbox.style = style.FromCss([]byte(textbox.Style), loader)
	textbox.componentLoader = theme.NewComponentLoader(loader, "textbox")
	
	textbox.id = component.NewIdentifier(textbox.Class, textbox.Id, Element, component.StateDefault)
	return textbox, nil
}

func (b *Textbox) Listen(key interface{}, listener event.Callback) {
	b.mouseHandler.Listen(key, listener)
}

func (b *Textbox) Identify() component.Identifier {
	return b.id
}

func (b *Textbox) RegisterWindowHandler(handler event.Handler) {
	b.mouseHandler.Bind(handler)
	b.mouseHandler.Listen(event.MouseClick{}, b.activate)
	b.mouseHandler.Listen(event.Blur{}, b.deactivate)
	b.keyboardHandler.Bind(handler)
	b.keyboardHandler.Listen(event.KeyDown{}, b.keyDown)
	b.keyboardHandler.Listen(event.KeyChar{}, b.keyChar)
}

func (b *Textbox) keyDown(ev interface{}) {
	e := ev.(event.KeyDown)
	if e.Scancode == sdl.SCANCODE_LEFT {
		b.navigate(-1)
	} else if e.Scancode == sdl.SCANCODE_RIGHT {
		b.navigate(1)
	} else if e.Scancode == sdl.SCANCODE_BACKSPACE {
		b.delete()
	}
}

func (b *Textbox) keyChar(ev interface{}) {
	e := ev.(event.KeyChar)
	b.insert(string(e.Rune))
}


func (b *Textbox) navigate(i int) {
	newPos := b.cursorPos + i
	if newPos < 0 {
		newPos = 0
	} else if newPos > len(b.Text) {
		newPos = len(b.Text)
	}
	b.cursorPos = newPos
}

func (b *Textbox) activate(event interface{}) {
	b.id.ChangeState(component.StateActive)
	b.active = true
	b.cursorPos = len(b.Text)
}

func (b *Textbox) deactivate(event interface{}) {
	b.id.ChangeState(component.StateDefault)
	b.active = false
}

func (b *Textbox) Update(ctx context.Context) ([]byte, error) {
	absoluteStyle := b.style.Build(ctx)
	b.mouseHandler.UpdateBox(ctx, b.Rect(absoluteStyle.Ctx))
	b.keyboardHandler.UpdateState(b.active)
	
	return b.componentLoader.Load(ctx, Tpl{
		Style:      absoluteStyle,
		Text:       b.Text,
		CursorPos:  b.cursorPos,
		Active:     b.active,
		BlinkState: b.calculateBlinkState(),
	})
}

func (b *Textbox) Rect(ctx context.Context) sdl.FRect {
	absoluteStyle := b.style.Build(ctx)
	return absoluteStyle.OuterBox
}

func (b *Textbox) Styler() *style.Style {
	return b.style
}

func (b *Textbox) calculateBlinkState() bool {
	if !b.active {
		return false
	}
	b.blinkCounter = (b.blinkCounter + 1) % 24
	if b.blinkCounter == 0 {
		b.blinkState = !b.blinkState
	}
	return b.blinkState
	
}


func (b *Textbox) insert(char string) {
	if b.cursorPos == 0 {
		b.Text = char + b.Text
	} else if b.cursorPos == len(b.Text) {
		b.Text = b.Text + char
	} else {
		b.Text = b.Text[0:b.cursorPos] + char + b.Text[b.cursorPos:len(b.Text)]
	}
	b.cursorPos++
}

func (b *Textbox) delete() {
	if b.cursorPos > 0 {
		if b.cursorPos == len(b.Text) {
			b.Text = b.Text[0 : len(b.Text)-1]
		} else {
			b.Text = b.Text[0:b.cursorPos-1] + b.Text[b.cursorPos:len(b.Text)]
		}
		b.cursorPos--
	}
}

