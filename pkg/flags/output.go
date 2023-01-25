package flags

import (
	"fmt"
	"sort"
	"strings"

	pkgConstants "github.com/vaguecoder/firefox-backups/pkg/constants"
	pkgEncoding "github.com/vaguecoder/firefox-backups/pkg/encoding"
)

const (
	outputFormatFilenameDelimiter = `:`
	outputFilesDelimiter          = `,`
)

type OutputFile struct {
	Format   pkgConstants.Constant[pkgConstants.OutputFormat]
	Filename string
}

type outputs []OutputFile

func (o *outputs) Slice() []OutputFile {
	return []OutputFile(*o)
}

func (o *outputs) String() string {
	var outputs []string

	for _, output := range *o {
		outputStr := fmt.Sprintf("%s%s%s", output.Format,
			outputFormatFilenameDelimiter, output.Filename)
		outputs = append(outputs, outputStr)
	}

	return strings.Join(outputs, ",")
}

func (o *outputs) Set(s string) error {
	var (
		index                    int
		splits                   []string
		format                   pkgConstants.Constant[pkgConstants.OutputFormat]
		filename, formatFilename string
		output                   OutputFile
	)

	// Iterate over format-filename sets
	for index, formatFilename = range strings.Split(s, outputFilesDelimiter) {

		// Split format and filename of current set
		splits = strings.Split(formatFilename, outputFormatFilenameDelimiter)

		// Input validation (1/2): Arguments number validation
		if len(splits) < 2 || splits[1] == "" {
			// When both format & filename are missing, or filename is empty

			if len(splits) == 0 || splits[0] == "" {
				// When both format is missing or is empty

				if index == 0 {
					// When current format-filename is the first set and
					// both format is missing or is empty

					return fmt.Errorf("missing argument in format --%s=<format>%s<filename>[%s<format>%s<filename>%s...]",
						pkgConstants.OutputFiles, outputFormatFilenameDelimiter, outputFilesDelimiter,
						outputFormatFilenameDelimiter, outputFilesDelimiter)
				}

				// When current format-filename is not the first set and
				// both format is missing or is empty

				return fmt.Errorf("missing arguments after delimiter %q in --%s=%s%s<format>%s<filename>",
					outputFilesDelimiter, pkgConstants.OutputFiles, o, outputFilesDelimiter, outputFormatFilenameDelimiter)
			}

			// When only filename is missing

			if index == 0 {
				// When current format-filename is the first set and
				// only filename is missing
				return fmt.Errorf("missing filename in --%s=%s%s<filename>",
					pkgConstants.OutputFiles, splits[0], outputFormatFilenameDelimiter)
			}

			// When current format-filename is not the first set and
			// only filename is missing

			return fmt.Errorf("missing filename in --%s=%s%s%s%s<filename>",
				pkgConstants.OutputFiles, o, outputFilesDelimiter, splits[0], outputFormatFilenameDelimiter)
		}

		format = pkgConstants.Constant[pkgConstants.OutputFormat](splits[0])
		filename = splits[1]

		// Clean-up
		output = OutputFile{
			Filename: filename,
		}

		// Input validation (2/2): File format validation
		switch format {
		case pkgConstants.CSVFormat, pkgConstants.JSONFormat, pkgConstants.TabularFormat, pkgConstants.YAMLFormat:
			output.Format = format
		default:
			return fmt.Errorf("invalid output format in --%s=<format>%s<filename> (allowed formats: %v)",
				pkgConstants.OutputFiles, outputFormatFilenameDelimiter, pkgEncoding.AllEncoders)
		}

		// Add current format-filename set to output files list
		*o = append(*o, output)
	}

	// Sort output sets based on alphabetical order of file formats
	sort.Slice(o.Slice(), func(i, j int) bool {
		// Convert Outputs type to []Output type
		outputs := o.Slice()

		// Get the first letter of Output.Format value
		firstLetter := func(output OutputFile) rune {
			// Rune of format's first letter
			return rune(output.Format[0])
		}

		// Less-than condition for slice sorting
		return firstLetter(outputs[i]) < firstLetter(outputs[j])
	})

	return nil
}
