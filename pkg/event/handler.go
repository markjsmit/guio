package event

type Callback func(event interface{})

type Listenable interface {
	Listen(key interface{}, callback Callback)
}

type Dispatcher interface {
	Dispatch(key interface{}, data interface{})
}

type Handler interface {
	Listenable
	Dispatcher
}

type handler struct {
	events map[interface{}][]Callback
}

func (h *handler) Listen(key interface{}, listener Callback) {
	if _, ok := h.events[key]; !ok {
		h.events[key] = []Callback{}
	}
	
	h.events[key] = append(h.events[key], listener)
}

func (h *handler) Dispatch(key interface{}, data interface{}) {
	if listeners, ok := h.events[key]; ok {
		for _, listener := range listeners {
			listener(data)
		}
	}
}

func NewHandler() *handler {
	return &handler{
		events: map[interface{}][]Callback{},
	}
}
