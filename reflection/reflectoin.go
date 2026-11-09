package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	walkWithVisited(x, fn, map[uintptr]bool{})
}

func walkWithVisited(x interface{}, fn func(input string), visited map[uintptr]bool) {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		if val.IsNil() || visited[val.Pointer()] {
			return
		}
		visited[val.Pointer()] = true
		val = val.Elem()
	}

	walkValue := func(value reflect.Value) {
		walkWithVisited(value.Interface(), fn, visited)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		valFuncResult := val.Call(nil)

		for _, res := range valFuncResult {
			walkValue(res)
		}
	}
}
