package window

import (
	"context"
	"reflect"
	"strings"
	
	"github.com/maxpower89/guio/pkg/component"
)

type Handler interface {
	Init(c context.Context, w Window)
}

func bindHandler(group component.ComponentGroup, handler Handler) {
	hv := reflect.ValueOf(handler)
	for hv.Kind() == reflect.Interface || hv.Kind() == reflect.Ptr {
		hv = hv.Elem()
	}
	
	if hv.Kind() == reflect.Struct {
		for i := 0; i < hv.NumField(); i++ {
			ft := hv.Type().Field(i)
			fv := hv.Field(i)
			if tag, ok := ft.Tag.Lookup("component"); ok {
				setupHandlerField(fv, group, tag)
			}
		}
	}
}

func setupHandlerField(fv reflect.Value, group component.ComponentGroup, tag string) {
	filteredGroup := group.Select(strings.Split(tag, " ")...)
	
	if reflect.TypeOf(filteredGroup).AssignableTo(fv.Type()) {
		fv.Type()
		fv.Set(reflect.ValueOf(filteredGroup))
	} else if fv.Kind() == reflect.Slice {
		slice:=reflect.New(fv.Type()).Elem();
		filteredGroup.Each(func(component component.Component) {
			cv:=reflect.ValueOf(component)
			slice=reflect.Append(slice,cv)
		})
		fv.Set(slice);
	} else if filteredGroup.Length() > 0 {
		first, err := filteredGroup.First()
		if err == nil {
			fv.Set(reflect.ValueOf(first))
		}
	}
}


func setuphandler(w *window) {
	if w.handler != nil {
		bindHandler(w.componentGroup, w.handler)
		w.handler.Init(w.ctx, w)
	}
}
