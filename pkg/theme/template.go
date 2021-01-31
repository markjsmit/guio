package theme

import (
	"fmt"
	"strconv"
	"text/template"
)

var funcMap = template.FuncMap{
	"add": func(args ...interface{}) float64 {
		var cur float64 = 0
		for _, arg := range args {
			val,_ := strconv.ParseFloat(fmt.Sprint(arg), 64)
			cur+= val;
		}
		
		return cur
	},
}
