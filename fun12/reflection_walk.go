package fun12

import "reflect"

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func oldWalk(x interface{}, fn func(input string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			oldWalk(val.Field(i).Interface(), fn)
		}
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			oldWalk(val.Index(i).Interface(), fn)
		}
	case reflect.String:
		fn(val.String())
	}
}

// walk is little wonky because the abstraction of for loop does not cover the whole
// switch statement. we are re-implementing the for loop for reflect.Map
func wonkyWalk(x interface{}, fn func(input string)) {
	val := getValue(x)

	numberOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getField = val.Index
	case reflect.Map:
		for _, key := range val.MapKeys() {
			wonkyWalk(val.MapIndex(key).Interface(), fn)
		}
	}

	for i := 0; i < numberOfValues; i++ {
		wonkyWalk(getField(i).Interface(), fn)
	}
}

/*
How do you feel right now? It felt like maybe a nice abstraction at the time but now the code feels a little wonky.

This is OK! Refactoring is a journey and sometimes we will make mistakes. A major point of TDD is it gives us the freedom to try these things out.

By taking small steps backed by tests this is in no way an irreversible situation. Let's just put it back to how it was before the refactor.
*/
func walk(x interface{}, fn func(input string)) {
	val := getValue(x)
	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
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
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}

	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
