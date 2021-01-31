package main

import (
	"context"
	"fmt"
	
	"github.com/maxpower89/guio/pkg/app"
	"github.com/maxpower89/guio/pkg/component"
	"github.com/maxpower89/guio/pkg/event"
)

func main() {
	a := app.NewApp("./theme/gray", "./windows")
	w, _ := a.NewWindow(context.Background(), "main")
	w = w.WithHandler(&handler{})
	w.Open()
}

type handler struct {
	Button1 *component.Button `component:"button1"`
	Button2 *component.Button `component:"button2"`
}

func (h handler) Init(c context.Context) {
	h.Button1.Listen(event.MouseClick{}, func(event interface{}) {
		fmt.Println("button 1 clicked")
	})
	h.Button2.Listen(event.MouseClick{}, func(event interface{}) {
		fmt.Println("button 2 clicked")
	})
}
