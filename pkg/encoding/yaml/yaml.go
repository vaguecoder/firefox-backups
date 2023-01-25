package yaml

import (
	"fmt"
	"io"

	yaml "gopkg.in/yaml.v3"

	"github.com/vaguecoder/firefox-backups/pkg/bookmark"
	"github.com/vaguecoder/firefox-backups/pkg/constants"
	"github.com/vaguecoder/firefox-backups/pkg/encoding"
	"github.com/vaguecoder/firefox-backups/pkg/files"
)

var EncoderName = encoding.ToEncoder(constants.YAMLFormat)

func init() {
	encoding.AllEncoders = append(encoding.AllEncoders, EncoderName)
}

type Encoder struct {
	yamlEncoder *yaml.Encoder
	filename    string
}

func NewEncoder(out io.Writer) *Encoder {
	var filename string
	if file, ok := any(out).(files.File); ok {
		filename = file.Name()
	}

	encoder := yaml.NewEncoder(out)
	encoder.SetIndent(8)
	return &Encoder{
		yamlEncoder: encoder,
		filename:    filename,
	}
}

func (e *Encoder) Encode(bookmarks []bookmark.Bookmark) error {
	err := e.yamlEncoder.Encode(bookmarks)
	if err != nil {
		return fmt.Errorf("failed to marshal YAML: %v", err)
	}

	return nil
}

func (e *Encoder) String() string {
	return EncoderName.String()
}

func (e *Encoder) Filename() string {
	return e.filename
}
