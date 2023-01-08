package yaml

import "github.com/vaguecoder/firefox-backups/pkg/encoding"

const EncoderName encoding.EncoderName = `yaml`

func init() {
	encoding.AllEncoders = append(encoding.AllEncoders, EncoderName)
}
