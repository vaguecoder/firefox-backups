package flags

import (
	"fmt"

	pkgEncoding "github.com/vaguecoder/firefox-backups/pkg/encoding"
	"github.com/vaguecoder/firefox-backups/pkg/filters"
	pkgText "github.com/vaguecoder/firefox-backups/pkg/text"
	"github.com/vaguecoder/firefox-backups/pkg/util"
)

const (
	// Flag default values
	silentFlagDefaultVal               = false
	inputSQLiteFileFlagDefaultVal      = `places.sqlite`
	rawFlagDefaultVal                  = false
	filterIgnoreDefaultsFlagDefaultVal = false
	filterDenormalizeFlagDefaultVal    = false
)

var (
	// Aliasing function for localization
	whitespace = util.Whitespace
	table      = pkgText.Table
	appendAll  = util.AppendAll

	// Flag descriptions
	silentFlagDesc = description(
		`Discard all the app logs.`,
		silentFlagDefaultVal,
		[]string{`Enabled by default if writing to stdout to avoid conflict.`},
	)
	inputSQLiteFileFlagDesc = description[quotedString](
		`Input places.sqlite file path`,
		inputSQLiteFileFlagDefaultVal,
		nil,
	)
	rawFlagDesc = description(
		"Fetch all bookmarks without filtering.",
		rawFlagDefaultVal,
		[]string{
			fmt.Sprintf("Available filters: [%s].", filters.AllFilterNames),
			"Folder path in records in only set after denormalization. " +
				"It should be blank in raw mode.",
		},
	)
	stdOutFormatFlagDesc = description[quotedString](
		`Stdout data format.`,
		"",
		appendAll(
			fmt.Sprintf("Available formats: [%s].", pkgEncoding.AllEncoders),
			`Empty string "" for no bookmarks on stdout, i.e., to print app logs.`,
		),
	)
	filterIgnoreDefaultsFlagDesc = description(
		"Ignore the default mozilla bookmarks from result.",
		filterIgnoreDefaultsFlagDefaultVal,
		[]string{`false if --raw is enabled.`},
	)
	filterDenormalizeFlagDesc = description(
		"Minimizes the list of bookmarks to only leaf bookmark records.",
		filterDenormalizeFlagDefaultVal,
		appendAll(
			`false if --raw is enabled.`,
			`Update the full directory path in leaf bookmarks and eliminate parent directory records.`,
			whitespace(2)+"Eg.",
			whitespace(4)+"Raw:",
			table([][]string{
				{"url", "title", "folder", "id", "parent"},
				{"", "Projects", "", "1", "0"},
				{"", "GitHub", "", "2", "1"},
				{"github.com/vaguecoder/firefox-backups", "Firefox Backups", "", "3", "2"},
				{"github.com/vaguecoder/gorilla-mux", "Gorilla Mux Fork", "", "4", "2"},
			}, true, whitespace(6)),
			"",
			whitespace(4)+"Denormalized:",
			table([][]string{
				{"url", "title", "folder", "id", "parent"},
				{"github.com/vaguecoder/firefox-backups", "Firefox Backups", "Projects/GitHub", "3", "2"},
				{"github.com/vaguecoder/gorilla-mux", "Gorilla Mux Fork", "Projects/GitHub", "4", "2"},
			}, true, whitespace(6)),
		),
	)
	outputFilesFlagDesc = description("", &outputs{}, []string{})
)
