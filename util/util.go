package util

import (
	"strings"
)

// Comparator is a function type that compares two elements.
// The return value is negative if a < b, 0 if a == b, and positive if a > b.
type Comparator[T any] func(o1, o2 T) int

// Equals is a function type that compares the equality of two elements.
type Equals[E any] func(a, b E) bool

// defaultComparator is the default comparator function, used when the element is comparable.
// It is useful when the element is an int, float64, string.
func DefaultComparator[T any]() Comparator[T] {
	return func(a, b T) int {
		switch aTyped := any(a).(type) {
		case int:
			bTyped := any(b).(int)
			return aTyped - bTyped
		case float64:
			bTyped := any(b).(float64)
			if aTyped < bTyped {
				return -1
			} else if aTyped > bTyped {
				return 1
			}
			return 0
		case string:
			bTyped := any(b).(string)
			return strings.Compare(aTyped, bTyped)
		default:
			return 0
		}
	}
}
