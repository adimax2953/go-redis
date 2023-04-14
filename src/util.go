package src

import (
	"encoding"
	"net"
	"reflect"
	"strings"
	"time"
)

func appendArgs(dst, args []interface{}) []interface{} {
	if len(args) == 1 {
		return appendArg(dst, args[0])
	}

	dst = append(dst, args...)
	return dst
}

func appendArg(dst []interface{}, arg interface{}) []interface{} {
	switch arg := arg.(type) {
	case []string:
		for _, s := range arg {
			dst = append(dst, s)
		}
		return dst
	case []interface{}:
		dst = append(dst, arg...)
		return dst
	case map[string]interface{}:
		for k, v := range arg {
			dst = append(dst, k, v)
		}

		return dst
	case map[string]string:
		for k, v := range arg {
			dst = append(dst, k, v)
		}

		return dst
	case time.Time, time.Duration, encoding.BinaryMarshaler, net.IP:
		return append(dst, arg)
	default:
		// scan struct field
		v := reflect.ValueOf(arg)
		if v.Type().Kind() == reflect.Ptr {
			if v.IsNil() {
				// error: arg is not a valid object
				return dst
			}
			v = v.Elem()
		}

		if v.Type().Kind() == reflect.Struct {
			return appendStructField(dst, v)
		}

		return append(dst, arg)
	}
}

// appendStructField 將結構體 v 持有的字段和值附加到 dst，並返回附加的 dst。
func appendStructField(dst []interface{}, v reflect.Value) []interface{} {
	typ := v.Type()
	for i := 0; i < typ.NumField(); i++ {
		tag := typ.Field(i).Tag.Get("redis")
		if tag == "" || tag == "-" {
			continue
		}
		name, opt, _ := strings.Cut(tag, ",")
		if name == "" {
			continue
		}

		field := v.Field(i)
		// 缺失字段
		if omitEmpty(opt) && isEmptyValue(field) {
			continue
		}

		if field.CanInterface() {
			dst = append(dst, name, field.Interface())
		}
	}

	return dst
}

func omitEmpty(opt string) bool {
	for opt != "" {
		var name string
		name, opt, _ = strings.Cut(opt, ",")
		if name == "omitempty" {
			return true
		}
	}
	return false
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Pointer:
		return v.IsNil()
	}
	return false
}
