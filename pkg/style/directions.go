package style

import (
	"context"
	"strconv"
)


type Directions struct {
	Top    float32
	Right  float32
	Bottom float32
	Left   float32
}

func directionsFromPadValues(ctx context.Context, pt string, pr string, pb string, pl string) Directions {
	gl, _ := strconv.ParseFloat(Width(ctx, pl),64)
	gt, _ := strconv.ParseFloat(Height(ctx, pt),64)
	gr, _ := strconv.ParseFloat(Width(ctx, pr),64)
	gb, _ := strconv.ParseFloat(Height(ctx, pb),64)
	return Directions{
		Top:    float32(gt),
		Right:   float32(gr),
		Bottom:  float32(gb),
		Left:    float32(gl),
	}
}

