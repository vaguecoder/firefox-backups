package constants

type (
	OutputFormat string // Output format constants: JSON, YAML, CSV, Tabular
	Filter       string // Bookmark filter constants: denormalize, ignore-defaults
	Flag         string // Input flag name constants: input-sqlite-file, output-filename, etc.
)

// stringer is a custom stringer interface which defines String method on underlying types
type stringer interface {
	OutputFormat | Filter | Flag
}

// Constant is a stringer type wound on string
type Constant[S stringer] string

// String methods returns the string representation of the receiver type.
// This makes the receiver type to implement a fmt.Stringer & stringer
func (s Constant[stringer]) String() string {
	return string(s)
}

const (
	// Output format constants
	JSONFormat    Constant[OutputFormat] = `json`
	YAMLFormat    Constant[OutputFormat] = `yaml`
	CSVFormat     Constant[OutputFormat] = `csv`
	TabularFormat Constant[OutputFormat] = `table`

	// Bookmark filter constants
	DenormalizeFilter    Constant[Filter] = `denormalize`
	IgnoreDefaultsFilter Constant[Filter] = `ignore-defaults`

	// Input flag name constants
	InputSQLiteFileFlag Constant[Flag] = `input-sqlite-file`
	RawFlag             Constant[Flag] = `raw`
	IgnoreDefaultsFlag  Constant[Flag] = `ignore-defaults`
	SilentFlag          Constant[Flag] = `silent`
	StdOutFormatFlag    Constant[Flag] = `stdout-format`
	DenormalizeFlag     Constant[Flag] = `denormalize`
	OutputFiles         Constant[Flag] = `output-files`
)
