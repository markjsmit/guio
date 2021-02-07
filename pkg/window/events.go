package window

import (
	"github.com/tfriedel6/canvas/sdlcanvas"
	
	"github.com/maxpower89/guio/pkg/component"
	"github.com/maxpower89/guio/pkg/event"
)

func setupEvents(wnd *sdlcanvas.Window) event.WindowHandler {
	eventHandler := event.NewWindowHandler()
	eventHandler.Bind(wnd)
	return eventHandler
}

func bindEvents(w *window, handler event.WindowHandler) {
	w.componentGroup.EachRecursive(func(c component.Component) {
		if reactive, ok := c.(component.Reactive); ok {
			reactive.RegisterWindowHandler(handler)
		}
	})
}
