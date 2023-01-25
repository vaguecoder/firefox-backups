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

var EncoderName = encoding.ToEncoder(constants.JSONFormat)

func init() {
	encoding.AllEncoders = append(encoding.AllEncoders, EncoderName)
}

type Encoder struct {
	jsonEncoder *json.Encoder
	filename    string
}

func NewEncoder(out io.Writer) *Encoder {
	var filename string
	if file, ok := any(out).(files.File); ok {
		filename = file.Name()
	}

	encoder := json.NewEncoder(out)
	encoder.SetIndent("", "\t")
	return &Encoder{
		jsonEncoder: encoder,
		filename:    filename,
	}
}

func (e *Encoder) Encode(bookmarks []bookmark.Bookmark) error {
	err := e.jsonEncoder.Encode(bookmarks)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	return nil
}

func (e *Encoder) String() string {
	return EncoderName.String()
}

func (e *Encoder) Filename() string {
	return e.filename
}
