package algorithm

import (
	"reflect"
)

// AllOf checks that all elements in slice satisfy the condition
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
