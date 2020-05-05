package algorithm

import (
	"reflect"
)

// AllOf checks that all elements in slice satisfy the condition.
func AllOf(slice interface{}, predicate func(i int) bool) bool {
	rv := reflect.ValueOf(slice)
	length := rv.Len()
	for i := 0; i < length; i++ {
		if !predicate(i) {
			return false
		}
	}
	return true
}

// AllOfMap checks that all elements in slice satisfy the condition
func AllOfMap(mapv interface{}, predicate func(key reflect.Value) bool) bool {
	rv := reflect.ValueOf(mapv)
	for _, k := range rv.MapKeys() {
		if !predicate() {
			return false
		}
	}
	return true
}
