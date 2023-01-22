package tabular

import (
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/vaguecoder/firefox-backups/pkg/bookmark"
	"github.com/vaguecoder/firefox-backups/pkg/constants"
	"github.com/vaguecoder/firefox-backups/pkg/encoding"
)

var EncoderName = encoding.ToEncoder(constants.TabularFormat)

func init() {
	encoding.AllEncoders = append(encoding.AllEncoders, EncoderName)
}

type Encoder struct {
	tabEncoder   *tabwriter.Writer
	enableHeader bool
}

func NewEncoder(out io.Writer, header bool) *Encoder {
	encoder := tabwriter.NewWriter(out, 0, 8, 8, ' ', tabwriter.TabIndent)
	return &Encoder{
		tabEncoder:   encoder,
		enableHeader: header,
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
