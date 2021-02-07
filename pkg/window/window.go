package window

import (
	"context"
	
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/sdlcanvas"
	
	"github.com/maxpower89/guio/pkg/binding"
	"github.com/maxpower89/guio/pkg/component"
	"github.com/maxpower89/guio/pkg/event"
	"github.com/maxpower89/guio/pkg/render"
	"github.com/maxpower89/guio/pkg/style"
)

type Window interface {
	event.Listenable
	Open() error
	Close()
	WithHandler(handler Handler) Window
	SetRootComponent(handler component.RootComponent)
	Bind(binding binding.Binding, err error) error
	GetCanvas() *canvas.Canvas
	GetSdlWindow() *sdlcanvas.Window
}

type window struct {
	ctx            context.Context
	wnd            *sdlcanvas.Window
	rootComponent  component.RootComponent
	renderer       render.Renderer
	handler        Handler
	componentGroup component.ComponentGroup
	stylesheet     *style.Stylesheet
	canvas         *canvas.Canvas
	binder         binding.Binder
	eventHandler   event.Handler
}

func NewWindow(ctx context.Context, renderer render.Renderer) Window {
	return &window{
		ctx:      ctx,
		renderer: renderer,
	}
}

func (w *window) Open() error {
	meta, wnd, canvas, err := w.setup()
	if err != nil {
		return err
	}
	defer wnd.Destroy()
	
	wnd.MainLoop(func() {
		w.update(wnd, meta, canvas)
	})
	
	return err
}

func (w *window) update(wnd *sdlcanvas.Window, meta component.WindowMeta, canvas *canvas.Canvas) {
	updateStyling(w)
	sw, sh := wnd.Window.GetSize()
	if meta.Width != int(sw) || meta.Height != int(sh) {
		meta.Width = int(sw)
		meta.Height = int(sh)
		w.rootComponent.UpdateMeta(meta)
		w.eventHandler.Dispatch(event.Resize{}, event.Resize{
			Width: sw, Height: sh,
		})
	}
	w.eventHandler.Dispatch(event.Update{}, event.Update{})
	w.binder.Update()
	renderData, _ := w.rootComponent.Update(w.ctx)
	w.renderer.Render(w.ctx, canvas, renderData)
}

func (w *window) setup() (component.WindowMeta, *sdlcanvas.Window, *canvas.Canvas, error) {
	meta := w.rootComponent.Meta()
	wnd, canvas, err := sdlcanvas.CreateWindow(meta.Width, meta.Height, meta.Title)
	w.wnd = wnd
	w.canvas = canvas
	if err != nil {
		return meta, wnd, canvas, err
	}
	
	Binder := binding.NewBinder()
	w.binder = Binder
	eventHandler := setupEvents(wnd)
	w.eventHandler = eventHandler.(event.Handler)
	setuphandler(w)
	bindEvents(w, eventHandler)
	bindStyle(w)
	
	return meta, wnd, canvas, err
}


func (w *window) Close() {
	w.wnd.Destroy()
}

func (w *window) Listen(key interface{}, callback event.Callback) {
	w.eventHandler.Listen(key, callback)
}

func (w *window) GetCanvas() *canvas.Canvas {
	return w.canvas
}

func (w *window) GetSdlWindow() *sdlcanvas.Window {
	return w.wnd
}

func (w window) Bind(binding binding.Binding, err error) error {
	if err != nil {
		return err
	}
	w.binder.Add(binding)
	return nil
}

func (w window) WithHandler(handler Handler) Window {
	w.handler = handler
	return &w
}

func (w *window) SetRootComponent(c component.RootComponent) {
	w.rootComponent = c
	w.componentGroup = component.NewComponentGroup([]component.Component{w.rootComponent})
}
