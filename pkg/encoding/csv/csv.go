package csv

import (
	"encoding/csv"
	"fmt"
	"io"

	"github.com/vaguecoder/firefox-backups/pkg/bookmark"
	"github.com/vaguecoder/firefox-backups/pkg/constants"
	"github.com/vaguecoder/firefox-backups/pkg/encoding"
	"github.com/vaguecoder/firefox-backups/pkg/files"
)

var EncoderName = encoding.ToEncoder(constants.CSVFormat)

func init() {
	encoding.AllEncoders = append(encoding.AllEncoders, EncoderName)
}

type Encoder struct {
	csvEncoder   *csv.Writer
	enableHeader bool
	filename     string
}

func NewEncoder(out io.Writer, header bool) *Encoder {
	var filename string
	if file, ok := any(out).(files.File); ok {
		filename = file.Name()
	}

	return &Encoder{
		csvEncoder:   csv.NewWriter(out),
		enableHeader: header,
		filename:     filename,
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

func (e *Encoder) Filename() string {
	return e.filename
}
