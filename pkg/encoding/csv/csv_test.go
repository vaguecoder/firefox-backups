package csv

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vaguecoder/firefox-backups/pkg/bookmark"
	"github.com/vaguecoder/firefox-backups/pkg/constants"
	"github.com/vaguecoder/firefox-backups/pkg/mocks"
	"github.com/vaguecoder/firefox-backups/pkg/util"
)

var (
	ptrStr                 = util.PtrStr
	stringSliceToFlatBytes = util.StringSliceToFlatBytes
)

func TestEncoder_Encode(t *testing.T) {
	type fields struct {
		enableHeader bool
		filename     string
		outputWriter io.Writer
	}

	type args struct {
		bookmarks []bookmark.Bookmark
	}

	type toggles struct {
		isExpectedError      bool
		isFileWriter         bool
		isErrorAtCSVWriteAll bool
	}

	type testData struct {
		name     string
		fields   fields
		args     args
		toggles  toggles
		expected []string
	}

	var (
		err            error
		tests          []testData
		testCase       testData
		encoder        *Encoder
		filename       string
		expectedOutput []byte

		fileWriter    = new(mocks.File)
		nonFileWriter = new(mocks.NonFileWriter)
	)

	tests = []testData{
		{
			name: "Valid-Case-Disable-Header",
			fields: fields{
				enableHeader: false,
				filename:     "firefox-bookmarks.csv",
				outputWriter: fileWriter,
			},
			args: args{
				bookmarks: []bookmark.Bookmark{
					{
						URL:    ptrStr("https://github.com/vaguecoder"),
						Title:  "Vague Coder",
						Folder: "Profiles/GitHub",
						ID:     1,
						Parent: 0,
					},
					{
						URL:    ptrStr("https://github.com/random"),
						Title:  "Random",
						Folder: "Profiles/GitHub",
						ID:     2,
						Parent: 0,
					},
				},
			},
			toggles: toggles{
				isExpectedError:      false,
				isFileWriter:         true,
				isErrorAtCSVWriteAll: false,
			},
			expected: []string{
				`https://github.com/vaguecoder,Vague Coder,Profiles/GitHub,1,0`,
				`https://github.com/random,Random,Profiles/GitHub,2,0`,
			},
		},
		{
			name: "Valid-Case-With-Non-File-Writer-Disable-Header",
			fields: fields{
				enableHeader: false,
				filename:     "firefox-bookmarks.csv",
				outputWriter: nonFileWriter,
			},
			args: args{
				bookmarks: []bookmark.Bookmark{
					{
						URL:    ptrStr("https://github.com/vaguecoder"),
						Title:  "Vague Coder",
						Folder: "Profiles/GitHub",
						ID:     1,
						Parent: 0,
					},
					{
						URL:    ptrStr("https://github.com/random"),
						Title:  "Random",
						Folder: "Profiles/GitHub",
						ID:     2,
						Parent: 0,
					},
				},
			},
			expected: []string{
				`https://github.com/vaguecoder,Vague Coder,Profiles/GitHub,1,0`,
				`https://github.com/random,Random,Profiles/GitHub,2,0`,
			},
			toggles: toggles{
				isExpectedError:      false,
				isFileWriter:         false,
				isErrorAtCSVWriteAll: false,
			},
		},
		{
			name: "No-Bookmarks-Disable-Header",
			fields: fields{
				enableHeader: false,
				filename:     "firefox-bookmarks.csv",
				outputWriter: fileWriter,
			},
			args: args{
				bookmarks: []bookmark.Bookmark{},
			},
			toggles: toggles{
				isExpectedError:      false,
				isFileWriter:         true,
				isErrorAtCSVWriteAll: false,
			},
			expected: []string{},
		},
		{
			name: "No-Bookmarks-With-Non-File-Writer-Disable-Header",
			fields: fields{
				enableHeader: false,
				filename:     "firefox-bookmarks.csv",
				outputWriter: nonFileWriter,
			},
			args: args{
				bookmarks: []bookmark.Bookmark{},
			},
			expected: []string{},
			toggles: toggles{
				isExpectedError:      false,
				isFileWriter:         false,
				isErrorAtCSVWriteAll: false,
			},
		},

		//
		{
			name: "Valid-Case-Enable-Header",
			fields: fields{
				enableHeader: true,
				filename:     "firefox-bookmarks.csv",
				outputWriter: fileWriter,
			},
			args: args{
				bookmarks: []bookmark.Bookmark{
					{
						URL:    ptrStr("https://github.com/vaguecoder"),
						Title:  "Vague Coder",
						Folder: "Profiles/GitHub",
						ID:     1,
						Parent: 0,
					},
					{
						URL:    ptrStr("https://github.com/random"),
						Title:  "Random",
						Folder: "Profiles/GitHub",
						ID:     2,
						Parent: 0,
					},
				},
			},
			toggles: toggles{
				isExpectedError:      false,
				isFileWriter:         true,
				isErrorAtCSVWriteAll: false,
			},
			expected: []string{
				`URL,TITLE,FOLDER,ID,PARENT`,
				`---,-----,------,--,------`,
				`https://github.com/vaguecoder,Vague Coder,Profiles/GitHub,1,0`,
				`https://github.com/random,Random,Profiles/GitHub,2,0`,
			},
		},
		{
			name: "Valid-Case-With-Non-File-Writer-Enable-Header",
			fields: fields{
				enableHeader: true,
				filename:     "firefox-bookmarks.csv",
				outputWriter: nonFileWriter,
			},
			args: args{
				bookmarks: []bookmark.Bookmark{
					{
						URL:    ptrStr("https://github.com/vaguecoder"),
						Title:  "Vague Coder",
						Folder: "Profiles/GitHub",
						ID:     1,
						Parent: 0,
					},
					{
						URL:    ptrStr("https://github.com/random"),
						Title:  "Random",
						Folder: "Profiles/GitHub",
						ID:     2,
						Parent: 0,
					},
				},
			},
			expected: []string{
				`URL,TITLE,FOLDER,ID,PARENT`,
				`---,-----,------,--,------`,
				`https://github.com/vaguecoder,Vague Coder,Profiles/GitHub,1,0`,
				`https://github.com/random,Random,Profiles/GitHub,2,0`,
			},
			toggles: toggles{
				isExpectedError:      false,
				isFileWriter:         false,
				isErrorAtCSVWriteAll: false,
			},
		},
		{
			name: "No-Bookmarks-Enable-Header",
			fields: fields{
				enableHeader: true,
				filename:     "firefox-bookmarks.csv",
				outputWriter: fileWriter,
			},
			args: args{
				bookmarks: []bookmark.Bookmark{},
			},
			toggles: toggles{
				isExpectedError:      false,
				isFileWriter:         true,
				isErrorAtCSVWriteAll: false,
			},
			expected: []string{}, // In case of no bookmarks, the bookmarks table is nil w/o header regardless
		},
		{
			name: "No-Bookmarks-With-Non-File-Writer-Enable-Header",
			fields: fields{
				enableHeader: true,
				filename:     "firefox-bookmarks.csv",
				outputWriter: nonFileWriter,
			},
			args: args{
				bookmarks: []bookmark.Bookmark{}, // In case of no bookmarks, the bookmarks table is nil w/o header regardless
			},
			expected: []string{},
			toggles: toggles{
				isExpectedError:      false,
				isFileWriter:         false,
				isErrorAtCSVWriteAll: false,
			},
		},

		//

		{
			name: "Failure-At-CSV-Write-All",
			fields: fields{
				enableHeader: true,
				filename:     "firefox-bookmarks.csv",
				outputWriter: fileWriter,
			},
			args: args{
				bookmarks: []bookmark.Bookmark{
					{
						URL:    ptrStr("https://github.com/vaguecoder"),
						Title:  "Vague Coder",
						Folder: "Profiles/GitHub",
						ID:     1,
						Parent: 0,
					},
					{
						URL:    ptrStr("https://github.com/random"),
						Title:  "Random",
						Folder: "Profiles/GitHub",
						ID:     2,
						Parent: 0,
					},
				},
			},
			expected: []string{
				`URL,TITLE,FOLDER,ID,PARENT`,
				`---,-----,------,--,------`,
				`https://github.com/vaguecoder,Vague Coder,Profiles/GitHub,1,0`,
				`https://github.com/random,Random,Profiles/GitHub,2,0`,
			},
			toggles: toggles{
				isExpectedError:      true,
				isFileWriter:         true,
				isErrorAtCSVWriteAll: true,
			},
		},
		{
			name: "Failure-At-CSV-Write-All-With-Non-File-Writer",
			fields: fields{
				enableHeader: true,
				filename:     "firefox-bookmarks.csv",
				outputWriter: nonFileWriter,
			},
			args: args{
				bookmarks: []bookmark.Bookmark{
					{
						URL:    ptrStr("https://github.com/vaguecoder"),
						Title:  "Vague Coder",
						Folder: "Profiles/GitHub",
						ID:     1,
						Parent: 0,
					},
					{
						URL:    ptrStr("https://github.com/random"),
						Title:  "Random",
						Folder: "Profiles/GitHub",
						ID:     2,
						Parent: 0,
					},
				},
			},
			expected: []string{
				`URL,TITLE,FOLDER,ID,PARENT`,
				`---,-----,------,--,------`,
				`https://github.com/vaguecoder,Vague Coder,Profiles/GitHub,1,0`,
				`https://github.com/random,Random,Profiles/GitHub,2,0`,
			},
			toggles: toggles{
				isExpectedError:      true,
				isFileWriter:         false,
				isErrorAtCSVWriteAll: true,
			},
		},
	}
	for _, testCase = range tests {
		t.Run(testCase.name, func(t *testing.T) {
			expectedOutput = stringSliceToFlatBytes(testCase.expected)

			var (
				lengthOfExpectedOutput       = len(expectedOutput)
				csvWriteAllErr         error = nil
			)

			filename = ""
			if testCase.toggles.isErrorAtCSVWriteAll {
				csvWriteAllErr = fmt.Errorf("some error")
			}

			if testCase.toggles.isFileWriter {
				filename = testCase.fields.filename
				fileWriter.On("Name").Return(testCase.fields.filename).Once()
				fileWriter.On("Write", expectedOutput).Return(lengthOfExpectedOutput, csvWriteAllErr).Once()
			} else {
				nonFileWriter.On("Write", expectedOutput).Return(lengthOfExpectedOutput, csvWriteAllErr).Once()
			}

			encoder = NewEncoder(testCase.fields.outputWriter, testCase.fields.enableHeader)
			err = encoder.Encode(testCase.args.bookmarks)
			if testCase.toggles.isExpectedError {
				assert.Error(t, err, "Expected error from encode")
			} else {
				assert.NoError(t, err, "Unexpected error from encode")
			}

			assert.Equal(t, filename, encoder.Filename(), "Mismatch of filename in writer")
			assert.Equal(t, constants.CSVFormat.String(), encoder.String(), "Mismatch of encoder name string")
		})
	}
}
