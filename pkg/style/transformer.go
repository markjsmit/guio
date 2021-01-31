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

func Rect(ctx context.Context, x string, y string, width string, height string) sdl.Rect {
	cx, _ := strconv.Atoi(Left(ctx, x))
	cy, _ := strconv.Atoi(Top(ctx, y))
	cw, _ := strconv.Atoi(Width(ctx, width))
	ch, _ := strconv.Atoi(Height(ctx, height))
	return sdl.Rect{
		X: int32(cx),
		Y: int32(cy),
		W: int32(cw),
		H: int32(ch),
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

var valueRegex = regexp.MustCompile(`(?m)([0-9]*)(.*)`)

func valueUnit(input string) (float64, string) {
	matches := valueRegex.FindStringSubmatch(input)
	var value float64 = 0
	unit := "px"
	if len(matches) >= 2 {
		value, _ = strconv.ParseFloat(matches[1],64)
	}
	if len(matches) >= 3 {
		unit = matches[2]
	}
	return value, unit
}
