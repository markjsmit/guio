package binding

import "github.com/maxpower89/guio/pkg/conversion"

type Binding interface {
	update()
}

func T(a interface{}, b interface{}) (Binding, error) {
	la, err := GetChangeListener(a)
	if err != nil {
		return nil, err
	}
	
	lb, err := GetChangeListener(b)
	if err != nil {
		return nil, err
	}
	
	va, _ := la.get()
	vb, _ := lb.get()
	catob, cbtoa, err := conversion.GetConversionFunc(va, vb)
	
	if err != nil {
		return nil, err
	}
	
	lb.set(catob(va))
	
	return &genericBinding{
		la:    la,
		lb:    lb,
		catob: catob,
		cbtoa: cbtoa,
	}, nil
}

type genericBinding struct {
	la    ChangeListener
	lb    ChangeListener
	catob conversion.ConversionFunc
	cbtoa conversion.ConversionFunc
}

func (b *genericBinding) update() {
	va, changed := b.la.get()
	if changed {
		b.lb.set(b.catob(va))
		return
	}
	
	vb, changed := b.lb.get()
	if changed {
		b.la.set(b.cbtoa(vb))
	}
	
}
