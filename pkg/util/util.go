package util

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	whitespaceRegex = `\s+`
	whitespace      = ` `
)

func StrWhitespacesCleanup(s string) string {
	space := regexp.MustCompile(whitespaceRegex)
	return space.ReplaceAllString(s, whitespace)
}

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

func Whitespace(n uint) string {
	return strings.Repeat(` `, int(n))
}
