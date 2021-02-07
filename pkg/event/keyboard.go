package event

type KeyboardHandler interface {
	Handler
	UpdateState(bool)
	Bind(handler Handler)
}

type keyboardHandler struct {
	internalHandler Handler
	active          bool
}

func NewKeyboardHandler() KeyboardHandler {
	return &keyboardHandler{
		internalHandler: NewHandler(),
	}
}

func (k *keyboardHandler) Listen(key interface{}, callback Callback) {
	k.internalHandler.Listen(key, callback)
}

func (k *keyboardHandler) Dispatch(key interface{}, data interface{}) {
	k.internalHandler.Dispatch(key, data)
}

func (k *keyboardHandler) UpdateState(active bool) {
	k.active = active
}

func (k *keyboardHandler) Bind(handler Handler) {
	
	handler.Listen(KeyChar{}, func(event interface{}) {
		if k.active {
			kc := event.(KeyChar)
			k.internalHandler.Dispatch(KeyChar{}, KeyChar{
				Rune: kc.Rune,
			})
		}
	})
	
	handler.Listen(KeyDown{}, func(event interface{}) {
		if k.active {
			kd := event.(KeyDown)
			k.internalHandler.Dispatch(KeyDown{}, KeyDown{
				Scancode: kd.Scancode,
				Rune:     kd.Rune,
				Name:     kd.Name,
			})
		}
	})
	
	handler.Listen(KeyUp{}, func(event interface{}) {
		ku := event.(KeyUp)
		k.internalHandler.Dispatch(KeyUp{}, KeyUp{
			Scancode: ku.Scancode,
			Rune:     ku.Rune,
			Name:     ku.Name,
		})
	})
}

type KeyChar struct {
	Rune rune
}
type KeyUp struct {
	Scancode int
	Rune     rune
	Name     string
}
type KeyDown struct {
	Scancode int
	Rune     rune
	Name     string
}
