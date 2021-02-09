package event

import "github.com/google/uuid"

type Callback func(event interface{})

type Listener interface {
	Callback(callback Callback)
	Trigger(data interface{})
	Dispose()
	Id() string
	Key() interface{}
}

type listener struct {
	cb      Callback
	handler Handler
	id      string
	key     interface{}
}

func NewListener(handler Handler, key interface{}) *listener {
	ider := uuid.New()
	return &listener{
		cb:      func(event interface{}) {},
		handler: handler,
		id:      ider.String(),
		key:     key,
	}
}

func (l *listener) Callback(callback Callback) {
	l.cb = callback
}

func (l *listener) Trigger(data interface{}) {
	l.cb(data)
}

func (l *listener) Dispose() {
	l.handler.Dispose(l)
}

func (l *listener) Id() string {
	return l.id
}
func (l *listener) Key() interface{} {
	return l.key
}
