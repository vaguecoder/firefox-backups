package json

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/vaguecoder/firefox-backups/pkg/bookmark"
	"github.com/vaguecoder/firefox-backups/pkg/constants"
	"github.com/vaguecoder/firefox-backups/pkg/encoding"
	"github.com/vaguecoder/firefox-backups/pkg/files"
)

// indentation is the JSON indentation string
const indentation = "\t"

// EncoderName is name of the encoder in current package, i.e., JSON.
// JSONFormat constant is parsed as EncoderName type here.
var EncoderName = encoding.ToEncoder(constants.JSONFormat)

func init() {
	// Register the encoder name in pkg/encoding.AllEncoders
	encoding.AllEncoders = append(encoding.AllEncoders, EncoderName)
}

// Encoder is the manager for JSON encoder
type Encoder struct {
	jsonEncoder *json.Encoder
	filename    string
}

// NewEncoder initializes new Encoder
func NewEncoder(out io.Writer) *Encoder {
	var filename string

	// If the output stream is a file, specifically pkg/files.File,
	// the filename can be extracted here. Just an optional requirement.
	if file, ok := any(out).(files.File); ok {
		filename = file.Name()
	}

	// Create new JSON encoder and set the indentation per constant.
	encoder := json.NewEncoder(out)
	encoder.SetIndent("", indentation)

	return &Encoder{
		jsonEncoder: encoder,
		filename:    filename,
	}
}

// Encode encodes the input bookmarks in JSON format to already set output stream
func (e *Encoder) Encode(bookmarks []bookmark.Bookmark) error {
	err := e.jsonEncoder.Encode(bookmarks)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	return nil
}

// String returns the encoder name derived in EncoderName.
// This returns the same value as EncoderName, but using the receiver.
func (e *Encoder) String() string {
	return EncoderName.String()
}

// Filename returns the file name string derived from output stream,
// iff the output stream is of pkg/files.File type.
func (e *Encoder) Filename() string {
	return e.filename
}
