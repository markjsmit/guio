package style

import (
	"context"
	"strconv"
	"strings"
)

const BoxX = "boxX"
const BoxY = "boxY"
const BoxW = "boxW"
const BoxH = "boxH"

func BoxContext(ctx context.Context, x string, y string, width string, height string) context.Context {
	bx, _ := strconv.Atoi(Left(ctx, x))
	by, _ := strconv.Atoi(Top(ctx, y))
	ctx = context.WithValue(ctx, BoxX, bx)
	ctx = context.WithValue(ctx, BoxY, by)
	bw, _ := strconv.Atoi(Width(ctx, width))
	bh, _ := strconv.Atoi(Height(ctx, height))
	ctx = context.WithValue(ctx, BoxW, bw)
	ctx = context.WithValue(ctx, BoxH, bh)
	return ctx
}

func PadBoxContext(ctx context.Context, top string, right string, bottom string, left string) context.Context {
	bx, _ := strconv.Atoi(Left(ctx, left))
	by, _ := strconv.Atoi(Top(ctx, top))
	
	ctx = context.WithValue(ctx, BoxX, bx)
	ctx = context.WithValue(ctx, BoxY, by)
	
	fullWidth, _ := strconv.Atoi(Width(ctx, "100%"))
	fullHeight, _ := strconv.Atoi(Height(ctx, "100%"))
	gapLeft, _ := strconv.Atoi(Width(ctx, left))
	gapTop, _ := strconv.Atoi(Height(ctx, top))
	gapRight, _ := strconv.Atoi(Width(ctx, right))
	gapBottom, _ := strconv.Atoi(Height(ctx, bottom))
	
	ctx = context.WithValue(ctx, BoxW, fullWidth-gapLeft-gapRight)
	ctx = context.WithValue(ctx, BoxH, fullHeight-gapTop-gapBottom)
	return ctx
}

func PaddingString(input string) (top string, right string, bottom string, left string) {
	split := strings.Split(input, " ")
	if len(split) == 4 {
		return split[0], split[1], split[2], split[3]
	} else if len(split) == 2 {
		return split[0], split[1], split[0], split[1]
	}
	return split[0], split[0], split[0], split[0]
}

func BoxFromContext(ctx context.Context) (x float64, y float64, w float64, h float64) {
	cx := ctx.Value(BoxX)
	cy := ctx.Value(BoxY)
	cw := ctx.Value(BoxW)
	ch := ctx.Value(BoxH)
	if cx != nil {
		x = float64(cx.(int))
	}
	if cy != nil {
		y = float64(cy.(int))
	}
	if cw != nil {
		w = float64(cw.(int))
	}
	if ch != nil {
		h = float64(ch.(int))
	}
	return
}

