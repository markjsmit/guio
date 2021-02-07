package component

import "strings"

type Identifier interface {
	HasIdentity(id string) bool
	ChangeState(state string)
}

type identifier struct {
	classes string
	id      string
	element string
	state   string
}

func (i *identifier) ChangeState(state string) {
	i.state = state
}

func NewIdentifier(class string, id string, element string, state string) *identifier {
	return &identifier{classes: " " + class + " ", id: id, element: element, state: state}
}

func (i identifier) HasIdentity(selector string) bool {
	
	if selector == "" {
		return false
	}
	
	buff := []rune{}
	for _, char := range selector {
		if char == '#' || char == '.' || char == ':' {
			if len(buff) > 0 && !i.singleIdentity(string(buff)) {
				return false
			}
			buff = []rune{}
		}
		buff = append(buff, char)
	}
	return i.singleIdentity(string(buff))
	
}

func (i identifier) singleIdentity(selector string) bool {
	switch selector[0] {
	case ':':
		return selector[1:] == i.state
	case '.':
		return strings.Contains(i.classes, " "+selector[1:]+" ")
	case '#':
		return selector[1:] == i.id
	default:
		return i.element == selector
	}
}
