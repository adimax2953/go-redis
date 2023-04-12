package src

import (
	"encoding"
	"net"
	"reflect"
	"strings"
	"time"

	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// UpdateHashBatch function - keys, args[] string - return string , error
func (s *MyScriptor) UpdateHashBatch(keys []string, args ...interface{}) (*[]RedisResult, error) {

	values := make([]interface{}, len(args)-1)
	values = appendArgs(values, args)

	res, err := s.Scriptor.ExecSha(UpdateHashBatchID, keys, values)
	if err != nil {
		logtool.LogError("UpdateHashBatch ExecSha Error", err)
		return nil, err
	}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	count := len(res.([]interface{})) / 2
	result := make([]RedisResult, count)

	for i := 0; i < count; i++ {
		r := &RedisResult{}
		r.Key = reader.ReadString()
		r.Value = reader.ReadString()
		result[i] = *r
		if err != nil {
			logtool.LogError("UpdateHashBatch Value Error", err)
		}
	}

	return &result, nil
}

// UpdateHashBatch - 批量更新Hash
const (
	UpdateHashBatchID       = "UpdateHashBatch"
	UpdateHashBatchTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   UpdateHashBatch
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {v1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = KEYS[4]
		local v1                                            = ARGV
		local sender                                        = "UpdateHashBatch.lua"		

		if DBKey and ProjectKey and TagKey and k1 and v1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			redis.call('hset', MAIN_KEY , unpack(v1))

			local r1 = ""
			local Tmp = redis.call('hgetall',MAIN_KEY)
			if Tmp~=nil and Tmp~="" and Tmp~=false then
				r1 = Tmp
			end
			return v1
		end
    `
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
		logtool.LogDebug("UpdateHashMap map[string]interface{} ", dst)

		return dst
	case map[string]string:
		for k, v := range arg {
			dst = append(dst, k, v)
		}
		logtool.LogDebug("UpdateHashMap map[string]string ", dst)

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

// appendStructField appends the field and value held by the structure v to dst, and returns the appended dst.
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

		// miss field
		if omitEmpty(opt) && isEmptyValue(field) {
			continue
		}

		if field.CanInterface() {
			dst = append(dst, name, field.Interface())
		}
	}
	logtool.LogDebug("UpdateHashMap struct ", dst)

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
