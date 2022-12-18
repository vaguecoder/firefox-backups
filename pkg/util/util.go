package util

import (
	"fmt"
	"regexp"
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

func ToStringers[T any](elems []T) []fmt.Stringer {
	var stringers []fmt.Stringer

	for _, e := range elems {
		if c, ok := any(e).(fmt.Stringer); ok {
			stringers = append(stringers, c)
		}

		// Ignored if element doesn't implement fmt.Stringer interface
	}

	return stringers
}
