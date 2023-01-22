package yaml

import (
	"fmt"
	"io"

	yaml "gopkg.in/yaml.v3"

	"github.com/vaguecoder/firefox-backups/pkg/bookmark"
	"github.com/vaguecoder/firefox-backups/pkg/constants"
	"github.com/vaguecoder/firefox-backups/pkg/encoding"
)

var EncoderName = encoding.ToEncoder(constants.YAMLFormat)

func init() {
	encoding.AllEncoders = append(encoding.AllEncoders, EncoderName)
}

type Encoder struct {
	yamlEncoder *yaml.Encoder
}

func NewEncoder(out io.Writer) *Encoder {
	encoder := yaml.NewEncoder(out)
	encoder.SetIndent(8)
	return &Encoder{
		yamlEncoder: encoder,
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
