package event

import (
	"github.com/tfriedel6/canvas/sdlcanvas"
	"github.com/veandco/go-sdl2/sdl"
)

type WindowHandler interface {
	Handler
	Bind(wnd *sdlcanvas.Window)
}

type windowHandler struct {
	handler Handler
	wnd     *sdlcanvas.Window
}

func (e *windowHandler) Dispose(listener Listener) {
	e.handler.Dispose(listener)
}

func NewWindowHandler() *windowHandler {
	return &windowHandler{handler: NewHandler()}
}

func (e *windowHandler) Bind(wnd *sdlcanvas.Window) {
	e.wnd = wnd
	wnd.MouseMove = e.handleMouseMove
	wnd.MouseDown = e.handleMouseDown
	wnd.MouseUp = e.handleMouseUp
	wnd.KeyDown = e.handleKeyDown
	wnd.KeyUp = e.handleKeyUp
	
	wnd.KeyChar = e.handleKeyChar
	
}

func (e *windowHandler) handleMouseMove(x, y int) {
	{
		e.Dispatch(MouseMove{}, MouseMove{Pos: sdl.FPoint{
			X: float32(x),
			Y: float32(y),
		}, RelPos: sdl.FPoint{
			X: float32(x),
			Y: float32(y),
		}})
	}
}

func (e *windowHandler) handleMouseDown(button int, x int, y int) {
	
	e.Dispatch(MouseButton{}, MouseButton{
		Button: uint8(button),
		State:  sdl.PRESSED,
		Pos: sdl.FPoint{
			X: float32(x),
			Y: float32(y),
		}})
}

func (e *windowHandler) handleMouseUp(button int, x int, y int) {
	
	e.Dispatch(MouseButton{}, MouseButton{
		Button: uint8(button),
		State:  sdl.RELEASED,
		Pos: sdl.FPoint{
			X: float32(x),
			Y: float32(y),
		}})
}

func (e *windowHandler) Listen(key interface{}) Listener {
	return e.handler.Listen(key)
}

func (e *windowHandler) Dispatch(key interface{}, data interface{}) {
	e.handler.Dispatch(key, data)
}

func (e *windowHandler) handleKeyDown(scancode int, rn rune, name string) {
	e.handler.Dispatch(KeyDown{}, KeyDown{
		Scancode: scancode,
		Rune:     rn,
		Name:     name,
	})
}

func (e *windowHandler) handleKeyUp(scancode int, rn rune, name string) {
	e.handler.Dispatch(KeyUp{}, KeyUp{
		Scancode: scancode,
		Rune:     rn,
		Name:     name,
	})
}

func (e *windowHandler) handleKeyChar(rn rune) {
	e.handler.Dispatch(KeyChar{}, KeyChar{
		Rune: rn,
	})
}

type Resize struct {
	Width  int32
	Height int32
}

type Update struct {
}
