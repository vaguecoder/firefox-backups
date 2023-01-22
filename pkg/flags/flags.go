package flags

import (
	"flag"

	"github.com/vaguecoder/firefox-backups/pkg/constants"
	"github.com/vaguecoder/firefox-backups/pkg/encoding"
	_ "github.com/vaguecoder/firefox-backups/pkg/filters/denormalize"
	_ "github.com/vaguecoder/firefox-backups/pkg/filters/ignore-defaults"
)

type Flags struct {
	SQLiteDBFilename     string               `json:"input-sqlite-file"`
	RawOutput            bool                 `json:"raw"`
	Silent               bool                 `json:"silent"`
	OutputFormat         encoding.EncoderName `json:"output-format"`
	FilterIgnoreDefaults bool                 `json:"ignore-defaults"`
	FilterDenormalize    bool                 `json:"denormalize"`
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
		err error

		flags = Flags{
			SQLiteDBFilename:     "",
			RawOutput:            false,
			FilterIgnoreDefaults: false,
			Silent:               false,
			OutputFormat:         "",
		}
		stdOutFormat string
	)

	o.flagSet.StringVar(&flags.SQLiteDBFilename, constants.InputSQLiteFileFlag.String(), "", inputSQLiteFileFlagDesc) // Lazy assignment of default value
	o.flagSet.BoolVar(&flags.RawOutput, constants.RawFlag.String(), rawFlagDefaultVal, rawFlagDesc)
	o.flagSet.BoolVar(&flags.Silent, constants.SilentFlag.String(), silentFlagDefaultVal, silentFlagDesc)
	o.flagSet.StringVar(&stdOutFormat, constants.StdOutFormatFlag.String(), "", stdOutFormatFlagDesc) // Lazy assignment of default value

	o.flagSet.BoolVar(&flags.FilterIgnoreDefaults, constants.IgnoreDefaultsFlag.String(), filterIgnoreDefaultsFlagDefaultVal, filterIgnoreDefaultsFlagDesc)
	o.flagSet.BoolVar(&flags.FilterDenormalize, constants.DenormalizeFilter.String(), filterDenormalizeFlagDefaultVal, filterDenormalizeFlagDesc)

	if err = o.flagSet.Parse(o.args); err != nil {
		return nil, err
	}

	if flags.SQLiteDBFilename == "" {
		// Input filename is missing; assign default
		// Lazy assignment to avoid printing of default value in default format
		flags.SQLiteDBFilename = inputSQLiteFileFlagDefaultVal
	}

	if stdOutFormat == "" {
		// Output format is missing; assign default
		// Lazy assignment to avoid printing of default value in default format
		stdOutFormat = outputFormatFlagDefaultVal.String()
	}

	return &flags, nil
}
