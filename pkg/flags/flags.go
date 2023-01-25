package flags

import (
	"flag"
	"fmt"
	"os"

	"github.com/vaguecoder/firefox-backups/pkg/constants"
	"github.com/vaguecoder/firefox-backups/pkg/encoding"
	pkgEncoding "github.com/vaguecoder/firefox-backups/pkg/encoding"
	pkgEncodingCSV "github.com/vaguecoder/firefox-backups/pkg/encoding/csv"
	pkgEncodingJSON "github.com/vaguecoder/firefox-backups/pkg/encoding/json"
	pkgEncodingTab "github.com/vaguecoder/firefox-backups/pkg/encoding/tabular"
	pkgEncodingYAML "github.com/vaguecoder/firefox-backups/pkg/encoding/yaml"
	_ "github.com/vaguecoder/firefox-backups/pkg/filters/denormalize"
	_ "github.com/vaguecoder/firefox-backups/pkg/filters/ignore-defaults"
)

type Flags struct {
	SQLiteDBFilename     string           `json:"input-sqlite-file"`
	RawOutput            bool             `json:"raw"`
	Silent               bool             `json:"silent"`
	OutputFiles          outputs          `json:"output-files"`
	StdOutFormat         encoding.Encoder `json:"stdout-format"`
	FilterIgnoreDefaults bool             `json:"ignore-defaults"`
	FilterDenormalize    bool             `json:"denormalize"`
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
	var (
		err          error
		stdOutFormat string

		flags = Flags{
			SQLiteDBFilename:     "",
			RawOutput:            false,
			Silent:               false,
			OutputFiles:          []OutputFile{},
			StdOutFormat:         nil,
			FilterIgnoreDefaults: false,
			FilterDenormalize:    false,
		}
		outputFiles = outputs{}
	)

	// General input flags
	o.flagSet.StringVar(&flags.SQLiteDBFilename, constants.InputSQLiteFileFlag.String(), "", inputSQLiteFileFlagDesc) // Lazy assignment of default value
	o.flagSet.BoolVar(&flags.RawOutput, constants.RawFlag.String(), rawFlagDefaultVal, rawFlagDesc)
	o.flagSet.BoolVar(&flags.Silent, constants.SilentFlag.String(), silentFlagDefaultVal, silentFlagDesc)
	o.flagSet.StringVar(&stdOutFormat, constants.StdOutFormatFlag.String(), "", stdOutFormatFlagDesc) // Lazy assignment of default value

	// Filter input flag
	o.flagSet.BoolVar(&flags.FilterIgnoreDefaults, constants.IgnoreDefaultsFlag.String(), filterIgnoreDefaultsFlagDefaultVal, filterIgnoreDefaultsFlagDesc)
	o.flagSet.BoolVar(&flags.FilterDenormalize, constants.DenormalizeFilter.String(), filterDenormalizeFlagDefaultVal, filterDenormalizeFlagDesc)

	// Output file format input flag with custom implementation of flags.Value interface
	o.flagSet.Var(&outputFiles, constants.OutputFiles.String(), outputFilesFlagDesc)

	if err = o.flagSet.Parse(o.args); err != nil {
		// When parsing of input flag arguments failed
		return nil, fmt.Errorf("failed to parse input flag args: %v", err)
	}

	if flags.SQLiteDBFilename == "" {
		// Input filename is missing; assign default
		// Lazy assignment to avoid printing of default value in default format
		flags.SQLiteDBFilename = inputSQLiteFileFlagDefaultVal
	}

	if stdOutFormat != "" {
		// When flag --stdout-format is provided with a non-empty string
		switch constants.Constant[constants.OutputFormat](stdOutFormat) {
		case constants.CSVFormat:
			// CSV format
			flags.StdOutFormat = pkgEncodingCSV.NewEncoder(os.Stdout, true)
		case constants.JSONFormat:
			// JSON format
			flags.StdOutFormat = pkgEncodingJSON.NewEncoder(os.Stdout)
		case constants.TabularFormat:
			// Table format
			flags.StdOutFormat = pkgEncodingTab.NewEncoder(os.Stdout, true)
		case constants.YAMLFormat:
			// YAML format
			flags.StdOutFormat = pkgEncodingYAML.NewEncoder(os.Stdout)
		default:
			// Unaccepted output format to stdout-flag
			return nil, fmt.Errorf("invalid format '%s' to --%s flag (available formats: [%s])",
				stdOutFormat, constants.StdOutFormatFlag, pkgEncoding.AllEncoders)
		}

		// When the resultant bookmarks be printed on stdout, the app logs should be suppressed
		flags.Silent = true
	}

	// Append output format-filename sets after validation.
	// It is validated and formatted at flags.Value interface level.
	flags.OutputFiles = append(flags.OutputFiles, outputFiles...)

	return &flags, nil
}
