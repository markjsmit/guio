package textbox

import "github.com/maxpower89/guio/pkg/style"

type Tpl struct {
	Style      style.AbsoluteStyle
	Text       string
	CursorPos  int
	Active     bool
	BlinkState bool
}
