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

func ToQuotedString(s fmt.Stringer) quotedString {
	return quotedString(s.String())
}

func description[T flagTypes](usage string, defaultVal T, additionals []string) string {
	desc := fmt.Sprintf("%s (default %v)", usage, defaultVal)

	for _, a := range additionals {
		desc += fmt.Sprintf("\n%s", a)
	}

	return desc
}
