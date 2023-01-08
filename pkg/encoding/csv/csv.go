package csv

import "github.com/vaguecoder/firefox-backups/pkg/encoding"

const EncoderName encoding.EncoderName = `csv`

func init() {
	encoding.AllEncoders = append(encoding.AllEncoders, EncoderName)
}
