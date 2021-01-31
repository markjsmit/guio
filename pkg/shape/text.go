package shape

import (
	"context"
	
	"github.com/mitchellh/mapstructure"
	"github.com/tfriedel6/canvas"
	
	"github.com/maxpower89/guio/pkg/theme"
)

func NewText(attributes map[string]string) (Component, error) {
	txt := Text{
		X:    0,
		Y:    0,
		Text: "",
		Fill: "#000",
	}
	mapstructure.WeakDecode(attributes, &txt)
	return &txt, nil
}

type Text struct {
	X, Y   float64
	Text   string
	Fill   string
	Family string
	Size   float64
}

func (t Text) Draw(ctx context.Context, renderFunc RenderFunc, canvas *canvas.Canvas, themeLoader theme.Loader) error {
	canvas.SetFillStyle(t.Fill)
	canvas.SetFont(themeLoader.GetPathFor("font/"+t.Family), t.Size)
	canvas.FillText(t.Text, t.X, t.Y)
	return nil
}
