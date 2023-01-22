package json

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/vaguecoder/firefox-backups/pkg/bookmark"
	"github.com/vaguecoder/firefox-backups/pkg/constants"
	"github.com/vaguecoder/firefox-backups/pkg/encoding"
)

var EncoderName = encoding.ToEncoder(constants.JSONFormat)

func init() {
	encoding.AllEncoders = append(encoding.AllEncoders, EncoderName)
}

type Encoder struct {
	jsonEncoder *json.Encoder
}

func NewEncoder(out io.Writer) *Encoder {
	encoder := json.NewEncoder(out)
	encoder.SetIndent("", "\t")
	return &Encoder{
		jsonEncoder: encoder,
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
