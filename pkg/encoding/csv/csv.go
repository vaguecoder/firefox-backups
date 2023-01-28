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

// EncoderName is name of the encoder in current package, i.e., CSV.
// CSVFormat constant is parsed as EncoderName type here.
var EncoderName = encoding.ToEncoder(constants.CSVFormat)

func init() {
	// Register the encoder name in pkg/encoding.AllEncoders
	encoding.AllEncoders = append(encoding.AllEncoders, EncoderName)
}

// Encoder is the manager for CSV encoder
type Encoder struct {
	csvEncoder   *csv.Writer
	enableHeader bool
	filename     string
}

// NewEncoder initializes new Encoder
func NewEncoder(out io.Writer, header bool) *Encoder {
	var filename string

	// If the output stream is a file, specifically pkg/files.File,
	// the filename can be extracted here. Just an optional requirement.
	if file, ok := any(out).(files.File); ok {
		filename = file.Name()
	}

	return &Encoder{
		csvEncoder:   csv.NewWriter(out),
		enableHeader: header,
		filename:     filename,
	}
}

// Encode encodes the input bookmarks in CSV format to already set output stream
func (e *Encoder) Encode(bookmarks []bookmark.Bookmark) error {
	records := bookmark.BookmarksTable(bookmarks, e.enableHeader)

	err := e.csvEncoder.WriteAll(records)
	if err != nil {
		return fmt.Errorf("failed to marshal CSV: %v", err)
	}

	return nil
}

// String returns the encoder name derived in EncoderName.
// This returns the same value as EncoderName, but using the receiver.
func (e *Encoder) String() string {
	return EncoderName.String()
}

// Filename returns the file name string derived from output stream,
// iff the output stream is of pkg/files.File type.
func (e *Encoder) Filename() string {
	return e.filename
}
