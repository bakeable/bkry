package utils

import (
	"reflect"
	"strconv"
	"time"
)

func TryDecodeStringArray(m interface{}) []string {
	if m == nil {
		return nil
	}
	
	var s []string
	for _, v := range m.([]interface{}) {
		s = append(s, v.(string))
	}
	return s
}

func TryGetFirst(m interface{}) interface{} {
	if m == nil {
		return nil
	}
	if len(m.([]interface{})) == 0 {
		return nil
	}
	return m.([]interface{})[0]
}

func TryDecodeMap(m interface{}) map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.(map[string]interface{})
}

func TryDecodeBool(m interface{}, fallback bool) bool {
	if m == nil {
		return fallback
	}
	return m.(bool)
}

func TryDecodeString(m interface{}) string {
	if m == nil {
		return ""
	}
	if v, ok := m.(string); ok {
		return v
	}
	
	if v, ok := m.(int); ok {
		return strconv.Itoa(v)
	}
	if v, ok := m.(int64); ok {
		return strconv.Itoa(int(v))
	}

	return ""
}

func TryDecodeInt(m interface{}) int {
	if m == nil {
		return 0
	}
	return int(m.(int64))
}

func TryDecodeFloat64(m interface{}) float64 {
	if m == nil {
		return 0
	}
	return m.(float64)
}

func TryDecodeTimestamp(m interface{}) string {
	if m == nil {
		return ""
	}
	str, ok := m.(time.Time); if ok {
		return str.Format(time.RFC3339)
	}
	return ""
}

func DecodeMap(m interface{}, defaultValue map[string]interface{}) map[string]interface{} {
	if m == nil {
		return defaultValue
	}
	
	if v, ok := m.(map[string]interface{}); ok {
		return v
	}

	return defaultValue
}

func DecodeString(m interface{}, defaultValue string) string {
	return DefaultDecode(m, defaultValue).(string)
}

func DecodeInt(m interface{}, defaultValue int) int {
	return DefaultDecode(m, defaultValue).(int)
}

func DecodeInt64(m interface{}, defaultValue int64) int64 {
	return DefaultDecode(m, defaultValue).(int64)
}

func DecodeBool(m interface{}, defaultValue bool) bool {
	return DefaultDecode(m, defaultValue).(bool)
}

func DecodeFloat64(m interface{}, defaultValue float64) float64 {
	return DefaultDecode(m, defaultValue).(float64)
}

func DecodeTime(m interface{}) time.Time {
	return DefaultDecode(m, time.Now()).(time.Time)
}
func DecodeStringArray(m interface{}, defaultValue []string) []string {
	if m == nil {
		return defaultValue
	}

	// Check if the input is a slice
	if v, ok := m.([]interface{}); ok {
		var s []string
		for _, i := range v {
			x := DefaultDecode(i, "").(string)
			if x != "" {
				s = append(s, x)
			}
		}
		return s
	}

	// Check if the input is a string
	if v, ok := m.(string); ok {
		return []string{v}
	}

	return defaultValue
}

func DecodeIntArray(m interface{}, defaultValue []int) []int {
	if m == nil {
		return defaultValue
	}

	// Check if the input is a slice
	if v, ok := m.([]interface{}); ok {
		var s []int
		for _, i := range v {
			x := DefaultDecode(i, 0).(int)
			s = append(s, x)
		}
		return s
	}

	// Check if the input is a string
	if v, ok := m.(int); ok {
		return []int{v}
	}

	return defaultValue
}


func DecodeIntArrayArray(m interface{}, defaultValue [][]int) [][]int {
	if m == nil {
		return defaultValue
	}

	// Check if the input is a slice
	if v, ok := m.([]interface{}); ok {
		var s [][]int
		for i, a := range v {
			if i >= len(defaultValue) {
				x := DecodeIntArray(a, defaultValue[i])
				s = append(s, x)
			} else {
				x := DecodeIntArray(a, []int{})
				s = append(s, x)
			}
		}
		return s
	}

	return defaultValue
}

func DecodeFloatArray(m interface{}, defaultValue []float64) []float64 {
	if m == nil {
		return defaultValue
	}

	// Check if the input is a slice
	if v, ok := m.([]interface{}); ok {
		var s []float64
		for _, i := range v {
			x := DefaultDecode(i, 0.0).(float64)
			s = append(s, x)
		}
		return s
	}

	// Check if the input is a float
	if v, ok := m.(float64); ok {
		return []float64{v}
	}

	return defaultValue
}

