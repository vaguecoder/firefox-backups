package files_test

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vaguecoder/firefox-backups/pkg/files"
	"github.com/vaguecoder/firefox-backups/pkg/logs"
)

// TODO: Remove unused types after full coverage
type writable interface {
	[]byte | string | []string
}

// TODO: Remove unused types after full coverage
type readers interface {
	*os.File | string
}

func read[R readers](reader R) ([]byte, error) {
	var (
		err        error
		fileReader *os.File
		filename   string
		data       []byte
	)
	switch any(reader).(type) {
	case *os.File:
		fileReader = any(reader).(*os.File)

		data, err = io.ReadAll(fileReader)
		if err != nil {
			return nil, fmt.Errorf("failed to read from file in test: %v", err)
		}

		return data, nil

	case string:
		filename = any(reader).(string)
		fileReader, err = os.Open(filename)
		if err != nil {
			return nil, fmt.Errorf("failed to open file %q in test: %v", filename, err)
		}

		data, err = io.ReadAll(fileReader)
		if err != nil {
			return nil, fmt.Errorf("failed to read from filename in test: %v", err)
		}

		return data, nil
	default:
		return nil, fmt.Errorf("invalid type to read data in test. " +
			"Probably added new type readers interface which is missing in switch case")
	}
}

func write[W writable](writer io.Writer, data W) error {
	var (
		err       error
		dataBytes []byte
		dataStr   string
		dataStrs  []string

		bufWriter = bufio.NewWriter(writer)
	)

	switch any(data).(type) {
	case []byte:
		dataBytes = any(data).([]byte)

		_, err = bufWriter.Write(dataBytes)
		if err != nil {
			return fmt.Errorf("failed to write byte data to writer in test: %v", err)
		}

		err = bufWriter.Flush()
		if err != nil {
			return fmt.Errorf("failed to flush writer in test: %v", err)
		}

		return nil
	case string:
		dataStr = any(data).(string)

		_, err = bufWriter.WriteString(dataStr)
		if err != nil {
			return fmt.Errorf("failed to write string data to writer in test: %v", err)
		}

		err = bufWriter.Flush()
		if err != nil {
			return fmt.Errorf("failed to flush writer in test: %v", err)
		}

		return nil
	case []string:
		dataStrs = any(data).([]string)
		dataStr = strings.Join(dataStrs, "\n")

		_, err = bufWriter.WriteString(dataStr)
		if err != nil {
			return fmt.Errorf("failed to write strings data to writer in test: %v", err)
		}

		err = bufWriter.Flush()
		if err != nil {
			return fmt.Errorf("failed to flush writer in test: %v", err)
		}

		return nil
	default:
		return fmt.Errorf("invalid type to write data in test. " +
			"Probably added new type writable interface which is missing in switch case")
	}
}

func createFile(filename string) (*os.File, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0600)
	if err != nil {
		return nil, fmt.Errorf("failed to create test file %s: %v", filename, err)
	}

	return file, nil
}

func deleteFile(filename string) error {
	err := os.RemoveAll(filename)
	if err != nil {
		return fmt.Errorf("failed to delete file/path %q: %v", filename, err)
	}

	return nil
}

func TestOperator_Copy(t *testing.T) {
	var (
		ctx     = context.Background()
		file    *os.File
		err     error
		fileOps files.FileOperator
		data    []byte
	)

	type args struct {
		src  string
		dest string
	}
	tests := []struct {
		name    string
		args    args
		data    []byte
		wantErr bool
	}{
		{
			name: "Simple",
			args: args{
				src:  "testdata/source.txt",
				dest: "testdata/desrination.txt",
			},
			data:    []byte("This is a test"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logOutStream := bytes.NewBuffer([]byte{})
			ctx, _ = logs.NewLogger(ctx, logOutStream, logs.LevelDebug)

			file, err = createFile(tt.args.src)
			require.NoError(t, err, "Unexpected error while creating a test file")

			err = write(file, tt.data)
			require.NoError(t, err, "Unexpected error while writing to the test file")

			fileOps = files.NewOperator(ctx)
			err = fileOps.Copy(tt.args.src, tt.args.dest)
			if tt.wantErr {
				assert.Error(t, err, "Expected error from Copy")
			} else {
				assert.NoError(t, err, "Unexpected error from Copy")
			}

			data, err = read(tt.args.dest)
			require.NoError(t, err, "Unexpected error from test reading destination file")

			assert.Equal(t, string(tt.data), string(data), "Mismatch of source and destination data")

			err = deleteFile(tt.args.src)
			require.NoError(t, err, "Unexpected error from deleting test source file")

			err = deleteFile(tt.args.dest)
			require.NoError(t, err, "Unexpected error from deleting test destination file")
		})
	}
}
