package files

import (
	"context"
	"fmt"
	"io"
	"io/fs"
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
	Chmod(filename string, permission int) (ChmodRevertFunc, error)
}

type File interface {
	io.Reader
	io.Writer
	io.Closer
	Name() string
}

func NewOperator(ctx context.Context) FileOperator {
	return &Operator{
		logger: logs.FromContext(ctx),
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

type ChmodRevertFunc func(filename string) (os.FileMode, error)

func (o *Operator) Chmod(filename string, permission int) (ChmodRevertFunc, error) {
	fileStat, err := os.Stat(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch file %q 's stats: %v", filename, err)
	}

	revertFunc := func(filename string) (os.FileMode, error) {
		fileMode := fileStat.Mode()
		if err := chmod(filename, fileMode); err != nil {
			return fileMode, err
		}

		return fileMode, nil
	}

	if err = chmod(filename, fs.FileMode(permission)); err != nil {
		return nil, fmt.Errorf("failed to change the file %q 's permissions: %v", filename, err)
	}

	return revertFunc, nil
}

func chmod(filename string, permission fs.FileMode) error {
	err := os.Chmod(filename, permission)
	if err != nil {
		return fmt.Errorf("failed to change the permissions of %q to %d: %v",
			filename, permission, err)
	}

	return nil
}
