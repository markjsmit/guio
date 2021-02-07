package main

import (
	"context"
	
	"github.com/maxpower89/guio/pkg/binding"
	"github.com/maxpower89/guio/pkg/component/widget/button"
	"github.com/maxpower89/guio/pkg/component/widget/textbox"
	"github.com/maxpower89/guio/pkg/event"
	appwnd "github.com/maxpower89/guio/pkg/window"
)

type handler struct {
	Button1 *button.Button `component:"button#button1"`
	Button2 *button.Button `component:"#button2"`
	Textbox *textbox.Textbox `component:"#text1"`
}

func (h handler) Init(c context.Context, w appwnd.Window) {
	
	count := 0
	h.Button1.Listen(event.MouseClick{}, func(event interface{}) {
		count++
	})
	w.Bind(binding.T(&count, &h.Button2.Text))
	w.Bind(binding.T(&count, &h.Textbox.Text))
	
}
