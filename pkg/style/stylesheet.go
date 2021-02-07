package style

import (
	"encoding/json"
	"sort"
	"strings"
	
	"github.com/maxpower89/guio/pkg/theme"
)

type StylesheetElement struct {
	Style    *Style
	Selector string
	Priority int
}
type Stylesheet struct {
	Elements []StylesheetElement
}

func StylesheetFromJson(input []byte, loader theme.Loader) (*Stylesheet, error) {
	inputMap := map[string]Style{}
	elements := []StylesheetElement{}
	err := json.Unmarshal(input, &inputMap)
	if err != nil {
		return nil, err
	}
	for selector, style := range inputMap {
		s:=style;
		elements = append(elements, StylesheetElement{
			Style:    &s,
			Selector: selector,
			Priority: calculatePriority(selector),
		})
	}
	sort.Slice(elements, func(i, j int) bool {
		return elements[i].Priority > elements[j].Priority
	})
	
	return &Stylesheet{Elements: elements}, nil
}

func calculatePriority(selector string) int {
	priority := 0
	priority += strings.Count(selector, " ")
	priority += strings.Count(selector, ".") * 100
	priority += strings.Count(selector, ":") * 10000
	priority += strings.Count(selector, "#") * 10000000
	return priority
}
