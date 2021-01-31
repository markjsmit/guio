package window

import (
	"context"
	
	"github.com/tfriedel6/canvas/sdlcanvas"
	"github.com/veandco/go-sdl2/sdl"
	
	"github.com/maxpower89/guio/pkg/component"
	"github.com/maxpower89/guio/pkg/event"
	"github.com/maxpower89/guio/pkg/render"
)

type Window interface {
	Open() error
	Close()
	WithHandler(handler Handler) Window
	SetRootComponent(handler component.RootComponent)
}

type window struct {
	ctx           context.Context
	wnd           *sdlcanvas.Window
	rootComponent component.RootComponent
	renderer      render.Renderer
	handler       Handler
}

func (w window) WithHandler(handler Handler) Window {
	w.handler = handler
	return &w
}

func (w *window) SetRootComponent(c component.RootComponent) {
	w.rootComponent = c
}

func NewWindow(ctx context.Context, renderer render.Renderer) Window {
	return &window{
		ctx:      ctx,
		renderer: renderer,
	}
}

func (w window) Open() error {
	meta := w.rootComponent.Meta()
	wnd, canvas, err := sdlcanvas.CreateWindow(meta.Width, meta.Height, meta.Title)
	if err != nil {
		return err
	}
	
	eventHandler := event.NewWindowHandler(wnd.WindowID)
	defer wnd.Destroy()
	w.wnd = wnd
	
	w.rootComponent.(component.Reactive).RegisterWindowHandler(eventHandler)
	
	if w.handler != nil {
		bindHandler(w.rootComponent, w.handler)
	}
	w.handler.Init(w.ctx)
	eventHandler.Bind(wnd)
	wnd.MainLoop(func() {
		sw, sh := wnd.Window.GetSize()
		if meta.Width != int(sw) || meta.Height != int(sh) {
			meta.Width = int(sw)
			meta.Height = int(sh)
			w.rootComponent.UpdateMeta(meta)
		}
		renderData, _ := w.rootComponent.Update(w.ctx)

		if err == nil {
		
		}
		w.renderer.Render(w.ctx, canvas, renderData)
		
		sdl.Delay(1000 / 60)
	})

	
	return err
}

func (w window) Close() {
	w.wnd.Destroy()
}
