package algorithm

import (
	"testing"
)

func TestAllOf(t *testing.T) {
	{
		s := []uint{1, 3, 5}
		expected := true
		got := AllOf(s, func(i int) bool { return s[i]%2 == 1 })
		if got != expected {
			t.Errorf("unexpected result: got:%v, expected:%v", got, expected)
		}
	}
	{
		s := []uint{1, 4, 5}
		expected := false
		got := AllOf(s, func(i int) bool { return s[i]%2 == 1 })
		if got != expected {
			t.Errorf("unexpected result: got:%v, expected:%v", got, expected)
		}
	}
	{
		s := []string{"apple", "banana", "orange"}
		expected := true
		got := AllOf(s, func(i int) bool { return len(s[i]) >= 5 })
		if got != expected {
			t.Errorf("unexpected result: got:%v, expected:%v", got, expected)
		}
	}
}
