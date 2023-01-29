package tabular

import (
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/vaguecoder/firefox-backups/pkg/bookmark"
	"github.com/vaguecoder/firefox-backups/pkg/constants"
	"github.com/vaguecoder/firefox-backups/pkg/encoding"
	"github.com/vaguecoder/firefox-backups/pkg/files"
)

// fixedTabWidth is the tab width in table
const fixedTabWidth = 8

// EncoderName is name of the encoder in current package, i.e., Table.
// CSVFormat constant is parsed as EncoderName type here.
var EncoderName = encoding.ToEncoder(constants.TabularFormat)

func init() {
	// Register the encoder name in pkg/encoding.AllEncoders
	encoding.AllEncoders = append(encoding.AllEncoders, EncoderName)
}

// Encoder is the manager for table encoder
type Encoder struct {
	tabEncoder   *tabwriter.Writer
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

	// Set the tab width and indentation for the tab writer
	encoder := tabwriter.NewWriter(out, 0, 8, fixedTabWidth, ' ', tabwriter.TabIndent)
	return &Encoder{
		tabEncoder:   encoder,
		enableHeader: header,
		filename:     filename,
	}
}

// Encode encodes the input bookmarks in tabular format to already set output stream
func (e *Encoder) Encode(bookmarks []bookmark.Bookmark) error {
	var (
		err       error
		recordStr string
		record    []string

		records = bookmark.BookmarksTable(bookmarks, e.enableHeader)
	)

	for _, record = range records {
		// Additional tab character at the end for tabwriter to format the closing end
		recordStr = strings.Join(record, "\t") + "\t"

		fmt.Fprintln(e.tabEncoder, recordStr)
	}

	// Flush the tabwriter for avoiding any conflicts.
	// The errors in actual write are returned while flushing.
	if err = e.tabEncoder.Flush(); err != nil {
		return fmt.Errorf("failed at flushing tabwriter: %v", err)
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
