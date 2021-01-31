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
	windowId uint32
	handler  Handler
}

func NewWindowHandler(windowId uint32) *windowHandler {
	return &windowHandler{windowId: windowId, handler: NewHandler()}
}

func (e *windowHandler) Bind(wnd *sdlcanvas.Window) {
	wnd.MouseMove = e.handleMouseMove
	wnd.MouseDown = e.handleMouseDown
	wnd.MouseUp = e.handleMouseUp
}

func (e *windowHandler) handleMouseMove(x, y int) {
	{
		e.Dispatch(MouseMove{}, MouseMove{Pos: sdl.Point{
			X: int32(x),
			Y: int32(y),
		}, RelPos: sdl.Point{
			X: int32(x),
			Y: int32(y),
		}})
	}
}

func (e *windowHandler) handleMouseDown(button int, x int, y int) {
	
	e.Dispatch(MouseButton{}, MouseButton{
		Button: uint8(button),
		State:  sdl.PRESSED,
		Pos: sdl.Point{
			X: int32(x),
			Y: int32(y),
		}})
}

func (e *windowHandler) handleMouseUp(button int, x int, y int) {
	
	e.Dispatch(MouseButton{}, MouseButton{
		Button: uint8(button),
		State:  sdl.RELEASED,
		Pos: sdl.Point{
			X: int32(x),
			Y: int32(y),
		}})
}

func (e *windowHandler) Listen(key interface{}, listener Callback) {
	e.handler.Listen(key, listener)
}

func (e *windowHandler) Dispatch(key interface{}, data interface{}) {
	e.handler.Dispatch(key, data)
}
