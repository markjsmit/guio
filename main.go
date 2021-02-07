package main

import (
	"context"
	
	"github.com/maxpower89/guio/pkg/app"
)

func main() {
	a := app.NewApp("./theme/gray", "./windows")
	w, _ := a.NewWindow(context.Background(), "main")
	w = w.WithHandler(&handler{})
	w.Open()
}

