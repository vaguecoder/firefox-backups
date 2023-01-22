package csv

import (
	"encoding/csv"
	"fmt"
	"io"

	"github.com/vaguecoder/firefox-backups/pkg/bookmark"
	"github.com/vaguecoder/firefox-backups/pkg/constants"
	"github.com/vaguecoder/firefox-backups/pkg/encoding"
)

var EncoderName = encoding.ToEncoder(constants.CSVFormat)

func init() {
	encoding.AllEncoders = append(encoding.AllEncoders, EncoderName)
}

type Encoder struct {
	csvEncoder   *csv.Writer
	enableHeader bool
}

func NewEncoder(out io.Writer, header bool) *Encoder {
	return &Encoder{
		csvEncoder:   csv.NewWriter(out),
		enableHeader: header,
	}
}

func (e *Encoder) Encode(bookmarks []bookmark.Bookmark) error {
	records := bookmark.BookmarksTable(bookmarks, e.enableHeader)

	err := e.csvEncoder.WriteAll(records)
	if err != nil {
		return fmt.Errorf("failed to marshal CSV: %v", err)
	}

	return nil
}

func (e *Encoder) String() string {
	return EncoderName.String()
}
