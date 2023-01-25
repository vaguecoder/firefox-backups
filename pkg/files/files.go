package files

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/vaguecoder/firefox-backups/pkg/logs"
)

type Operator struct {
	logger logs.Logger
}

type FileOperator interface {
	Copy(src, dest string) error
	Delete(filename string) error
	Open(filename string) (File, error)
}

type File interface {
	io.Reader
	io.Writer
	io.Closer
	Name() string
}

func NewOperator(l logs.Logger) FileOperator {
	return &Operator{
		logger: l,
	}
}

func (o *Operator) Copy(src, dest string) error {
	logger := o.logger.With().Str("src", src).
		Str("dest", dest).Logger()

	input, err := os.ReadFile(src)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to read the source file")
		return fmt.Errorf("failed to read the file %q: %v", src, err)
	}

	err = os.WriteFile(dest, input, 0644)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to write to destination file")
		return fmt.Errorf("failed to write the file %q: %v", dest, err)
	}

	logger.Info().Msg("Successfully copied file")

	return nil
}

func (o *Operator) Delete(filename string) error {
	filenamePattern := fmt.Sprintf("%s*", filename)
	logger := o.logger.With().Str("filename", filename).
		Str("pattern", filenamePattern).Logger()

	files, err := filepath.Glob(filenamePattern)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to fetch list of files matching pattern")
		return fmt.Errorf("failed to fetch list of files matching pattern %q: %v", filenamePattern, err)
	}

	logger = logger.With().Strs("matching-files", files).
		Int("totalFiles", len(files)).Logger()

	for i, f := range files {
		if err := os.RemoveAll(f); err != nil {
			logger.Error().Err(err).
				Str("current-file", f).
				Int("deleted-count", i-1).
				Msg("Failed to revove file")
			return fmt.Errorf("failed to remove file %q: %v", f, err)
		}
		logger.Info()
	}

	logger.Info().Msg("File(s) deleted successfully")

	return nil
}

func (o *Operator) Open(filename string) (File, error) {
	logger := o.logger.With().Str("filename", filename).Logger()

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to open/create file")
		return nil, fmt.Errorf("failed to open/create file %q: %v", filename, err)
	}

	o.logger.Info().Msg("Successfully opened/created file")

	return file, nil
}
