package flags

import (
	"flag"
	"fmt"

	"github.com/vaguecoder/firefox-backups/pkg/database"
	"github.com/vaguecoder/firefox-backups/pkg/encoding"
	_ "github.com/vaguecoder/firefox-backups/pkg/encoding/csv"
	_ "github.com/vaguecoder/firefox-backups/pkg/encoding/json"
	"github.com/vaguecoder/firefox-backups/pkg/encoding/tabular"
	_ "github.com/vaguecoder/firefox-backups/pkg/encoding/yaml"
)

const (
	silentFlagDefaultVal          = false
	inputSQLiteFileFlagDefaultVal = "places.sqlite"
	outputFilenameFlagDefaultVal  = ""
	rawFlagDefaultVal             = false
	ignoreDefaultsFlagDefaultVal  = false
)

var (
	outputFormatFlagDefaultVal = quotedString(tabular.EncoderName.String())
)

var (
	silentFlagDesc = description(`Discard/suppress all the app logs`,
		silentFlagDefaultVal, []string{`Enabled by default if writing to stdout`})
	inputSQLiteFileFlagDesc = description[quotedString](`Input places.sqlite file path`, inputSQLiteFileFlagDefaultVal, nil)
	outputFilenameFlagDesc  = description[quotedString]("Output filename",
		outputFilenameFlagDefaultVal, []string{`Empty value "" will write output to stdout`})
	rawFlagDesc = description("Fetch all bookmarks without filtering",
		rawFlagDefaultVal, []string{fmt.Sprintf("Available filters: [%s]", database.AllFilters)})
	ignoreDefaultsFlagDesc = description("Ignore the default mozilla bookmarks from result",
		ignoreDefaultsFlagDefaultVal, []string{`Overwritten as false if --raw is enabled`})
	outputFormatFlagDesc = description(`Output data format`,
		outputFormatFlagDefaultVal, []string{fmt.Sprintf("Available formats: [%s]", encoding.AllEncoders)})
)

type Flags struct {
	SQLiteDBFilename string  `json:"input-sqlite-file"`
	OutputFilename   *string `json:"output-filename"`
	RawOutput        bool    `json:"raw"`
	IgnoreDefaults   bool    `json:"ignore-defaults"`
	Silent           bool    `json:"silent"`
	OutputFormat     string  `json:"output-format"`
}

type Operator struct {
	args    []string
	flagSet *flag.FlagSet
}

func NewOperator(args []string) *Operator {
	return &Operator{
		args:    args,
		flagSet: flag.NewFlagSet("config flags", flag.ExitOnError),
	}
}

func (o *Operator) Parse() (*Flags, error) {
	var f = Flags{
		SQLiteDBFilename: "",
		OutputFilename:   ptrToStr(""),
		RawOutput:        false,
		IgnoreDefaults:   false,
		Silent:           false,
		OutputFormat:     "",
	}

	o.flagSet.StringVar(&f.SQLiteDBFilename, "input-sqlite-file", "", inputSQLiteFileFlagDesc) // Lazy assignment of default value
	o.flagSet.StringVar(f.OutputFilename, "output-filename", outputFilenameFlagDefaultVal, outputFilenameFlagDesc)
	o.flagSet.BoolVar(&f.RawOutput, "raw", rawFlagDefaultVal, rawFlagDesc)
	o.flagSet.BoolVar(&f.IgnoreDefaults, "ignore-defaults", ignoreDefaultsFlagDefaultVal, ignoreDefaultsFlagDesc)
	o.flagSet.BoolVar(&f.Silent, "silent", silentFlagDefaultVal, silentFlagDesc)
	o.flagSet.StringVar(&f.OutputFormat, "output-format", "", outputFormatFlagDesc) // Lazy assignment of default value

	if err := o.flagSet.Parse(o.args); err != nil {
		return nil, err
	}

	if f.SQLiteDBFilename == "" {
		// Input filename is missing; assign default
		// Lazy assignment to avoid printing of default value in default format
		f.SQLiteDBFilename = inputSQLiteFileFlagDefaultVal
	}

	if f.OutputFormat == "" {
		// Output format is missing; assign default
		// Lazy assignment to avoid printing of default value in default format
		f.OutputFormat = outputFormatFlagDefaultVal.String()
	}

	switch {
	case contains(encoding.EncoderName(f.OutputFormat), encoding.AllEncoders):
	default:
		return nil, fmt.Errorf("invalid output data format: %q", f.OutputFormat)
	}

	if *f.OutputFilename == "" {
		// Empty output file means printing to stdout
		f.OutputFilename = nil
	}

	if f.OutputFilename == nil {
		// When writing to stdout, silent mode should be enabled by default
		// Else, the app logs might coincide with output
		f.Silent = true
	}

	return &f, nil
}
