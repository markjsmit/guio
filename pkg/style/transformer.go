package style

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	
	"github.com/veandco/go-sdl2/sdl"
)

const UnitPixel = "px"
const UnitPercent = "%"

func Left(ctx context.Context, value string) string {
	bx, _, bw, _ := BoxFromContext(ctx)
	
	v, u := valueUnit(value)
	if u == UnitPercent {
		v = (bw / 100 * v)
	}
	v += bx
	return fmt.Sprint(v)
}

func Top(ctx context.Context, value string) string {
	_, by, _, bh := BoxFromContext(ctx)
	v, u := valueUnit(value)
	if u == UnitPercent {
		v = bh / 100 * v
	}
	v += by
	return fmt.Sprint(v)
}

func Width(ctx context.Context, value string) string {
	_, _, bw, _ := BoxFromContext(ctx)
	v, u := valueUnit(value)
	if u == UnitPercent {
		v = (bw / 100 * v)
	}
	return fmt.Sprint(v)
}

func Height(ctx context.Context, value string) string {
	_, _, _, bh := BoxFromContext(ctx)
	v, u := valueUnit(value)
	if u == UnitPercent {
		v = bh / 100 * v
	}
	return fmt.Sprint(v)
	return value
}

func Rect(ctx context.Context, x string, y string, width string, height string) sdl.FRect {
	cx, _ := strconv.ParseFloat(Left(ctx, x),64)
	cy, _ := strconv.ParseFloat(Top(ctx, y),64)
	cw, _ := strconv.ParseFloat(Width(ctx, width),64)
	ch, _ := strconv.ParseFloat(Height(ctx, height),64)
	return sdl.FRect{
		X: float32(cx),
		Y: float32(cy),
		W: float32(cw),
		H: float32(ch),
	}
}


func Stroke(ctx context.Context, value string) string {
	return value
}

func Fill(ctx context.Context, value string) string {
	return value
}

func Bool(ctx context.Context, b bool) string {
	if b {
		return "1"
	}
	return ""
}

var valueRegex = regexp.MustCompile(`(?m)([0-9\.]*)(.*)`)

func valueUnit(input string) (float64, string) {
	matches := valueRegex.FindStringSubmatch(input)
	var value float64 = 0
	unit := "px"
	if len(matches) >= 2 {
		value, _ = strconv.ParseFloat(matches[1],64)
	}
	if len(matches) >= 3 {
		if matches[2]!="" {
			unit = matches[2]
		}
	}
	return value, unit
}
