package style

import (
	"context"
	"encoding/json"
	"strconv"
	
	"github.com/mitchellh/mapstructure"
	"github.com/veandco/go-sdl2/sdl"
	
	"github.com/maxpower89/guio/pkg/theme"
)

type state struct {
	abs   AbsoluteStyle
	style Style
	box   sdl.FRect
	hash  uint64
}

type Style struct {
	Left       string `json:"left"`
	Top        string `json:"top"`
	Width      string `json:"width"`
	Height     string `json:"height"`
	Padding    string `json:"padding"`
	Margin     string `json:"margin"`
	Background string `json:"background-color"`
	Border     string `json:"border-color"`
	Color      string `json:"color"`
	FontSize   string `json:"font-size"`
	FontFamily string `json:"font-family"`
	subStyles  []*Style
	loader     theme.Loader
	lastState  *state
}

type AbsoluteStyle struct {
	Box        sdl.FRect
	OuterBox   sdl.FRect
	InnerBox   sdl.FRect
	Margin     Directions
	Padding    Directions
	Background string
	Border     string
	Color      string
	FontSize   float64
	FontFamily string
	Ctx        context.Context
	OuterCtx   context.Context
	InnerCtx   context.Context
}

func FromCss(str []byte, loader theme.Loader) *Style {
	input := map[string]string{}
	json.Unmarshal(str, &input)
	return NewStyle(input, loader)
}

func NewStyle(attributes map[string]string, loader theme.Loader) *Style {
	style := &Style{
		subStyles: []*Style{},
		loader:    loader,
	}
	mapstructure.WeakDecode(attributes, style)
	return style
}

func (s *Style) Build(ctx context.Context) AbsoluteStyle {
	c := s.GetCompleteStyle()
	if c.Changed(ctx) {
		OuterBox := Rect(ctx, c.Left, c.Top, c.Width, c.Height)
		mt, mr, mb, ml := PaddingString(c.Margin)
		mctx := PadBoxContext(ctx, mt, mr, mb, ml)
		Box := Rect(mctx, c.Left, c.Top, c.Width, c.Height)
		pt, pr, pb, pl := PaddingString(c.Padding)
		pctx := PadBoxContext(mctx, pt, pr, pb, pl)
		InnerBox := Rect(pctx, c.Left, c.Top, c.Width, c.Height)
		
		margin := directionsFromPadValues(ctx, mt, mr, mb, ml)
		OuterBox.H += margin.Bottom
		OuterBox.W += margin.Right
		
		if s.lastState == nil {
			s.lastState = &state{}
		}
		
		fontSize,_:=strconv.ParseFloat(c.FontSize,64)
		
		s.lastState.abs = AbsoluteStyle{
			Box:        Box,
			OuterBox:   OuterBox,
			InnerBox:   InnerBox,
			Margin:     margin,
			Padding:    directionsFromPadValues(pctx, pt, pr, pb, pl),
			Background: c.Background,
			Border:     c.Border,
			Color:      c.Color,
			FontSize:   fontSize,
			FontFamily: c.FontFamily,
			Ctx:        mctx,
			OuterCtx:   ctx,
			InnerCtx:   pctx,
		}
		
		s.lastState.style = *s
		s.lastState.box = RectFromContext(ctx)
	}
	
	return s.lastState.abs
}

func (s *Style) Apply(substyle *Style) {
	s.subStyles = append(s.subStyles, substyle)
}

func (s *Style) Clear() {
	s.subStyles = []*Style{}
}

func (s Style) Changed(ctx context.Context) bool {
	
	return s.lastState == nil || !s.lastState.style.Equals(s) || s.lastState.box != RectFromContext(ctx)
	
}

func (a Style) Equals(b Style) bool {
	return a.Left == b.Left &&
		a.Top == b.Top &&
		a.Width == b.Width &&
		a.Height == b.Height &&
		a.Padding == b.Padding &&
		a.Margin == b.Margin &&
		a.Background == b.Background &&
		a.Border == b.Border &&
		a.Color == b.Color &&
		a.FontSize == b.FontSize &&
		a.FontFamily == b.FontFamily
}

func (s Style) GetCompleteStyle() Style {
	for _, sub := range s.subStyles {
		if s.Left == "" {
			s.Left = sub.Left
		}
		if s.Top == "" {
			s.Top = sub.Top
		}
		if s.Width == "" {
			s.Width = sub.Width
		}
		if s.Height == "" {
			s.Height = sub.Height
		}
		if s.Padding == "" {
			s.Padding = sub.Padding
		}
		if s.Margin == "" {
			s.Margin = sub.Margin
		}
		if s.Background == "" {
			s.Background = sub.Background
		}
		if s.Border == "" {
			s.Border = sub.Border
		}
		if s.Color == "" {
			s.Color = sub.Color
		}
		if s.FontSize == "" {
			s.FontSize = sub.FontSize
		}
		if s.FontFamily == "" {
			s.FontFamily = sub.FontFamily
		}
	}
	return s
}
