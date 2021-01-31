package window

import (
	"context"
	"reflect"
	
	"github.com/maxpower89/guio/pkg/component"
)

type Handler interface {
	Init(context context.Context)
}

func bindHandler(rootComponent component.Component, handler Handler)  {
	hv := reflect.ValueOf(handler)
	for hv.Kind() == reflect.Interface || hv.Kind() == reflect.Ptr {
		hv = hv.Elem()
	}
	
	if hv.Kind() == reflect.Struct {
		idMap := getIdMap(rootComponent, map[string]component.Component{})
		for i := 0; i < hv.NumField(); i++ {
			ft := hv.Type().Field(i)
			fv := hv.Field(i)
			if tag, ok := ft.Tag.Lookup("component"); ok {
				if component, ok := idMap[tag]; ok {
					fv.Set(reflect.ValueOf(component))
				}
			}
		}
	}
}

func getIdMap(c component.Component, m map[string]component.Component) map[string]component.Component {
	
	if ider, ok := c.(component.Identifiable); ok && ider.Identify() != "" {
		m[ider.Identify()] = c
	}
	
	if container, ok := c.(component.Container); ok {
		for _, child := range container.GetChildren() {
			m = getIdMap(child, m)
		}
	}
	
	return m
}
