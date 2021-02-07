package style

import (
	"context"
	"strconv"
	"strings"
	
	"github.com/veandco/go-sdl2/sdl"
)

const BoxX = "boxX"
const BoxY = "boxY"
const BoxW = "boxW"
const BoxH = "boxH"

func BoxContext(ctx context.Context, x string, y string, width string, height string) context.Context {
	bx, _ := strconv.ParseFloat(Left(ctx, x), 64)
	by, _ := strconv.ParseFloat(Top(ctx, y), 64)
	ctx = context.WithValue(ctx, BoxX, bx)
	ctx = context.WithValue(ctx, BoxY, by)
	bw, _ := strconv.ParseFloat(Width(ctx, width), 64)
	bh, _ := strconv.ParseFloat(Height(ctx, height), 64)
	ctx = context.WithValue(ctx, BoxW, bw)
	ctx = context.WithValue(ctx, BoxH, bh)
	return ctx
}

func PadBoxContext(ctx context.Context, top string, right string, bottom string, left string) context.Context {
	bx, _ := strconv.ParseFloat(Left(ctx, left), 64)
	by, _ := strconv.ParseFloat(Top(ctx, top), 64)
	
	ctx = context.WithValue(ctx, BoxX, bx)
	ctx = context.WithValue(ctx, BoxY, by)
	
	fullWidth, _ := strconv.ParseFloat(Width(ctx, "100%"), 64)
	fullHeight, _ := strconv.ParseFloat(Height(ctx, "100%"), 64)
	gapLeft, _ := strconv.ParseFloat(Width(ctx, left), 64)
	gapTop, _ := strconv.ParseFloat(Height(ctx, top), 64)
	gapRight, _ := strconv.ParseFloat(Width(ctx, right), 64)
	gapBottom, _ := strconv.ParseFloat(Height(ctx, bottom), 64)
	
	ctx = context.WithValue(ctx, BoxW, fullWidth-gapLeft-gapRight)
	ctx = context.WithValue(ctx, BoxH, fullHeight-gapTop-gapBottom)
	return ctx
}

func PaddingString(input string) (top string, right string, bottom string, left string) {
	if input == "" {
		input = "0"
	}
	split := strings.Split(input, " ")
	if len(split) == 4 {
		return split[0], split[1], split[2], split[3]
	} else if len(split) == 2 {
		return split[0], split[1], split[0], split[1]
	}
	return split[0], split[0], split[0], split[0]
}

func RectFromContext(ctx context.Context) sdl.FRect {
	x, y, w, h := BoxFromContext(ctx)
	return sdl.FRect{
		X: float32(x),
		Y: float32(y),
		W: float32(w),
		H: float32(h),
	}
}

func BoxFromContext(ctx context.Context) (x float64, y float64, w float64, h float64) {
	cx := ctx.Value(BoxX)
	cy := ctx.Value(BoxY)
	cw := ctx.Value(BoxW)
	ch := ctx.Value(BoxH)
	if cx != nil {
		x = cx.(float64)
	}
	if cy != nil {
		y = cy.(float64)
	}
	if cw != nil {
		w = cw.(float64)
	}
	if ch != nil {
		h = ch.(float64)
	}
	return
}
