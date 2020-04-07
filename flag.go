package tags

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

const (
	Flag = "flag"
)

var (
	fs = flag.CommandLine
)

func init() {
	tagHandlers[Flag] = setFlag
}

func setFlag(rawValue reflect.Value, tagValue string) error {
	if tagValue == "" {
		return nil
	}
	vv := strings.SplitN(tagValue, ",", 2)
	flagName := vv[0]
	hasDefaultValue := len(vv) == 2

	setFlagByTypeName := func(name string) (miss bool) {
		switch name {
		case "bool":
			var b bool
			if hasDefaultValue {
				b, _ = strconv.ParseBool(vv[1])
			}
			fs.BoolVar((*bool)(unsafe.Pointer(rawValue.UnsafeAddr())), flagName, b, tagValue)
		case "int":
			var i int64
			if hasDefaultValue {
				i, _ = strconv.ParseInt(vv[1], 0, 0)
			}
			fs.IntVar((*int)(unsafe.Pointer(rawValue.UnsafeAddr())), flagName, int(i), tagValue)
		case "uint":
			var i uint64
			if hasDefaultValue {
				i, _ = strconv.ParseUint(vv[1], 0, 0)
			}
			fs.UintVar((*uint)(unsafe.Pointer(rawValue.UnsafeAddr())), flagName, uint(i), tagValue)
		case "int64":
			var i int64
			if hasDefaultValue {
				i, _ = strconv.ParseInt(vv[1], 0, 0)
			}
			fs.Int64Var((*int64)(unsafe.Pointer(rawValue.UnsafeAddr())), flagName, i, tagValue)
		case "uint64":
			var i uint64
			if hasDefaultValue {
				i, _ = strconv.ParseUint(vv[1], 0, 0)
			}
			fs.Uint64Var((*uint64)(unsafe.Pointer(rawValue.UnsafeAddr())), flagName, i, tagValue)
		case "float64":
			var f float64
			if hasDefaultValue {
				f, _ = strconv.ParseFloat(vv[1], 64)
			}
			fs.Float64Var((*float64)(unsafe.Pointer(rawValue.UnsafeAddr())), flagName, f, tagValue)
		case "string":
			var s string
			if hasDefaultValue {
				s = vv[1]
			}
			fs.StringVar((*string)(unsafe.Pointer(rawValue.UnsafeAddr())), flagName, s, tagValue)
		case "time.Duration":
			var d time.Duration
			if hasDefaultValue {
				d, _ = time.ParseDuration(vv[1])
			}
			fs.DurationVar((*time.Duration)(unsafe.Pointer(rawValue.UnsafeAddr())), flagName, d, tagValue)
		default:
			miss = true
		}
		return
	}

	miss := setFlagByTypeName(rawValue.Type().String())
	if miss {
		miss = setFlagByTypeName(rawValue.Kind().String())
		if !miss {
			return fmt.Errorf("unsupport type :%s", rawValue.Type().String())
		}
	}

	return nil
}
