package conversion

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type ConversionFunc func(a interface{}) interface{}

func GetConversionFunc(a interface{}, b interface{}) (atob ConversionFunc, btoa ConversionFunc, err error) {
	va := reflect.ValueOf(a)
	for va.Kind() == reflect.Interface {
		va = va.Elem()
	}
	vb := reflect.ValueOf(b)
	for vb.Kind() == reflect.Interface {
		vb = vb.Elem()
	}
	
	atobc, ok := conversionTable[va.Kind()][vb.Kind()]
	if !ok {
		return nil, nil, errors.New("No conversion found")
	}
	
	btoac, ok := conversionTable[vb.Kind()][va.Kind()]
	if !ok {
		return nil, nil, errors.New("No conversion found")
	}
	
	return atobc, btoac, nil
}

var conversionTable map[reflect.Kind]map[reflect.Kind]ConversionFunc = map[reflect.Kind]map[reflect.Kind]ConversionFunc{
	reflect.String: {
		reflect.String: func(a interface{}) interface{} {
			return a
		},
		reflect.Bool: func(a interface{}) interface{} {
			v, _ := strconv.ParseBool(a.(string))
			return v
		},
		reflect.Int: func(a interface{}) interface{} {
			v, _ := strconv.Atoi(a.(string))
			return v
		},
		reflect.Int64: func(a interface{}) interface{} {
			v, _ := strconv.ParseInt(a.(string), 10, 64)
			return v
		},
		reflect.Int32: func(a interface{}) interface{} {
			v, _ := strconv.ParseInt(a.(string), 10, 32)
			return v
		},
		reflect.Int16: func(a interface{}) interface{} {
			v, _ := strconv.ParseInt(a.(string), 10, 16)
			return v
		},
		reflect.Int8: func(a interface{}) interface{} {
			v, _ := strconv.ParseInt(a.(string), 10, 8)
			return v
		},
		reflect.Float32: func(a interface{}) interface{} {
			v, _ := strconv.ParseFloat(a.(string), 32)
			return v
		},
		reflect.Float64: func(a interface{}) interface{} {
			v, _ := strconv.ParseFloat(a.(string), 64)
			return v
		},
		reflect.Uint: func(a interface{}) interface{} {
			v, _ := strconv.ParseUint(a.(string), 10, 64)
			return v
		},
		reflect.Uint64: func(a interface{}) interface{} {
			v, _ := strconv.ParseUint(a.(string), 10, 64)
			return v
		},
		reflect.Uint32: func(a interface{}) interface{} {
			v, _ := strconv.ParseUint(a.(string), 10, 32)
			return v
		},
		reflect.Uint16: func(a interface{}) interface{} {
			v, _ := strconv.ParseUint(a.(string), 10, 16)
			return v
		},
		reflect.Uint8: func(a interface{}) interface{} {
			v, _ := strconv.ParseUint(a.(string), 10, 8)
			return v
		},
	},
	reflect.Int: {
		reflect.String: func(a interface{}) interface{} {
			return fmt.Sprint(a)
		},
		reflect.Bool: func(a interface{}) interface{} {
			return a.(int) != 0
		},
		reflect.Int: func(a interface{}) interface{} {
			return int(a.(int))
		},
		reflect.Int64: func(a interface{}) interface{} {
			return int64(a.(int))
		},
		reflect.Int32: func(a interface{}) interface{} {
			return int32(a.(int))
		},
		reflect.Int16: func(a interface{}) interface{} {
			return int16(a.(int))
		},
		reflect.Int8: func(a interface{}) interface{} {
			return int8(a.(int))
		},
		reflect.Float32: func(a interface{}) interface{} {
			return float32(a.(int))
		},
		reflect.Float64: func(a interface{}) interface{} {
			return float64(a.(int))
		},
		reflect.Uint: func(a interface{}) interface{} {
			return uint(a.(int))
		},
		reflect.Uint64: func(a interface{}) interface{} {
			return uint64(a.(int))
		},
		reflect.Uint32: func(a interface{}) interface{} {
			return uint32(a.(int))
		},
		reflect.Uint16: func(a interface{}) interface{} {
			return uint16(a.(int))
		},
		reflect.Uint8: func(a interface{}) interface{} {
			return uint8(a.(int))
		},
	},
	reflect.Int32: {
		reflect.String: func(a interface{}) interface{} {
			return fmt.Sprint(a)
		},
		reflect.Bool: func(a interface{}) interface{} {
			return a.(int32) != 0
		},
		reflect.Int: func(a interface{}) interface{} {
			return int(a.(int32))
		},
		reflect.Int64: func(a interface{}) interface{} {
			return int64(a.(int32))
		},
		reflect.Int32: func(a interface{}) interface{} {
			return int32(a.(int32))
		},
		reflect.Int16: func(a interface{}) interface{} {
			return int16(a.(int32))
		},
		reflect.Int8: func(a interface{}) interface{} {
			return int8(a.(int32))
		},
		reflect.Float32: func(a interface{}) interface{} {
			return float32(a.(int32))
		},
		reflect.Float64: func(a interface{}) interface{} {
			return float64(a.(int32))
		},
		reflect.Uint: func(a interface{}) interface{} {
			return uint(a.(int32))
		},
		reflect.Uint64: func(a interface{}) interface{} {
			return uint64(a.(int32))
		},
		reflect.Uint32: func(a interface{}) interface{} {
			return uint32(a.(int32))
		},
		reflect.Uint16: func(a interface{}) interface{} {
			return uint16(a.(int32))
		},
		reflect.Uint8: func(a interface{}) interface{} {
			return uint8(a.(int32))
		},
	},
	reflect.Int64: {
		reflect.String: func(a interface{}) interface{} {
			return fmt.Sprint(a)
		},
		reflect.Bool: func(a interface{}) interface{} {
			return a.(int64) != 0
		},
		reflect.Int: func(a interface{}) interface{} {
			return int(a.(int64))
		},
		reflect.Int64: func(a interface{}) interface{} {
			return int64(a.(int64))
		},
		reflect.Int32: func(a interface{}) interface{} {
			return int32(a.(int64))
		},
		reflect.Int16: func(a interface{}) interface{} {
			return int16(a.(int64))
		},
		reflect.Int8: func(a interface{}) interface{} {
			return int8(a.(int64))
		},
		reflect.Float32: func(a interface{}) interface{} {
			return float32(a.(int64))
		},
		reflect.Float64: func(a interface{}) interface{} {
			return float64(a.(int64))
		},
		reflect.Uint: func(a interface{}) interface{} {
			return uint(a.(int64))
		},
		reflect.Uint64: func(a interface{}) interface{} {
			return uint64(a.(int64))
		},
		reflect.Uint32: func(a interface{}) interface{} {
			return uint32(a.(int64))
		},
		reflect.Uint16: func(a interface{}) interface{} {
			return uint16(a.(int64))
		},
		reflect.Uint8: func(a interface{}) interface{} {
			return uint8(a.(int64))
		},
	},
	reflect.Int16: {
		reflect.String: func(a interface{}) interface{} {
			return fmt.Sprint(a)
		},
		reflect.Bool: func(a interface{}) interface{} {
			return a.(int16) != 0
		},
		reflect.Int: func(a interface{}) interface{} {
			return int(a.(int16))
		},
		reflect.Int64: func(a interface{}) interface{} {
			return int64(a.(int16))
		},
		reflect.Int32: func(a interface{}) interface{} {
			return int32(a.(int16))
		},
		reflect.Int16: func(a interface{}) interface{} {
			return int16(a.(int16))
		},
		reflect.Int8: func(a interface{}) interface{} {
			return int8(a.(int16))
		},
		reflect.Float32: func(a interface{}) interface{} {
			return float32(a.(int16))
		},
		reflect.Float64: func(a interface{}) interface{} {
			return float64(a.(int16))
		},
		reflect.Uint: func(a interface{}) interface{} {
			return uint(a.(int16))
		},
		reflect.Uint64: func(a interface{}) interface{} {
			return uint64(a.(int16))
		},
		reflect.Uint32: func(a interface{}) interface{} {
			return uint32(a.(int16))
		},
		reflect.Uint16: func(a interface{}) interface{} {
			return uint16(a.(int16))
		},
		reflect.Uint8: func(a interface{}) interface{} {
			return uint8(a.(int16))
		},
	},
	reflect.Int8: {
		reflect.String: func(a interface{}) interface{} {
			return fmt.Sprint(a)
		},
		reflect.Bool: func(a interface{}) interface{} {
			return a.(int8) != 0
		},
		reflect.Int: func(a interface{}) interface{} {
			return int(a.(int8))
		},
		reflect.Int64: func(a interface{}) interface{} {
			return int64(a.(int8))
		},
		reflect.Int32: func(a interface{}) interface{} {
			return int32(a.(int8))
		},
		reflect.Int16: func(a interface{}) interface{} {
			return int16(a.(int8))
		},
		reflect.Int8: func(a interface{}) interface{} {
			return int8(a.(int8))
		},
		reflect.Float32: func(a interface{}) interface{} {
			return float32(a.(int8))
		},
		reflect.Float64: func(a interface{}) interface{} {
			return float64(a.(int8))
		},
		reflect.Uint: func(a interface{}) interface{} {
			return uint(a.(int8))
		},
		reflect.Uint64: func(a interface{}) interface{} {
			return uint64(a.(int8))
		},
		reflect.Uint32: func(a interface{}) interface{} {
			return uint32(a.(int8))
		},
		reflect.Uint16: func(a interface{}) interface{} {
			return uint16(a.(int8))
		},
		reflect.Uint8: func(a interface{}) interface{} {
			return uint8(a.(int8))
		},
	},
	
	reflect.Uint: {
		reflect.String: func(a interface{}) interface{} {
			return fmt.Sprint(a)
		},
		reflect.Bool: func(a interface{}) interface{} {
			return a.(uint) != 0
		},
		reflect.Int: func(a interface{}) interface{} {
			return int(a.(uint))
		},
		reflect.Int64: func(a interface{}) interface{} {
			return int64(a.(uint))
		},
		reflect.Int32: func(a interface{}) interface{} {
			return int32(a.(uint))
		},
		reflect.Int16: func(a interface{}) interface{} {
			return int16(a.(uint))
		},
		reflect.Int8: func(a interface{}) interface{} {
			return int8(a.(uint))
		},
		reflect.Float32: func(a interface{}) interface{} {
			return float32(a.(uint))
		},
		reflect.Float64: func(a interface{}) interface{} {
			return float64(a.(uint))
		},
		reflect.Uint: func(a interface{}) interface{} {
			return uint(a.(uint))
		},
		reflect.Uint64: func(a interface{}) interface{} {
			return uint64(a.(uint))
		},
		reflect.Uint32: func(a interface{}) interface{} {
			return uint32(a.(uint))
		},
		reflect.Uint16: func(a interface{}) interface{} {
			return uint16(a.(uint))
		},
		reflect.Uint8: func(a interface{}) interface{} {
			return uint8(a.(uint))
		},
	},
	reflect.Uint32: {
		reflect.String: func(a interface{}) interface{} {
			return fmt.Sprint(a)
		},
		reflect.Bool: func(a interface{}) interface{} {
			return a.(uint32) != 0
		},
		reflect.Int: func(a interface{}) interface{} {
			return int(a.(uint32))
		},
		reflect.Int64: func(a interface{}) interface{} {
			return int64(a.(uint32))
		},
		reflect.Int32: func(a interface{}) interface{} {
			return int32(a.(uint32))
		},
		reflect.Int16: func(a interface{}) interface{} {
			return int16(a.(uint32))
		},
		reflect.Int8: func(a interface{}) interface{} {
			return int8(a.(uint32))
		},
		reflect.Float32: func(a interface{}) interface{} {
			return float32(a.(uint32))
		},
		reflect.Float64: func(a interface{}) interface{} {
			return float64(a.(uint32))
		},
		reflect.Uint: func(a interface{}) interface{} {
			return uint(a.(uint32))
		},
		reflect.Uint64: func(a interface{}) interface{} {
			return uint64(a.(uint32))
		},
		reflect.Uint32: func(a interface{}) interface{} {
			return uint32(a.(uint32))
		},
		reflect.Uint16: func(a interface{}) interface{} {
			return uint16(a.(uint32))
		},
		reflect.Uint8: func(a interface{}) interface{} {
			return uint8(a.(uint32))
		},
	},
	reflect.Uint64: {
		reflect.String: func(a interface{}) interface{} {
			return fmt.Sprint(a)
		},
		reflect.Bool: func(a interface{}) interface{} {
			return a.(uint64) != 0
		},
		reflect.Int: func(a interface{}) interface{} {
			return int(a.(uint64))
		},
		reflect.Int64: func(a interface{}) interface{} {
			return int64(a.(uint64))
		},
		reflect.Int32: func(a interface{}) interface{} {
			return int32(a.(uint64))
		},
		reflect.Int16: func(a interface{}) interface{} {
			return int16(a.(uint64))
		},
		reflect.Int8: func(a interface{}) interface{} {
			return int8(a.(uint64))
		},
		reflect.Float32: func(a interface{}) interface{} {
			return float32(a.(uint64))
		},
		reflect.Float64: func(a interface{}) interface{} {
			return float64(a.(uint64))
		},
		reflect.Uint: func(a interface{}) interface{} {
			return uint(a.(uint64))
		},
		reflect.Uint64: func(a interface{}) interface{} {
			return uint64(a.(uint64))
		},
		reflect.Uint32: func(a interface{}) interface{} {
			return uint32(a.(uint64))
		},
		reflect.Uint16: func(a interface{}) interface{} {
			return uint16(a.(uint64))
		},
		reflect.Uint8: func(a interface{}) interface{} {
			return uint8(a.(uint64))
		},
	},
	reflect.Uint16: {
		reflect.String: func(a interface{}) interface{} {
			return fmt.Sprint(a)
		},
		reflect.Bool: func(a interface{}) interface{} {
			return a.(uint16) != 0
		},
		reflect.Int: func(a interface{}) interface{} {
			return int(a.(uint16))
		},
		reflect.Int64: func(a interface{}) interface{} {
			return int64(a.(uint16))
		},
		reflect.Int32: func(a interface{}) interface{} {
			return int32(a.(uint16))
		},
		reflect.Int16: func(a interface{}) interface{} {
			return int16(a.(uint16))
		},
		reflect.Int8: func(a interface{}) interface{} {
			return int8(a.(uint16))
		},
		reflect.Float32: func(a interface{}) interface{} {
			return float32(a.(uint16))
		},
		reflect.Float64: func(a interface{}) interface{} {
			return float64(a.(uint16))
		},
		reflect.Uint: func(a interface{}) interface{} {
			return uint(a.(uint16))
		},
		reflect.Uint64: func(a interface{}) interface{} {
			return uint64(a.(uint16))
		},
		reflect.Uint32: func(a interface{}) interface{} {
			return uint32(a.(uint16))
		},
		reflect.Uint16: func(a interface{}) interface{} {
			return uint16(a.(uint16))
		},
		reflect.Uint8: func(a interface{}) interface{} {
			return uint8(a.(uint16))
		},
	},
	reflect.Uint8: {
		reflect.String: func(a interface{}) interface{} {
			return fmt.Sprint(a)
		},
		reflect.Bool: func(a interface{}) interface{} {
			return a.(uint8) != 0
		},
		reflect.Int: func(a interface{}) interface{} {
			return int(a.(uint8))
		},
		reflect.Int64: func(a interface{}) interface{} {
			return int64(a.(uint8))
		},
		reflect.Int32: func(a interface{}) interface{} {
			return int32(a.(uint8))
		},
		reflect.Int16: func(a interface{}) interface{} {
			return int16(a.(uint8))
		},
		reflect.Int8: func(a interface{}) interface{} {
			return int8(a.(uint8))
		},
		reflect.Float32: func(a interface{}) interface{} {
			return float32(a.(uint8))
		},
		reflect.Float64: func(a interface{}) interface{} {
			return float64(a.(uint8))
		},
		reflect.Uint: func(a interface{}) interface{} {
			return uint(a.(uint8))
		},
		reflect.Uint64: func(a interface{}) interface{} {
			return uint64(a.(uint8))
		},
		reflect.Uint32: func(a interface{}) interface{} {
			return uint32(a.(uint8))
		},
		reflect.Uint16: func(a interface{}) interface{} {
			return uint16(a.(uint8))
		},
		reflect.Uint8: func(a interface{}) interface{} {
			return uint8(a.(uint8))
		},
	},
	
	reflect.Float64: {
		reflect.String: func(a interface{}) interface{} {
			return fmt.Sprint(a)
		},
		reflect.Bool: func(a interface{}) interface{} {
			return a.(float64) != 0
		},
		reflect.Int: func(a interface{}) interface{} {
			return int(a.(float64))
		},
		reflect.Int64: func(a interface{}) interface{} {
			return int64(a.(float64))
		},
		reflect.Int32: func(a interface{}) interface{} {
			return int32(a.(float64))
		},
		reflect.Int16: func(a interface{}) interface{} {
			return int16(a.(float64))
		},
		reflect.Int8: func(a interface{}) interface{} {
			return int8(a.(float64))
		},
		reflect.Float32: func(a interface{}) interface{} {
			return float32(a.(float64))
		},
		reflect.Float64: func(a interface{}) interface{} {
			return float64(a.(float64))
		},
		reflect.Uint: func(a interface{}) interface{} {
			return uint(a.(float64))
		},
		reflect.Uint64: func(a interface{}) interface{} {
			return uint64(a.(float64))
		},
		reflect.Uint32: func(a interface{}) interface{} {
			return uint32(a.(float64))
		},
		reflect.Uint16: func(a interface{}) interface{} {
			return uint16(a.(float64))
		},
		reflect.Uint8: func(a interface{}) interface{} {
			return uint8(a.(float64))
		},
	},
	
	reflect.Float32: {
		reflect.String: func(a interface{}) interface{} {
			return fmt.Sprint(a)
		},
		reflect.Bool: func(a interface{}) interface{} {
			return a.(float32) != 0
		},
		reflect.Int: func(a interface{}) interface{} {
			return int(a.(float32))
		},
		reflect.Int64: func(a interface{}) interface{} {
			return int64(a.(float32))
		},
		reflect.Int32: func(a interface{}) interface{} {
			return int32(a.(float32))
		},
		reflect.Int16: func(a interface{}) interface{} {
			return int16(a.(float32))
		},
		reflect.Int8: func(a interface{}) interface{} {
			return int8(a.(float32))
		},
		reflect.Float32: func(a interface{}) interface{} {
			return float32(a.(float32))
		},
		reflect.Float64: func(a interface{}) interface{} {
			return float64(a.(float32))
		},
		reflect.Uint: func(a interface{}) interface{} {
			return uint(a.(float32))
		},
		reflect.Uint64: func(a interface{}) interface{} {
			return uint64(a.(float32))
		},
		reflect.Uint32: func(a interface{}) interface{} {
			return uint32(a.(float32))
		},
		reflect.Uint16: func(a interface{}) interface{} {
			return uint16(a.(float32))
		},
		reflect.Uint8: func(a interface{}) interface{} {
			return uint8(a.(float32))
		},
	},
	
	reflect.Bool: {
		reflect.String: func(a interface{}) interface{} {
			if a.(bool) {
				return "true"
			}
			return "false"
		},
		reflect.Bool: func(a interface{}) interface{} {
			return a.(bool)
		},
		reflect.Int: func(a interface{}) interface{} {
			if a.(bool) {
				return int(1)
			}
			return int(0)
		},
		reflect.Int64: func(a interface{}) interface{} {
			if a.(bool) {
				return int64(1)
			}
			return int64(0)
		},
		reflect.Int32: func(a interface{}) interface{} {
			if a.(bool) {
				return int32(1)
			}
			return int32(0)
		},
		reflect.Int16: func(a interface{}) interface{} {
			if a.(bool) {
				return int16(1)
			}
			return int16(0)
		},
		reflect.Int8: func(a interface{}) interface{} {
			if a.(bool) {
				return int8(1)
			}
			return int8(0)
		},
		reflect.Float32: func(a interface{}) interface{} {
			if a.(bool) {
				return float32(1)
			}
			return float32(0)
		},
		reflect.Float64: func(a interface{}) interface{} {
			if a.(bool) {
				return float64(1)
			}
			return float64(0)
		},
		reflect.Uint: func(a interface{}) interface{} {
			if a.(bool) {
				return uint(1)
			}
			return uint(0)
		},
		reflect.Uint64: func(a interface{}) interface{} {
			if a.(bool) {
				return uint64(1)
			}
			return uint64(0)
		},
		reflect.Uint32: func(a interface{}) interface{} {
			if a.(bool) {
				return uint32(1)
			}
			return uint32(0)
		},
		reflect.Uint16: func(a interface{}) interface{} {
			if a.(bool) {
				return uint16(1)
			}
			return uint16(0)
		},
		reflect.Uint8: func(a interface{}) interface{} {
			if a.(bool) {
				return uint8(1)
			}
			return uint8(0)
		},
	},
}
