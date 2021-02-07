package binding

type Binder interface {
	Add(binding Binding)
	Update()
}

type binder struct {

	close    chan bool
	bindings []Binding
}

func (bi *binder) Update() {
	for _, binding := range bi.bindings {
		binding.update()
	}
}

func NewBinder() Binder {
	return &binder{
		close:    make(chan bool),
		bindings: []Binding{},
	}
}

func (bi *binder) Add(binding Binding) {
	bi.bindings = append(bi.bindings, binding)
}
