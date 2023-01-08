package flags

import "fmt"

func ptrToStr(s string) *string {
	return &s
}

type flagTypes interface {
	bool | quotedString
}

type quotedString string

func (q quotedString) String() string {
	return fmt.Sprintf(`"%s"`, string(q))
}

func description[T flagTypes](usage string, defaultVal T, additionals []string) string {
	var desc string = fmt.Sprintf("%s (default %v)", usage, defaultVal)

	for _, a := range additionals {
		desc += fmt.Sprintf("\n* %s", a)
	}

	return desc
}

func contains[T comparable](value T, list []T) bool {
	for _, elem := range list {
		if elem == value {
			return true
		}
	}
	return false
}
