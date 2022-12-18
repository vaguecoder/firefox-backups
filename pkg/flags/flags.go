package flags

import "flag"

type Flags struct {
	SQLiteDBFilename    string `json:"sqlite-filename"`
	OutputJSONFilename  string `json:"output-filename"`
	EnableNormalization bool   `json:"normalize"`
	IgnoreDefaults      bool   `json:"ignore-defaults"`
	WriteToFile         bool   `json:"write-to-file"`
	Silent              bool   `json:"silent"`
}

type Operator struct {
	args    []string
	flagSet *flag.FlagSet
}

func NewOperator(args []string) *Operator {
	return &Operator{
		args:    args,
		flagSet: flag.NewFlagSet("configs", flag.ExitOnError),
	}
}

func (o *Operator) Parse() (*Flags, error) {
	var f Flags

	o.flagSet.StringVar(&f.SQLiteDBFilename, "sqlite-filename", "places.sqlite", "full path to places.sqlite file")
	o.flagSet.StringVar(&f.OutputJSONFilename, "output-filename", "bookmarks.json", "output JSON filename")
	o.flagSet.BoolVar(&f.WriteToFile, "write-to-file", false, "enable writing to file")
	o.flagSet.BoolVar(&f.EnableNormalization, "normalize", false, "enable normalize to fetch only the leaf bookmarks saved")
	o.flagSet.BoolVar(&f.IgnoreDefaults, "ignore-defaults", false, "enable this flag to ignore the default mozilla bookmarks from result")
	o.flagSet.BoolVar(&f.Silent, "silent", false, "enable silent mode to discard all the logs")

	if err := o.flagSet.Parse(o.args); err != nil {
		return nil, err
	}

	return &f, nil
}
