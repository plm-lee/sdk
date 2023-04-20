package utils

import (
	"reflect"
	"strconv"
)

func ToString(v interface{}) string {
	if v == nil {
		return ""
	}

	t := reflect.TypeOf(v)
	switch t.Name() {
	case "int":
		return strconv.Itoa(v.(int))
	case "int64":
		return strconv.FormatInt(v.(int64), 10)
	case "float64":
		return strconv.FormatFloat(v.(float64), 'f', -1, 64)
	case "string":
		return v.(string)
	}
	return ""
}
