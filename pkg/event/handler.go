package event

type Listenable interface {
	Listen(key interface{}) Listener
}

type Dispatcher interface {
	Dispatch(key interface{}, data interface{})
}

type Handler interface {
	Listenable
	Dispatcher
	Dispose(listener Listener)
}

type handler struct {
	events map[interface{}]map[string]Listener
}

func (h *handler) Listen(key interface{}) Listener {
	if _, ok := h.events[key]; !ok {
		h.events[key] = map[string]Listener{}
	}
	listener := NewListener(h, key)
	h.events[key][listener.Id()] = listener
	return listener
}

func (h *handler) Dispatch(key interface{}, data interface{}) {
	if listeners, ok := h.events[key]; ok {
		for _, listener := range listeners {
			listener.Trigger(data)
		}
	}
}

func NewHandler() *handler {
	return &handler{
		events: map[interface{}]map[string]Listener{},
	}
}

func (h *handler) Dispose(listener Listener) {
	delete(h.events[listener.Key()], listener.Id())
}