func DecodeInterfaceArray(m interface{}, defaultValue []interface{}) []interface{} {
	if m == nil {
		return defaultValue
	}

	// Check if the input is a slice
	if v, ok := m.([]interface{}); ok {
		return v
	}

	return defaultValue
}

func DecodeIntMap(m interface{}, defaultValue map[string]int) map[string]int {
	return DefaultDecode(m, defaultValue).(map[string]int)
}

func DecodeFloatMap(m interface{}, defaultValue map[string]float64) map[string]float64 {
	return DefaultDecode(m, defaultValue).(map[string]float64)
}

func DecodeBoolMap(m interface{}, defaultValue map[string]bool) map[string]bool {
	return DefaultDecode(m, defaultValue).(map[string]bool)
}

func DecodeStringMap(m interface{}, defaultValue map[string]string) map[string]string {
	return DefaultDecode(m, defaultValue).(map[string]string)
}

func DecodeIntIntMap(m interface{}, defaultValue map[int]int) map[int]int {
	if m == nil {
		return defaultValue
	}

	// Check if the input is a map
	if v, ok := m.(map[int]int); ok {
		return v
	}

	return defaultValue
}

// DefaultDecode decodes the input to the specified type with a default value
func DefaultDecode(m interface{}, defaultValue interface{}) interface{} {
	if m == nil {
		return defaultValue
	}

	// Get the type of the default value
	defaultType := reflect.TypeOf(defaultValue)

	// Use reflection to check the type of the input
	v := reflect.ValueOf(m)
	switch defaultType.Kind() {
	case reflect.Slice:
		// Check if the slice element type is compatible
		if v.Kind() == reflect.Slice && v.Type().Elem().Kind() == defaultType.Elem().Kind() {
			return v.Interface()
		}
	case reflect.Map:
		// Check if the map key and value types are compatible
		if v.Kind() == reflect.Map && v.Type().Key().Kind() == defaultType.Key().Kind() && v.Type().Elem().Kind() == defaultType.Elem().Kind() {
			return v.Interface()
		}
	case reflect.Int:
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return int(v.Int())
		case reflect.Float32, reflect.Float64:
			return int(v.Float())
		case reflect.String:
			i, err := strconv.Atoi(v.String())
			if err != nil {
				return defaultValue
			}
			return int(i)
		case reflect.Bool:
			if v.Bool() {
				return 1
			}
			return 0
		case reflect.Struct:
			if v.Type() == reflect.TypeOf(time.Time{}) {
				return int(v.Interface().(time.Time).Unix())
			} else if v.Type() == reflect.TypeOf(time.Duration(0)) {
				return int(v.Interface().(time.Duration))
			}
		}
	case reflect.String:
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return strconv.FormatInt(v.Int(), 10)
		case reflect.Float32, reflect.Float64:
			return strconv.FormatFloat(v.Float(), 'f', -1, 64)
		case reflect.String:
			return v.String()
		case reflect.Bool:
			if v.Bool() {
				return "true"
			}
			return "false"
		case reflect.Struct:
			if v.Type() == reflect.TypeOf(time.Time{}) {
				return v.Interface().(time.Time).Format(time.RFC3339)
			} else if v.Type() == reflect.TypeOf(time.Duration(0)) {
				return v.Interface().(time.Duration).String()
			}
		}
	case reflect.Bool:
		switch v.Kind() {
		case reflect.Bool:
			return v.Bool()
		case reflect.String:
			b, err := strconv.ParseBool(v.String())
			if err != nil {
				return defaultValue
			}
			return b
		}
	case reflect.Float64:
		switch v.Kind() {
		case reflect.Float32, reflect.Float64:
			return v.Float()
		case reflect.String:
			f, err := strconv.ParseFloat(v.String(), 64)
			if err != nil {
				return defaultValue
			}
			return f
		}
	case reflect.Struct:
		switch v.Kind() {
		case reflect.Struct:
			if v.Type() == reflect.TypeOf(time.Time{}) {
				return v.Interface().(time.Time)
			} else if v.Type() == reflect.TypeOf(time.Duration(0)) {
				return v.Interface().(time.Duration)
			}
		}
	}
	return defaultValue
}