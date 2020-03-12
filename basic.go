package tags

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

const (
	Default = "default"
	Min     = "min"
	Max     = "max"
)

var (
	tagHandlers = map[string]TagHandler{}
)

type TagHandler func(rawValue reflect.Value, tagValue string) (reflect.Value, error)

func init() {
	tagHandlers[Default] = getDefaultValue
	tagHandlers[Min] = getMinValue
	tagHandlers[Max] = getMaxValue
}

func Register(tagName string, h TagHandler) {
	tagHandlers[tagName] = h
}

func getMinValue(rawValue reflect.Value, tagValue string) (reflect.Value, error) {
	minValue, err := getValue(rawValue, tagValue)
	if err != nil {
		return rawValue, err
	}

	if valueCompare(rawValue, minValue) == -1 {
		return minValue, nil
	}
	return rawValue, nil
}

func getMaxValue(rawValue reflect.Value, tagValue string) (reflect.Value, error) {
	maxValue, err := getValue(rawValue, tagValue)
	if err != nil {
		return rawValue, err
	}
	if valueCompare(rawValue, maxValue) == 1 {
		return maxValue, nil
	}
	return rawValue, nil
}

func getDefaultValue(rawValue reflect.Value, tagValue string) (reflect.Value, error) {
	if !rawValue.IsZero() {
		return rawValue, nil
	}

	defaultValue, err := getValue(rawValue, tagValue)
	if err != nil {
		return rawValue, err
	}
	return defaultValue, nil
}

func valueCompare(v1 reflect.Value, v2 reflect.Value) int {
	switch v1.Type().String() {
	case "string":
		return strings.Compare(v1.String(), v2.String())
	case "int", "int16", "int32", "int64":
		i1, i2 := v1.Int(), v2.Int()
		if i1 > i2 {
			return 1
		} else if i1 < i2 {
			return -1
		}
		return 0
	case "uint", "uint16", "uint32", "uint64":
		i1, i2 := v1.Uint(), v2.Uint()
		if i1 > i2 {
			return 1
		} else if i1 < i2 {
			return -1
		}
		return 0
	case "float32", "float64":
		i1, i2 := v1.Float(), v2.Float()
		if i1 > i2 {
			return 1
		} else if i1 < i2 {
			return -1
		}
		return 0
	case "time.Duration":
		d1, d2 := v1.Interface().(time.Duration), v2.Interface().(time.Duration)
		if d1 > d2 {
			return 1
		} else if d1 < d2 {
			return -1
		}
		return 0
	case "bool":
		b1, b2 := v1.Bool(), v2.Bool()
		if b1 == b2 {
			return 0
		} else if b1 {
			return 1
		}
		return -1
	default:
		fmt.Printf("un-comparable type :%s, return equal", v1.Type().String())
	}
	return 0
}

func getValue(rawValue reflect.Value, tagValue string) (ret reflect.Value, err error) {
	ret = unsafeValueOf(rawValue) //reflect.Zero(rawValue.Type())
	switch rawValue.Type().String() {
	case "string":
		ret.SetString(tagValue)
		return
	case "int", "int16", "int32", "int64":
		var i int64
		i, err = strconv.ParseInt(tagValue, 10, 64)
		if err != nil {
			return
		}
		ret.SetInt(i)
	case "uint", "uint16", "uint32", "uint64":
		var i uint64
		i, err = strconv.ParseUint(tagValue, 10, 64)
		if err != nil {
			return
		}
		ret.SetUint(i)
	case "float32", "float64":
		var f float64
		f, err = strconv.ParseFloat(tagValue, 10)
		if err != nil {
			return
		}
		ret.SetFloat(f)
	case "bool":
		var b bool
		b, err = strconv.ParseBool(tagValue)
		if err != nil {
			return
		}
		ret.SetBool(b)
	case "time.Duration":
		var d time.Duration
		d, err = time.ParseDuration(tagValue)
		if err != nil {
			return
		}
		ret.Set(reflect.ValueOf(d))
	default:
		fmt.Printf("unsupported type:%T\n", rawValue)
		return rawValue, nil
	}
	return
}

func unsafeValueOf(val reflect.Value) reflect.Value {
	uptr := unsafe.Pointer(val.UnsafeAddr())
	return reflect.NewAt(val.Type(), uptr).Elem()
}
