package event

import (
	"context"
	
	"github.com/veandco/go-sdl2/sdl"
)

type MouseHandler interface {
	Handler
	UpdateMousePos(point sdl.FPoint)
	UpdateBox(ctx context.Context, rect sdl.FRect)
	UpdateButtonState(state uint8, button uint8)
	IsInBox() bool
	IsMouseDown() bool
	Bind(h Handler)
}

type mouseHandler struct {
	internalHandler Handler
	buttonState     uint8
	button          uint8
	isInBox         bool
	pos             sdl.FPoint
	relPos          sdl.FPoint
	box             sdl.FRect
}

func (m *mouseHandler) Bind(h Handler) {
	h.Listen(MouseMove{}, func(e interface{}) {
		me := e.(MouseMove)
		m.UpdateMousePos(me.Pos)
	})
	
	h.Listen(MouseButton{}, func(e interface{}) {
		me := e.(MouseButton)
		m.UpdateButtonState(me.State, me.Button)
	})
}

func (m *mouseHandler) UpdateButtonState(state uint8, button uint8) {
	m.update(m.pos, m.box, state, button)
}

func (m *mouseHandler) UpdateMousePos(point sdl.FPoint) {
	m.update(point, m.box, m.buttonState, m.button)
}

func (m *mouseHandler) UpdateBox(ctx context.Context, rect sdl.FRect) {
	
	m.update(m.pos, rect, m.buttonState, m.button)
	
}

func (m *mouseHandler) update(pos sdl.FPoint, box sdl.FRect, buttonState uint8, button uint8) {
	oldButtonState := m.buttonState
	oldIsInBox := m.isInBox
	
	m.box = box
	m.button = button
	m.buttonState = buttonState
	m.isInBox = pos.InRect(&box)
	m.pos = pos
	
	if m.isInBox {
		m.relPos = sdl.FPoint{
			X: m.pos.X - box.X,
			Y: m.pos.Y - box.Y,
		}
		
		if m.buttonState != oldButtonState {
			m.Dispatch(MouseButton{}, MouseButton{
				Pos:    m.pos,
				RelPos: m.relPos,
				Button: m.button,
				State:  m.buttonState,
			})
			
			if m.buttonState == sdl.PRESSED {
				m.Dispatch(MouseDown{}, MouseDown{Button: m.button})
			} else {
				m.Dispatch(MouseClick{}, MouseUp{Button: m.button})
				m.Dispatch(MouseUp{}, MouseUp{Button: m.button})
			}
		}
		
		m.internalHandler.Dispatch(MouseMove{}, MouseMove{
			Pos:    m.pos,
			RelPos: m.relPos,
		})
		if !oldIsInBox {
			m.internalHandler.Dispatch(MouseEnter{}, MouseEnter{})
		}
	} else {
		m.relPos = sdl.FPoint{}
		
		if m.buttonState != oldButtonState && m.buttonState == sdl.RELEASED {
			m.Dispatch(Blur{}, Blur{})
		}
		
		if oldIsInBox {
			m.internalHandler.Dispatch(MouseLeave{}, MouseLeave{})
		}
	}
	
}

func NewMouseHandler() MouseHandler {
	return &mouseHandler{
		internalHandler: NewHandler(),
	}
}

func (m mouseHandler) Listen(key interface{}, callback Callback) {
	m.internalHandler.Listen(key, callback)
}

func (m mouseHandler) Dispatch(key interface{}, data interface{}) {
	m.internalHandler.Dispatch(key, data)
}

func (m mouseHandler) IsInBox() bool {
	return m.isInBox
}

func (m mouseHandler) IsMouseDown() bool {
	return m.buttonState == sdl.PRESSED
}

type MouseMove struct {
	Pos    sdl.FPoint
	RelPos sdl.FPoint
}

type MouseEnter struct{}

type MouseLeave struct{}

type MouseClick struct {
	Button uint8
}


type MouseDown struct {
	Button uint8
}

type MouseUp struct {
	Button uint8
}

type MouseButton struct {
	Pos    sdl.FPoint
	RelPos sdl.FPoint
	Button uint8
	State  uint8
}
