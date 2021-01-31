package window

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"
	
	"github.com/maxpower89/guio/pkg/component"
	"github.com/maxpower89/guio/pkg/element"
	"github.com/maxpower89/guio/pkg/render"
	"github.com/maxpower89/guio/pkg/theme"
)

type Loader interface {
	Load(ctx context.Context, name string) (Window, error)
	RegisterComponent(key string, handler component.NewComponent)
}

type loader struct {
	components      map[string]component.NewComponent
	themeLoader     theme.Loader
	windowDirectory string
	renderer        render.Renderer
}

func NewLoader(windowDirectory string, themeLoader theme.Loader, renderer render.Renderer) Loader {
	return &loader{
		themeLoader:     themeLoader,
		windowDirectory: windowDirectory,
		renderer:        renderer,
		components:      map[string]component.NewComponent{},
	}
}

func (l *loader) Load(ctx context.Context, name string) (Window, error) {
	filename := fmt.Sprint(l.windowDirectory, "/", name, ".xml")
	fileContent, err := ioutil.ReadFile(filename)
	w := NewWindow(ctx, l.renderer)
	if err != nil {
		return w, err
	}
	
	windowElement, err := element.FromXml(fileContent, "window")
	if err != nil {
		return w, err
	}
	
	windowComponent, _ := component.NewWindow(l.themeLoader, windowElement.Attr)
	container := windowComponent.(component.Container)
	l.fillContainer(container, windowElement.Children)
	w.SetRootComponent(windowComponent.(component.RootComponent))
	
	return w, nil
}

func (w *loader) RegisterComponent(key string, handler component.NewComponent) {
	w.components[strings.ToLower(key)] = handler
}

type componentElement struct {
	key     string
	element interface{}
}

func (l *loader) fillContainer(container component.Container, elements []element.Element) {
	for _, elem := range elements {
		if constructor, ok := l.components[strings.ToLower(elem.Tag)]; ok {
			c, err := constructor(l.themeLoader, elem.Attr)
			if err == nil {
				if componentContainer, ok := c.(component.Container); ok {
					l.fillContainer(componentContainer, elem.Children)
				}
				container.AddChild(c)
			}
		}
	}
}
