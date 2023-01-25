package flags

import (
	"fmt"
)

// flagTypes holds a selective list of types that are used as default values
// in input flag's descriptions.
//
//  1. bool - Boolean flags are used for just true/false.
//  2. quotedString - Is a wrapper over string typed flags that encloses the result in quotes.
//  3. *outputs - Is a slice of set of output file format and filename.
//     Reference of outputs is required as one of the other function on outputs
//     i.e., Set() requires pointer receiver to update. Hence, outputs instance
//     usually is used with pointer.
type flagTypes interface {
	bool | quotedString | *outputs
}

// quotedString encloses string in quotes
type quotedString string

// String returns string value of quotedString
// making quotedString implement fmt.Stringer interface
func (q quotedString) String() string {
	return fmt.Sprintf("%q", string(q))
}

// toQuotedString converts string as quotedString
func toQuotedString(s fmt.Stringer) quotedString {
	return quotedString(s.String())
}

// description processes the inputs to return the description string
func description[T flagTypes](usage string, defaultValue T, additionals []string) string {
	var (
		desc, additionalStr string
		defValHolder        interface{}
		ok                  bool
		outputsVal          *outputs
		stringerVal         fmt.Stringer
	)

	// Convert *outputs type to quotedString as it adds quotes in description-level
	if outputsVal, ok = any(defaultValue).(*outputs); ok {
		defValHolder = toQuotedString(outputsVal)
	} else {
		defValHolder = defaultValue
	}

	// Print default value in desciption in string format for stringer
	// and value format for non-stringer type
	if stringerVal, ok = defValHolder.(fmt.Stringer); ok {
		desc = fmt.Sprintf("%s (default %s)", usage, stringerVal)
	} else {
		desc = fmt.Sprintf("%s (default %v)", usage, defValHolder)
	}

	// Add additional lines to desciption
	for _, additionalStr = range additionals {
		desc += fmt.Sprintf("\n%s", additionalStr)
	}

	return desc
}
