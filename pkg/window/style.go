package window

import (
	"github.com/maxpower89/guio/pkg/component"
	"github.com/maxpower89/guio/pkg/style"
)

func bindStyle(w *window) {
	styleData, _ := w.renderer.ThemeLoader().LoadStylesheet("style")
	stylesheet, _ := style.StylesheetFromJson(styleData, w.renderer.ThemeLoader())
	w.stylesheet = stylesheet
}

func updateStyling(w *window) {
	stylesheet := w.stylesheet
	w.componentGroup.EachRecursive(func(c component.Component) {
		if stylable, ok := c.(component.Stylable); ok {
			styler := stylable.Styler()
			styler.Clear()
		}
	})
	for _, element := range stylesheet.Elements {
		w.componentGroup.Select(element.Selector).Each(func(c component.Component) {
			if stylable, ok := c.(component.Stylable); ok {
				styler := stylable.Styler()
				styler.Apply(element.Style)
			}
		})
	}
}
