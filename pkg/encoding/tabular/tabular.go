package tabular

import "github.com/vaguecoder/firefox-backups/pkg/encoding"

const EncoderName encoding.EncoderName = `tabular`

func init() {
	encoding.AllEncoders = append(encoding.AllEncoders, EncoderName)
}
