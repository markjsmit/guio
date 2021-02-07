package component

import "errors"

type ComponentGroup interface {
	Select(selector ...string) ComponentGroup
	AddGroup(group ComponentGroup)
	Components() []Component
	Each(func(component Component))
	EachRecursive(func(component Component))
	Add(component Component)
	First() (Component, error)
	Length() int
}

type componentGroup struct {
	components []Component
}

func (g *componentGroup) Length() int {
	return len(g.components)
}

func (g *componentGroup) First() (Component, error) {
	if len(g.components) == 0 {
		return nil, errors.New("No components in group")
	}
	return g.components[0], nil
}

func NewComponentGroup(components []Component) *componentGroup {
	return &componentGroup{components: components}
}

func (g *componentGroup) Select(selector ...string) ComponentGroup {
	newGroup := NewComponentGroup([]Component{})
	if len(selector) == 0 {
		return newGroup
	}
	
	first := selector[0]
	for _, component := range g.components {
		if component.Identify().HasIdentity(first) {
			if len(selector) == 1 {
				newGroup.Add(component)
			}
			if container, ok := component.(Container); ok {
				if len(selector) == 1 {
					newGroup.AddGroup(container.Children().Select(selector[1:]...))
				}else{
					newGroup.AddGroup(container.Children().Select(selector...))
				}
			}
			
		} else if container, ok := component.(Container); ok {
			newGroup.AddGroup(container.Children().Select(selector...))
		}
	}
	
	return newGroup
}

func (g *componentGroup) AddGroup(group ComponentGroup) {
	newComponents := group.Components()
	newGroup := append(g.components, newComponents...)
	g.components = dedub(newGroup)
}

func (g *componentGroup) Components() []Component {
	return g.components
}

func (g *componentGroup) Each(f func(component Component)) {
	for _, component := range g.components {
		f(component)
	}
}

func (g *componentGroup) EachRecursive(f func(component Component)) {
	for _, component := range g.components {
		if container, ok := component.(Container); ok {
			container.Children().EachRecursive(f)
		}
		f(component)
	}
}

func (g *componentGroup) Add(c Component) {
	g.components = append(g.components, c)
}

func dedub(components []Component) []Component {
	keys := make(map[Component]bool)
	list := []Component{}
	for _, entry := range components {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
