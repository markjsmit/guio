package binding

import (
	"errors"
	"reflect"
)

type ChangeListener interface {
	get() (value interface{}, changed bool)
	set(value interface{})
}

type changeListenerConstructor = func(pointer interface{}) ChangeListener

var changeListenerKindMap = map[reflect.Kind]changeListenerConstructor{
	reflect.String:  newStringListener,
	reflect.Int:     newintListener,
	reflect.Int64:   newint64Listener,
	reflect.Int32:   newint32Listener,
	reflect.Int16:   newint16Listener,
	reflect.Int8:    newint8Listener,
	reflect.Uint:    newuintListener,
	reflect.Uint64:  newuint64Listener,
	reflect.Uint32:  newuint32Listener,
	reflect.Uint16:  newuint16Listener,
	reflect.Uint8:   newuint8Listener,
	reflect.Float64: newfloat64Listener,
	reflect.Float32: newfloat32Listener,
	reflect.Bool:    newboolListener,
}

func GetChangeListener(value interface{}) (ChangeListener, error) {
	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Ptr {
		return nil, errors.New("Value is not a pointer")
	}
	
	constructor, ok := changeListenerKindMap[v.Elem().Kind()]
	if !ok {
		return nil, errors.New("Listener not found")
	}
	
	return constructor(value), nil
}

type stringListener struct {
	pointer *string
	value   string
}

func newStringListener(pointer interface{}) ChangeListener {
	return &stringListener{pointer: pointer.(*string), value: *(pointer.(*string))}
}

func (l *stringListener) get() (interface{}, bool) {
	if l.value == *l.pointer {
		return l.value, false
	}
	l.value = *l.pointer
	return l.value, true
}

func (l *stringListener) set(value interface{}) {
	l.value = value.(string)
	*l.pointer = value.(string)
}

type intListener struct {
	pointer *int
	value   int
}

func newintListener(pointer interface{}) ChangeListener {
	return &intListener{pointer: pointer.(*int), value: *(pointer.(*int))}
}

func (l *intListener) get() (interface{}, bool) {
	if l.value == *l.pointer {
		return l.value, false
	}
	l.value = *l.pointer
	return l.value, true
}

func (l *intListener) set(value interface{}) {
	l.value = value.(int)
	*l.pointer = value.(int)
}

type int64Listener struct {
	pointer *int64
	value   int64
}

func newint64Listener(pointer interface{}) ChangeListener {
	return &int64Listener{pointer: pointer.(*int64), value: *(pointer.(*int64))}
}

func (l *int64Listener) get() (interface{}, bool) {
	if l.value == *l.pointer {
		return l.value, false
	}
	l.value = *l.pointer
	return l.value, true
}

func (l *int64Listener) set(value interface{}) {
	l.value = value.(int64)
	*l.pointer = value.(int64)
}

type int32Listener struct {
	pointer *int32
	value   int32
}

func newint32Listener(pointer interface{}) ChangeListener {
	return &int32Listener{pointer: pointer.(*int32), value: *(pointer.(*int32))}
}

func (l *int32Listener) get() (interface{}, bool) {
	if l.value == *l.pointer {
		return l.value, false
	}
	l.value = *l.pointer
	return l.value, true
}

func (l *int32Listener) set(value interface{}) {
	l.value = value.(int32)
	*l.pointer = value.(int32)
}

type int16Listener struct {
	pointer *int16
	value   int16
}

func newint16Listener(pointer interface{}) ChangeListener {
	return &int16Listener{pointer: pointer.(*int16), value: *(pointer.(*int16))}
}

func (l *int16Listener) get() (interface{}, bool) {
	if l.value == *l.pointer {
		return l.value, false
	}
	l.value = *l.pointer
	return l.value, true
}

func (l *int16Listener) set(value interface{}) {
	l.value = value.(int16)
	*l.pointer = value.(int16)
}

type int8Listener struct {
	pointer *int8
	value   int8
}

func newint8Listener(pointer interface{}) ChangeListener {
	return &int8Listener{pointer: pointer.(*int8), value: *(pointer.(*int8))}
}

func (l *int8Listener) get() (interface{}, bool) {
	if l.value == *l.pointer {
		return l.value, false
	}
	l.value = *l.pointer
	return l.value, true
}

func (l *int8Listener) set(value interface{}) {
	l.value = value.(int8)
	*l.pointer = value.(int8)
}

