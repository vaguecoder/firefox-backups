package util

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"

	"golang.org/x/exp/constraints"
)

const (
	whitespaceRegex      = `\s+`
	whitespace           = ` `
	lfChar          byte = 0xa
)

// StrsWhitespacesCleanup replaces multiple
// whitespaces/tab/newlines with single
// whitespace in all strings are slice
func StrsWhitespacesCleanup(strs []string) []string {
	var result []string
	for _, s := range strs {
		result = append(result, StrWhitespacesCleanup(s))
	}

	return result
}

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

// PtrInt returns the reference to integer of any type
func PtrInt[T constraints.Integer](i T) *T {
	return &i
}

// NonFileWriter is a wrapper over IO writer.
// This is for registering the interface for mocking.
type NonFileWriter interface {
	io.Writer
}

// StringSliceToFlatBytes joins string slice
// together with line feed (LF) character,
// making it []byte data.
// LF character in unix systems = 0xa
func StringSliceToFlatBytes(lines []string) []byte {
	var byteLine [][]byte

	for _, line := range lines {
		byteLine = append(byteLine, []byte(line))
	}

	byteLines := bytes.Join(byteLine, []byte{lfChar})

	return append(byteLines, lfChar)
}

// Joinables holds the types which are required
// to be joined together as string.
// Currently, it only has fmt.Stringer as required in project.
type Joinables interface {
	fmt.Stringer
}

// JoinAsString joins the Joinables type elements together
// with string seperator to return string value.
// Currently, it only has logic for fmt.Stringer alone.
func JoinAsString[E Joinables](elements []E, sep string) string {
	var joined string
	switch any(elements[0]).(type) {
	case fmt.Stringer:
		var strs []string
		for _, elem := range elements {
			strs = append(strs, elem.String())
		}

		joined = strings.Join(strs, sep)
	}

	return joined
}
