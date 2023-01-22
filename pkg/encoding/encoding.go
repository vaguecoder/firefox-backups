package encoding

import (
	"fmt"
	"sort"
	"strings"

	"github.com/vaguecoder/firefox-backups/pkg/bookmark"
)

// EncoderName holds encoder's name
type EncoderName string

// String converts EncoderName to string,
// making EncoderName type implement fmt.Stringer
func (e EncoderName) String() string {
	return string(e)
}

// encoderNames is an unexported collection of encoder names
type encoderNames []EncoderName

// String converts encoders to string,
// making encoders type implement fmt.Stringer
func (e encoderNames) String() string {
	var encoders []string

	for _, encoderName := range e {
		encoders = append(encoders, encoderName.String())
	}

	sort.Strings(encoders)

	return strings.Join(encoders, ", ")
}

// AllEncoders holds list of encoder names.
// All the encoder names in echo of the encoder packages should
// be appended to AllEncoders during respective init()
var AllEncoders encoderNames

// ToEncoder converts stringer to EncoderName type
func ToEncoder(s fmt.Stringer) EncoderName {
	return EncoderName(s.String())
}

type Encoder interface {
	Encode([]bookmark.Bookmark) error
	fmt.Stringer
}