type float64Listener struct {
	pointer *float64
	value   float64
}

func newfloat64Listener(pointer interface{}) ChangeListener {
	return &float64Listener{pointer: pointer.(*float64), value: *(pointer.(*float64))}
}

func (l *float64Listener) get() (interface{}, bool) {
	if l.value == *l.pointer {
		return l.value, false
	}
	l.value = *l.pointer
	return l.value, true
}

func (l *float64Listener) set(value interface{}) {
	l.value = value.(float64)
	*l.pointer = value.(float64)
}

type float32Listener struct {
	pointer *float32
	value   float32
}

func newfloat32Listener(pointer interface{}) ChangeListener {
	return &float32Listener{pointer: pointer.(*float32), value: *(pointer.(*float32))}
}

func (l *float32Listener) get() (interface{}, bool) {
	if l.value == *l.pointer {
		return l.value, false
	}
	l.value = *l.pointer
	return l.value, true
}

func (l *float32Listener) set(value interface{}) {
	l.value = value.(float32)
	*l.pointer = value.(float32)
}

type boolListener struct {
	pointer *bool
	value   bool
}

func newboolListener(pointer interface{}) ChangeListener {
	return &boolListener{pointer: pointer.(*bool), value: *(pointer.(*bool))}
}

func (l *boolListener) get() (interface{}, bool) {
	if l.value == *l.pointer {
		return l.value, false
	}
	l.value = *l.pointer
	return l.value, true
}

func (l *boolListener) set(value interface{}) {
	l.value = value.(bool)
	*l.pointer = value.(bool)
}

type uintListener struct {
	pointer *uint
	value   uint
}

func newuintListener(pointer interface{}) ChangeListener {
	return &uintListener{pointer: pointer.(*uint), value: *(pointer.(*uint))}
}

func (l *uintListener) get() (interface{}, bool) {
	if l.value == *l.pointer {
		return l.value, false
	}
	l.value = *l.pointer
	return l.value, true
}

func (l *uintListener) set(value interface{}) {
	l.value = value.(uint)
	*l.pointer = value.(uint)
}

type uint64Listener struct {
	pointer *uint64
	value   uint64
}

func newuint64Listener(pointer interface{}) ChangeListener {
	return &uint64Listener{pointer: pointer.(*uint64), value: *(pointer.(*uint64))}
}

func (l *uint64Listener) get() (interface{}, bool) {
	if l.value == *l.pointer {
		return l.value, false
	}
	l.value = *l.pointer
	return l.value, true
}

func (l *uint64Listener) set(value interface{}) {
	l.value = value.(uint64)
	*l.pointer = value.(uint64)
}

type uint32Listener struct {
	pointer *uint32
	value   uint32
}

func newuint32Listener(pointer interface{}) ChangeListener {
	return &uint32Listener{pointer: pointer.(*uint32), value: *(pointer.(*uint32))}
}

func (l *uint32Listener) get() (interface{}, bool) {
	if l.value == *l.pointer {
		return l.value, false
	}
	l.value = *l.pointer
	return l.value, true
}

func (l *uint32Listener) set(value interface{}) {
	l.value = value.(uint32)
	*l.pointer = value.(uint32)
}

type uint16Listener struct {
	pointer *uint16
	value   uint16
}

func newuint16Listener(pointer interface{}) ChangeListener {
	return &uint16Listener{pointer: pointer.(*uint16), value: *(pointer.(*uint16))}
}

func (l *uint16Listener) get() (interface{}, bool) {
	if l.value == *l.pointer {
		return l.value, false
	}
	l.value = *l.pointer
	return l.value, true
}

func (l *uint16Listener) set(value interface{}) {
	l.value = value.(uint16)
	*l.pointer = value.(uint16)
}

type uint8Listener struct {
	pointer *uint8
	value   uint8
}

func newuint8Listener(pointer interface{}) ChangeListener {
	return &uint8Listener{pointer: pointer.(*uint8), value: *(pointer.(*uint8))}
}

func (l *uint8Listener) get() (interface{}, bool) {
	if l.value == *l.pointer {
		return l.value, false
	}
	l.value = *l.pointer
	return l.value, true
}

func (l *uint8Listener) set(value interface{}) {
	l.value = value.(uint8)
	*l.pointer = value.(uint8)
}
