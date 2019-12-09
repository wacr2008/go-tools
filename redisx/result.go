package redisx

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
)

// 提供返回值的一些方法
type Result struct {
	val interface{}
	err error
}

func (r *Result) Value() (interface{}, error) {
	return r.val, wrapNil(r.err)
}

func (r *Result) String() (string, error) {
	if r.err != nil {
		return "", wrapNil(r.err)
	}
	switch val := r.val.(type) {
	case string:
		return val, nil
	case int64:
		return strconv.FormatInt(val, 10), nil
	default:
		return "", fmt.Errorf("redisx: unexpected type=%T for String", val)
	}
}

func (r *Result) Int64() (int64, error) {
	if r.err != nil {
		return 0, wrapNil(r.err)
	}
	switch val := r.val.(type) {
	case int64:
		return val, nil
	case string:
		return strconv.ParseInt(val, 10, 64)
	default:
		err := fmt.Errorf("redisx: unexpected type=%T for Int64", val)
		return 0, err
	}

}

func (r *Result) Float64() (float64, error) {
	if r.err != nil {
		return 0, wrapNil(r.err)
	}
	switch val := r.val.(type) {
	case int64:
		return float64(val), nil
	case string:
		return strconv.ParseFloat(val, 64)
	default:
		err := fmt.Errorf("redisx: unexpected type=%T for Float64", val)
		return 0, err
	}
}

func (r *Result) Bool() (bool, error) {
	if r.err != nil {
		return false, wrapNil(r.err)
	}
	switch val := r.val.(type) {
	case int64:
		return val != int64(0), nil
	case string:
		return strconv.ParseBool(val)
	default:
		err := fmt.Errorf("redisx: unexpected type=%T for Bool", val)
		return false, err
	}
}

func (r *Result) Bytes() ([]byte, error) {
	if r.err != nil {
		return nil, wrapNil(r.err)
	}
	switch val := r.val.(type) {
	case string:
		return []byte(val), nil
	default:
		err := fmt.Errorf("redisx: unexpected type=%T for Byte", val)
		return nil, err
	}
}

func (r *Result) Slice() ([]interface{}, error) {
	if r.err != nil {
		return nil, wrapNil(r.err)
	}
	switch val := r.val.(type) {
	case []interface{}:
		return val, nil
	default:
		err := fmt.Errorf("redisx: unexpected type=%T for []interface{}", val)
		return nil, err
	}
}

func wrapNil(err error) error {
	if err == redis.Nil {
		return nil
	}
	return err
}

func IsNil(err error) bool {
	return err == redis.Nil
}
