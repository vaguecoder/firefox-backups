package encoding

import (
	"sort"
	"strings"
)

type EncoderName string

func (e EncoderName) String() string {
	return string(e)
}

type encoderNames []EncoderName

func (e encoderNames) String() string {
	var encoders []string

	for _, encoder := range e {
		encoders = append(encoders, encoder.String())
	}

	sort.Strings(encoders)

	return strings.Join(encoders, ", ")
}

var AllEncoders encoderNames
