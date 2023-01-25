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

var EncoderName = encoding.ToEncoder(constants.TabularFormat)

func init() {
	encoding.AllEncoders = append(encoding.AllEncoders, EncoderName)
}

type Encoder struct {
	tabEncoder   *tabwriter.Writer
	enableHeader bool
	filename     string
}

func NewEncoder(out io.Writer, header bool) *Encoder {
	var filename string
	if file, ok := any(out).(files.File); ok {
		filename = file.Name()
	}

	encoder := tabwriter.NewWriter(out, 0, 8, 8, ' ', tabwriter.TabIndent)
	return &Encoder{
		tabEncoder:   encoder,
		enableHeader: header,
		filename:     filename,
	}
}

func (e *Encoder) Encode(bookmarks []bookmark.Bookmark) error {
	records := bookmark.BookmarksTable(bookmarks, e.enableHeader)

	for _, r := range records {
		record := strings.Join(r, "\t") + "\t"
		fmt.Fprintln(e.tabEncoder, record)
	}

	e.tabEncoder.Flush()

	return nil
}

func (e *Encoder) String() string {
	return EncoderName.String()
}

func (e *Encoder) Filename() string {
	return e.filename
}
