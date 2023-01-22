package flags

import (
	"fmt"

	"github.com/vaguecoder/firefox-backups/pkg/encoding"
	"github.com/vaguecoder/firefox-backups/pkg/encoding/tabular"
	"github.com/vaguecoder/firefox-backups/pkg/filters"
)

const (
	silentFlagDefaultVal               = false
	inputSQLiteFileFlagDefaultVal      = "places.sqlite"
	rawFlagDefaultVal                  = false
	filterIgnoreDefaultsFlagDefaultVal = false
	filterDenormalizeFlagDefaultVal    = false
)

var (
	outputFormatFlagDefaultVal = ToQuotedString(tabular.EncoderName)

	silentFlagDesc = description(
		`Discard/suppress all the app logs`,
		silentFlagDefaultVal,
		[]string{`Enabled by default if writing to stdout`},
	)
	inputSQLiteFileFlagDesc = description[quotedString](
		`Input places.sqlite file path`,
		inputSQLiteFileFlagDefaultVal,
		nil,
	)
	rawFlagDesc = description(
		"Fetch all bookmarks without filtering",
		rawFlagDefaultVal,
		[]string{fmt.Sprintf("Available filters: [%s]", filters.AllFilterNames)},
	)
	stdOutFormatFlagDesc = description(
		`Stdout data format`,
		outputFormatFlagDefaultVal,
		[]string{fmt.Sprintf("Available formats: [%s]", encoding.AllEncoders)},
	)
	filterIgnoreDefaultsFlagDesc = description(
		"Ignore the default mozilla bookmarks from result",
		filterIgnoreDefaultsFlagDefaultVal,
		[]string{`false if --raw is enabled`},
	)
	filterDenormalizeFlagDesc = description(
		"Minimizes the list of bookmarks to only leaf bookmark records.",
		filterDenormalizeFlagDefaultVal,
		[]string{
			`false if --raw is enabled`,
			`Update the full directory path in leaf bookmarks and eliminate parent directory records.`,
			"  Eg.\n    Raw:\n" +
				"      ---------------------------------------------------------------\n" +
				"      | URL                   | TITLE       | FOLDER | ID  | PARENT |\n" +
				"      ---------------------------------------------------------------\n" +
				"      |                       | Profiles    |        | 1   | 0      |\n" +
				"      |                       | GitHub      |        | 2   | 1      |\n" +
				"      | github.com/vaguecoder | Vague Coder |        | 3   | 2      |\n" +
				"      ---------------------------------------------------------------\n",
			"    Denormalized:\n" +
				"      ------------------------------------------------------------------------\n" +
				"      | URL                   | TITLE       | FOLDER          | ID  | PARENT |\n" +
				"      ------------------------------------------------------------------------\n" +
				"      | github.com/vaguecoder | Vague Coder | Profiles/GitHub | 3   | 2      |\n" +
				"      ------------------------------------------------------------------------\n",
		},
	)
)
