package goredis

import (
	"strconv"

	"gopkg.in/guregu/null.v3"
)

// EmptyRedisReplyValue -
var EmptyRedisReplyValue = &RedisReplyValue{value: nil}

// RedisReplyValue -
type RedisReplyValue struct {
	value interface{}
}

// NewRedisReplyValue -
func NewRedisReplyValue(value interface{}) *RedisReplyValue {
	return &RedisReplyValue{
		value: value,
	}
}

// Value -
func (v *RedisReplyValue) Value() interface{} {
	return v.value
}

// AsInt32 -
func (v *RedisReplyValue) AsInt32(defaultValue int32) (int32, error) {
	if v.value != nil {
		switch v.value.(type) {
		case string:
			{
				r, err := strconv.ParseFloat(v.value.(string), 32)
				if err != nil {
					return defaultValue, err
				}
				return int32(r), nil
			}
		case int:
			return int32(v.value.(int)), nil
		case int32:
			return v.value.(int32), nil
		case int64:
			return int32(v.value.(int64)), nil
		}
	}
	return defaultValue, nil
}

// AsInt64 -
func (v *RedisReplyValue) AsInt64(defaultValue int64) (int64, error) {
	if v.value != nil {
		switch v.value.(type) {
		case string:
			{
				r, err := strconv.ParseFloat(v.value.(string), 64)
				if err != nil {
					return defaultValue, err
				}
				return int64(r), nil
			}
		case int:
			return int64(v.value.(int)), nil
		case int32:
			return int64(v.value.(int32)), nil
		case int64:
			return v.value.(int64), nil
		}
	}
	return defaultValue, nil
}

// AsFloat64 -
func (v *RedisReplyValue) AsFloat64(defaultValue float64) (float64, error) {
	if v.value != nil {
		switch v.value.(type) {
		case string:
			{
				r, err := strconv.ParseFloat(v.value.(string), 64)
				if err != nil {
					return defaultValue, err
				}
				return r, nil
			}
		case int:
			return float64(v.value.(int)), nil
		case int32:
			return float64(v.value.(int32)), nil
		case int64:
			return float64(v.value.(int64)), nil
		case float32:
			return float64(v.value.(float32)), nil
		case float64:
			return v.value.(float64), nil
		}
	}
	return defaultValue, nil
}

// AsString -
func (v *RedisReplyValue) AsString() string {
	if v.value != nil {
		switch v.value.(type) {
		case string:
			return v.value.(string)
		case int:
			return strconv.FormatInt(int64(v.value.(int)), 10)
		case int32:
			return strconv.FormatInt(int64(v.value.(int32)), 10)
		case int64:
			return strconv.FormatInt(v.value.(int64), 10)
		}
	}
	return ""
}

// IsNil -
func (v *RedisReplyValue) IsNil() bool {
	return v.value == nil
}

// ToArrayReplyReader -
func (v *RedisReplyValue) ToArrayReplyReader() *RedisArrayReplyReader {
	i, ok := v.value.([]interface{})
	if ok {
		return NewRedisArrayReplyReader(i)
	}
	return nil
}

// GetNullableInt -
func GetNullableInt(value *RedisReplyValue) (null.Int, error) {
	if value.value == nil {
		return null.IntFromPtr(nil), nil
	}
	var result int64
	var err error
	if result, err = value.AsInt64(0); err != nil {
		return null.IntFromPtr(nil), err
	}
	return null.IntFrom(result), nil
}

// GetNullableString -
func GetNullableString(value *RedisReplyValue) null.String {
	if value.value == nil {
		return null.StringFromPtr(nil)
	}
	return null.StringFrom(value.AsString())
}

// RedisArrayReplyReader -
type RedisArrayReplyReader struct {
	redisReply []interface{}
	position   uint32
}

// NewRedisArrayReplyReader -
func NewRedisArrayReplyReader(redisReply []interface{}) *RedisArrayReplyReader {
	return &RedisArrayReplyReader{
		redisReply: redisReply,
		position:   0,
	}
}

// GetLength -
func (r *RedisArrayReplyReader) GetLength() int {
	return len(r.redisReply)
}

// HasNext -
func (r *RedisArrayReplyReader) HasNext() bool {
	values := r.redisReply
	pos := r.position
	return pos < uint32(len(values))
}

// ReadArray -
func (r *RedisArrayReplyReader) ReadArray() *RedisArrayReplyReader {
	return NewRedisArrayReplyReader(r.ReadValue().value.([]interface{}))
}

// ReadString -
func (r *RedisArrayReplyReader) ReadString() string {
	return r.ReadValue().AsString()
}

// ReadInt32 -
func (r *RedisArrayReplyReader) ReadInt32(defaultValue int32) (int32, error) {
	return r.ReadValue().AsInt32(defaultValue)
}

// ReadInt64 -
func (r *RedisArrayReplyReader) ReadInt64(defaultValue int64) (int64, error) {
	return r.ReadValue().AsInt64(defaultValue)
}

// ReadFloat64 -
func (r *RedisArrayReplyReader) ReadFloat64(defaultValue float64) (float64, error) {
	return r.ReadValue().AsFloat64(defaultValue)
}

// SkipValue -
func (r *RedisArrayReplyReader) SkipValue() {
	r.ReadValue()
}

// ReadValue -
func (r *RedisArrayReplyReader) ReadValue() *RedisReplyValue {
	values := r.redisReply
	pos := r.position
	r.position++
	if pos < uint32(len(values)) {
		return &RedisReplyValue{value: values[pos]}
	}
	return EmptyRedisReplyValue
}

// ForEach -
func (r *RedisArrayReplyReader) ForEach(action func(i int, v *RedisReplyValue) error) error {
	for i, v := range r.redisReply {
		err := action(i, &RedisReplyValue{value: v})
		if err != nil {
			return err
		}
	}
	return nil
}
