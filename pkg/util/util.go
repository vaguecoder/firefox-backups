package util

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/exp/constraints"
)

const (
	whitespaceRegex = `\s+`
	whitespace      = ` `
)

// StrWhitespacesCleanup replaces multiple
// whitespaces/tab/newlines with single whitespace.
func StrWhitespacesCleanup(s string) string {
	space := regexp.MustCompile(whitespaceRegex)
	return space.ReplaceAllString(s, whitespace)
}

// MapKeys returns slice of keys from map of any type
func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0)

	if len(m) == 0 {
		return keys
	}

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

// ToStringers wraps elements of input slice with stringer type,
// provided, all the elements on input slice should implement fmt.Stringer.
func ToStringers[T any](elems []T) ([]fmt.Stringer, error) {
	var stringers []fmt.Stringer

	for _, e := range elems {
		if c, ok := any(e).(fmt.Stringer); ok {
			stringers = append(stringers, c)
			continue
		}

		return nil, fmt.Errorf("element not a stringer: %+v", e)

	}

	return stringers, nil
}

// AppendAll appends strings and slice of strings together into slice of strings
func AppendAll(strs ...interface{}) []string {
	var result []string
	for _, s := range strs {
		switch v := s.(type) {
		case string:
			result = append(result, v)
		case []string:
			result = append(result, v...)
		default:
			// Ignored.
			// Should be the caller's responsibility to send only accepted types.
			// No return error as this function is expected to be used in arguments
			//  or global variables similar to built-in append().
		}
	}

	return result
}

// Whitespace returns whitespace string of length n
func Whitespace(n uint) string {
	return strings.Repeat(` `, int(n))
}

// PtrStr returns pointer to string
func PtrStr(s string) *string {
	return &s
}

// PtrInt
func PtrInt[T constraints.Integer](i T) *T {
	return &i
}
