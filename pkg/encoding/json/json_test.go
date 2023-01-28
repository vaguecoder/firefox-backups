package json

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
		filename     string
		outputWriter io.Writer
	}

	type args struct {
		bookmarks []bookmark.Bookmark
	}

	type toggles struct {
		isExpectedError    bool
		isFileWriter       bool
		isErrorAtJSONWrite bool
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
			name: "Valid-Case",
			fields: fields{
				filename:     "firefox-bookmarks.json",
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
				isExpectedError:    false,
				isFileWriter:       true,
				isErrorAtJSONWrite: false,
			},
			expected: []string{
				`[`,
				`	{`,
				`		"url": "https://github.com/vaguecoder",`,
				`		"title": "Vague Coder",`,
				`		"folder": "Profiles/GitHub",`,
				`		"id": 1,`,
				`		"parent": 0`,
				`	},`,
				`	{`,
				`		"url": "https://github.com/random",`,
				`		"title": "Random",`,
				`		"folder": "Profiles/GitHub",`,
				`		"id": 2,`,
				`		"parent": 0`,
				`	}`,
				`]`,
			},
		},
		{
			name: "Valid-Case-With-Non-File-Writer",
			fields: fields{
				filename:     "firefox-bookmarks.json",
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
				`[`,
				`	{`,
				`		"url": "https://github.com/vaguecoder",`,
				`		"title": "Vague Coder",`,
				`		"folder": "Profiles/GitHub",`,
				`		"id": 1,`,
				`		"parent": 0`,
				`	},`,
				`	{`,
				`		"url": "https://github.com/random",`,
				`		"title": "Random",`,
				`		"folder": "Profiles/GitHub",`,
				`		"id": 2,`,
				`		"parent": 0`,
				`	}`,
				`]`,
			},
			toggles: toggles{
				isExpectedError:    false,
				isFileWriter:       false,
				isErrorAtJSONWrite: false,
			},
		},
		{
			name: "No-Bookmarks",
			fields: fields{
				filename:     "firefox-bookmarks.json",
				outputWriter: fileWriter,
			},
			args: args{
				bookmarks: []bookmark.Bookmark{},
			},
			toggles: toggles{
				isExpectedError:    false,
				isFileWriter:       true,
				isErrorAtJSONWrite: false,
			},
			expected: []string{`[]`},
		},
		{
			name: "No-Bookmarks-With-Non-File-Writer",
			fields: fields{
				filename:     "firefox-bookmarks.json",
				outputWriter: nonFileWriter,
			},
			args: args{
				bookmarks: []bookmark.Bookmark{},
			},
			expected: []string{`[]`},
			toggles: toggles{
				isExpectedError:    false,
				isFileWriter:       false,
				isErrorAtJSONWrite: false,
			},
		},
		{
			name: "Failure-At-JSON-Write-All",
			fields: fields{
				filename:     "firefox-bookmarks.json",
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
				`[`,
				`	{`,
				`		"url": "https://github.com/vaguecoder",`,
				`		"title": "Vague Coder",`,
				`		"folder": "Profiles/GitHub",`,
				`		"id": 1,`,
				`		"parent": 0`,
				`	},`,
				`	{`,
				`		"url": "https://github.com/random",`,
				`		"title": "Random",`,
				`		"folder": "Profiles/GitHub",`,
				`		"id": 2,`,
				`		"parent": 0`,
				`	}`,
				`]`,
			},
			toggles: toggles{
				isExpectedError:    true,
				isFileWriter:       true,
				isErrorAtJSONWrite: true,
			},
		},
		{
			name: "Failure-At-JSON-Write-All-With-Non-File-Writer",
			fields: fields{
				filename:     "firefox-bookmarks.json",
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
				`[`,
				`	{`,
				`		"url": "https://github.com/vaguecoder",`,
				`		"title": "Vague Coder",`,
				`		"folder": "Profiles/GitHub",`,
				`		"id": 1,`,
				`		"parent": 0`,
				`	},`,
				`	{`,
				`		"url": "https://github.com/random",`,
				`		"title": "Random",`,
				`		"folder": "Profiles/GitHub",`,
				`		"id": 2,`,
				`		"parent": 0`,
				`	}`,
				`]`,
			},
			toggles: toggles{
				isExpectedError:    true,
				isFileWriter:       false,
				isErrorAtJSONWrite: true,
			},
		},
	}
	for _, testCase = range tests {
		t.Run(testCase.name, func(t *testing.T) {
			expectedOutput = stringSliceToFlatBytes(testCase.expected)

			var (
				lengthOfExpectedOutput       = len(expectedOutput)
				jsonWriteAllErr        error = nil
			)

			if testCase.toggles.isErrorAtJSONWrite {
				jsonWriteAllErr = fmt.Errorf("some error")
			}

			filename = ""
			if testCase.toggles.isFileWriter {
				filename = testCase.fields.filename
				fileWriter.On("Name").Return(testCase.fields.filename).Once()
				fileWriter.On("Write", expectedOutput).Return(lengthOfExpectedOutput, jsonWriteAllErr).Once()
			} else {
				nonFileWriter.On("Write", expectedOutput).Return(lengthOfExpectedOutput, jsonWriteAllErr).Once()
			}

			encoder = NewEncoder(testCase.fields.outputWriter)

			err = encoder.Encode(testCase.args.bookmarks)

			if testCase.toggles.isExpectedError {
				assert.Error(t, err, "Expected error from encode")
			} else {
				assert.NoError(t, err, "Unexpected error from encode")
			}

			assert.Equal(t, filename, encoder.Filename(), "Mismatch of filename in writer")
			assert.Equal(t, constants.JSONFormat.String(), encoder.String(), "Mismatch of encoder name string")
		})
	}
}
