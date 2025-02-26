package pkg

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Stringify(v interface{}) string {
	if v == nil {
		return "nil"
	}

	val := reflect.ValueOf(v)
	return stringifyValue(val)
}

func stringifyValue(val reflect.Value) string {
	switch val.Kind() {
	case reflect.Invalid:
		return "nil"
	case reflect.Ptr:
		if val.IsNil() {
			return fmt.Sprintf("*%s(nil)", val.Type().Elem().String())
		}
		return fmt.Sprintf("*%s(%s)", val.Type().Elem().String(), stringifyValue(val.Elem()))
	case reflect.Struct:
		return stringifyStruct(val)
	case reflect.Slice, reflect.Array:
		return stringifySliceOrArray(val)
	case reflect.Map:
		return stringifyMap(val)
	case reflect.String:
		return strconv.Quote(val.String())
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return fmt.Sprint(val.Interface())
	default:
		return fmt.Sprintf("%v", val.Interface())
	}
}

func stringifyStruct(val reflect.Value) string {
	var fields []string
	t := val.Type()
	for i := 0; i < val.NumField(); i++ {
		fieldName := t.Field(i).Name
		fieldValue := stringifyValue(val.Field(i))
		fields = append(fields, fmt.Sprintf("%s: %s", fieldName, fieldValue))
	}
	return fmt.Sprintf("struct{%s}", strings.Join(fields, ", "))
}

func stringifySliceOrArray(val reflect.Value) string {
	var elements []string
	for i := 0; i < val.Len(); i++ {
		elements = append(elements, stringifyValue(val.Index(i)))
	}
	typeStr := val.Type().String()
	return fmt.Sprintf("%s[%s]", typeStr, strings.Join(elements, ", "))
}

func stringifyMap(val reflect.Value) string {
	var pairs []string
	for _, key := range val.MapKeys() {
		value := val.MapIndex(key)
		pairs = append(pairs, fmt.Sprintf("%s: %s",
			stringifyValue(key),
			stringifyValue(value)))
	}

	return fmt.Sprintf("{%s}", strings.Join(pairs, ", "))
}
