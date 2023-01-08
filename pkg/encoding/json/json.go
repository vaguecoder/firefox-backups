package json

import "github.com/vaguecoder/firefox-backups/pkg/encoding"

const EncoderName encoding.EncoderName = `json`

func init() {
	encoding.AllEncoders = append(encoding.AllEncoders, EncoderName)
}
