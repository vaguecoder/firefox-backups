package tabular

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vaguecoder/firefox-backups/pkg/bookmark"
	"github.com/vaguecoder/firefox-backups/pkg/constants"
	"github.com/vaguecoder/firefox-backups/pkg/mocks"
	"github.com/vaguecoder/firefox-backups/pkg/util"
)

var (
	ptrStr     = util.PtrStr
	whitespace = util.Whitespace
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
		isExpectedError     bool
		isFileWriter        bool
		isErrorAtTableWrite bool
	}

	type testData struct {
		name     string
		fields   fields
		args     args
		toggles  toggles
		expected [][]string
	}

	var (
		err            error
		tests          []testData
		testCase       testData
		encoder        *Encoder
		filename, cell string
		index          int

		fileWriter         = new(mocks.File)
		nonFileWriter      = new(mocks.NonFileWriter)
		tabWidthWhitespace = whitespace(fixedTabWidth) // Highest width cell has the hardcoded tabwidth, i.e., 8 here
		lfChar             = []byte{0xa}
	)

	tests = []testData{
		{
			name: "Valid-Case-Disable-Header",
			fields: fields{
				enableHeader: false,
				filename:     "firefox-bookmarks.txt",
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
				isExpectedError:     false,
				isFileWriter:        true,
				isErrorAtTableWrite: false,
			},
			expected: [][]string{
				{
					`https://github.com/vaguecoder`,
					tabWidthWhitespace,
					`Vague Coder`,
					tabWidthWhitespace,
					`Profiles/GitHub`,
					tabWidthWhitespace,
					`1`,
					tabWidthWhitespace,
					`0`,
					tabWidthWhitespace, // Closing tab added.
				},
				{
					`https://github.com/random`,
					// Width difference of max tab width and previous cell (29 - 25 = 4) is padded to maintain the table format
					whitespace(uint(len(`https://github.com/vaguecoder`) - len(`https://github.com/random`))),
					tabWidthWhitespace,
					`Random`,
					// Width difference of max tab width and previous cell (11 - 6 = 5) is padded to maintain the table format
					whitespace(uint(len(`Vague Coder`) - len(`Random`))),
					tabWidthWhitespace,
					`Profiles/GitHub`,
					tabWidthWhitespace,
					`2`,
					tabWidthWhitespace,
					`0`,
					tabWidthWhitespace, // Closing tab added.
				},
			},
		},
		{
			name: "Valid-Case-With-Non-File-Writer-Disable-Header",
			fields: fields{
				enableHeader: false,
				filename:     "firefox-bookmarks.txt",
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
			expected: [][]string{
				{
					`https://github.com/vaguecoder`,
					tabWidthWhitespace,
					`Vague Coder`,
					tabWidthWhitespace,
					`Profiles/GitHub`,
					tabWidthWhitespace,
					`1`,
					tabWidthWhitespace,
					`0`,
					tabWidthWhitespace, // Closing tab added.
				},
				{
					`https://github.com/random`,
					// Width difference of max tab width and previous cell (29 - 25 = 4) is padded to maintain the table format
					whitespace(uint(len(`https://github.com/vaguecoder`) - len(`https://github.com/random`))),
					tabWidthWhitespace,
					`Random`,
					// Width difference of max tab width and previous cell (11 - 6 = 5) is padded to maintain the table format
					whitespace(uint(len(`Vague Coder`) - len(`Random`))),
					tabWidthWhitespace,
					`Profiles/GitHub`,
					tabWidthWhitespace,
					`2`,
					tabWidthWhitespace,
					`0`,
					tabWidthWhitespace, // Closing tab added.
				},
			},
			toggles: toggles{
				isExpectedError:     false,
				isFileWriter:        false,
				isErrorAtTableWrite: false,
			},
		},
		{
			name: "No-Bookmarks-Disable-Header",
			fields: fields{
				enableHeader: false,
				filename:     "firefox-bookmarks.txt",
				outputWriter: fileWriter,
			},
			args: args{
				bookmarks: []bookmark.Bookmark{},
			},
			toggles: toggles{
				isExpectedError:     false,
				isFileWriter:        true,
				isErrorAtTableWrite: false,
			},
			expected: [][]string{},
		},
		{
			name: "No-Bookmarks-With-Non-File-Writer-Disable-Header",
			fields: fields{
				enableHeader: false,
				filename:     "firefox-bookmarks.txt",
				outputWriter: nonFileWriter,
			},
			args: args{
				bookmarks: []bookmark.Bookmark{},
			},
			expected: [][]string{},
			toggles: toggles{
				isExpectedError:     false,
				isFileWriter:        false,
				isErrorAtTableWrite: false,
			},
		},
		{
			name: "Valid-Case-Enable-Header",
			fields: fields{
				enableHeader: true,
				filename:     "firefox-bookmarks.txt",
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
				isExpectedError:     false,
				isFileWriter:        true,
				isErrorAtTableWrite: false,
			},
			expected: [][]string{
				{
					`URL`,
					// Width difference of max tab width and previous cell (29 - 3 = 26) is padded to maintain the table format.
					// Since the padding is too higher than the tab width (26 > 8), it is split into 2 + 8 + 8 + 8
					whitespace(2),
					tabWidthWhitespace,
					tabWidthWhitespace,
					tabWidthWhitespace,
					tabWidthWhitespace,
					`TITLE`,
					// Width difference of max tab width and previous cell (11 - 5 = 6) is padded to maintain the table format
					whitespace(uint(len(`Vague Coder`) - len(`TITLE`))),
					tabWidthWhitespace,
					`FOLDER`,
					// Width difference of max tab width and previous cell (15 - 6 = 9) is padded to maintain the table format.
					// Since the padding is higher than the tab width (9 > 8), it is split into 1 + 8.
					whitespace(1),
					tabWidthWhitespace,
					tabWidthWhitespace,
					`ID`,
					tabWidthWhitespace,
					`PARENT`,
					tabWidthWhitespace, // Closing tab added.
				},
				{
					`---`,
					// Width difference of max tab width and previous cell (29 - 3 = 26) is padded to maintain the table format.
					// Since the padding is too higher than the tab width (26 > 8), it is split into 2 + 8 + 8 + 8
					whitespace(2),
					tabWidthWhitespace,
					tabWidthWhitespace,
					tabWidthWhitespace,
					tabWidthWhitespace,
					`-----`,
					// Width difference of max tab width and previous cell (11 - 5 = 6) is padded to maintain the table format
					whitespace(uint(len(`Vague Coder`) - len(`-----`))),
					tabWidthWhitespace,
					`------`,
					// Width difference of max tab width and previous cell (15 - 6 = 9) is padded to maintain the table format.
					// Since the padding is higher than the tab width (9 > 8), it is split into 1 + 8.
					whitespace(1),
					tabWidthWhitespace,
					tabWidthWhitespace,
					`--`,
					tabWidthWhitespace,
					`------`,
					tabWidthWhitespace, // Closing tab added.
				},
				{
					`https://github.com/vaguecoder`,
					tabWidthWhitespace,
					`Vague Coder`,
					tabWidthWhitespace,
					`Profiles/GitHub`,
					tabWidthWhitespace,
					`1`,
					// Width difference of max tab width and previous cell (2 - 1 = 1) is padded to maintain the table format
					whitespace(uint(len(`ID`) - len(`1`))),
					tabWidthWhitespace,
					`0`,
					// Width difference of max tab width and previous cell (6 - 1 = 5) is padded to maintain the table format
					whitespace(uint(len(`PARENT`) - len(`0`))),
					tabWidthWhitespace, // Closing tab added.
				},
				{
					`https://github.com/random`,
					// Width difference of max tab width and previous cell (29 - 25 = 4) is padded to maintain the table format
					whitespace(uint(len(`https://github.com/vaguecoder`) - len(`https://github.com/random`))),
					tabWidthWhitespace,
					`Random`,
					// Width difference of max tab width and previous cell (11 - 6 = 5) is padded to maintain the table format
					whitespace(uint(len(`Vague Coder`) - len(`Random`))),
					tabWidthWhitespace,
					`Profiles/GitHub`,
					tabWidthWhitespace,
					`2`,
					// Width difference of max tab width and previous cell (2 - 1 = 1) is padded to maintain the table format
					whitespace(uint(len(`ID`) - len(`2`))),
					tabWidthWhitespace,
					`0`,
					// Width difference of max tab width and previous cell (6 - 1 = 5) is padded to maintain the table format
					whitespace(uint(len(`PARENT`) - len(`0`))),
					tabWidthWhitespace, // Closing tab added.
				},
			},
		},
		{
			name: "Valid-Case-With-Non-File-Writer-Enable-Header",
			fields: fields{
				enableHeader: true,
				filename:     "firefox-bookmarks.txt",
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
			expected: [][]string{
				{
					`URL`,
					// Width difference of max tab width and previous cell (29 - 3 = 26) is padded to maintain the table format.
					// Since the padding is too higher than the tab width (26 > 8), it is split into 2 + 8 + 8 + 8
					whitespace(2),
					tabWidthWhitespace,
					tabWidthWhitespace,
					tabWidthWhitespace,
					tabWidthWhitespace,
					`TITLE`,
					// Width difference of max tab width and previous cell (11 - 5 = 6) is padded to maintain the table format
					whitespace(uint(len(`Vague Coder`) - len(`TITLE`))),
					tabWidthWhitespace,
					`FOLDER`,
					// Width difference of max tab width and previous cell (15 - 6 = 9) is padded to maintain the table format.
					// Since the padding is higher than the tab width (9 > 8), it is split into 1 + 8.
					whitespace(1),
					tabWidthWhitespace,
					tabWidthWhitespace,
					`ID`,
					tabWidthWhitespace,
					`PARENT`,
					tabWidthWhitespace, // Closing tab added.
				},
				{
					`---`,
					// Width difference of max tab width and previous cell (29 - 3 = 26) is padded to maintain the table format.
					// Since the padding is too higher than the tab width (26 > 8), it is split into 2 + 8 + 8 + 8
					whitespace(2),
					tabWidthWhitespace,
					tabWidthWhitespace,
					tabWidthWhitespace,
					tabWidthWhitespace,
					`-----`,
					// Width difference of max tab width and previous cell (11 - 5 = 6) is padded to maintain the table format
					whitespace(uint(len(`Vague Coder`) - len(`-----`))),
					tabWidthWhitespace,
					`------`,
					// Width difference of max tab width and previous cell (15 - 6 = 9) is padded to maintain the table format.
					// Since the padding is higher than the tab width (9 > 8), it is split into 1 + 8.
					whitespace(1),
					tabWidthWhitespace,
					tabWidthWhitespace,
					`--`,
					tabWidthWhitespace,
					`------`,
					tabWidthWhitespace, // Closing tab added.
				},
				{
					`https://github.com/vaguecoder`,
					tabWidthWhitespace,
					`Vague Coder`,
					tabWidthWhitespace,
					`Profiles/GitHub`,
					tabWidthWhitespace,
					`1`,
					// Width difference of max tab width and previous cell (2 - 1 = 1) is padded to maintain the table format
					whitespace(uint(len(`ID`) - len(`1`))),
					tabWidthWhitespace,
					`0`,
					// Width difference of max tab width and previous cell (6 - 1 = 5) is padded to maintain the table format
					whitespace(uint(len(`PARENT`) - len(`0`))),
					tabWidthWhitespace, // Closing tab added.
				},
				{
					`https://github.com/random`,
					// Width difference of max tab width and previous cell (29 - 25 = 4) is padded to maintain the table format
					whitespace(uint(len(`https://github.com/vaguecoder`) - len(`https://github.com/random`))),
					tabWidthWhitespace,
					`Random`,
					// Width difference of max tab width and previous cell (11 - 6 = 5) is padded to maintain the table format
					whitespace(uint(len(`Vague Coder`) - len(`Random`))),
					tabWidthWhitespace,
					`Profiles/GitHub`,
					tabWidthWhitespace,
					`2`,
					// Width difference of max tab width and previous cell (2 - 1 = 1) is padded to maintain the table format
					whitespace(uint(len(`ID`) - len(`2`))),
					tabWidthWhitespace,
					`0`,
					// Width difference of max tab width and previous cell (6 - 1 = 5) is padded to maintain the table format
					whitespace(uint(len(`PARENT`) - len(`0`))),
					tabWidthWhitespace, // Closing tab added.
				},
			},
			toggles: toggles{
				isExpectedError:     false,
				isFileWriter:        false,
				isErrorAtTableWrite: false,
			},
		},
		{
			name: "No-Bookmarks-Enable-Header",
			fields: fields{
				enableHeader: true,
				filename:     "firefox-bookmarks.txt",
				outputWriter: fileWriter,
			},
			args: args{
				bookmarks: []bookmark.Bookmark{},
			},
			toggles: toggles{
				isExpectedError:     false,
				isFileWriter:        true,
				isErrorAtTableWrite: false,
			},
			expected: [][]string{}, // In case of no bookmarks, the bookmarks table is nil w/o header regardless
		},
		{
			name: "No-Bookmarks-With-Non-File-Writer-Enable-Header",
			fields: fields{
				enableHeader: true,
				filename:     "firefox-bookmarks.txt",
				outputWriter: nonFileWriter,
			},
			args: args{
				bookmarks: []bookmark.Bookmark{}, // In case of no bookmarks, the bookmarks table is nil w/o header regardless
			},
			expected: [][]string{},
			toggles: toggles{
				isExpectedError:     false,
				isFileWriter:        false,
				isErrorAtTableWrite: false,
			},
		},
		{
			name: "Failure-At-Table-Write",
			fields: fields{
				enableHeader: true,
				filename:     "firefox-bookmarks.txt",
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
			expected: [][]string{
				{
					`URL`,
					// Width difference of max tab width and previous cell (29 - 3 = 26) is padded to maintain the table format.
					// Since the padding is too higher than the tab width (26 > 8), it is split into 2 + 8 + 8 + 8
					whitespace(2),
					tabWidthWhitespace,
					tabWidthWhitespace,
					tabWidthWhitespace,
					tabWidthWhitespace,
					`TITLE`,
					// Width difference of max tab width and previous cell (11 - 5 = 6) is padded to maintain the table format
					whitespace(uint(len(`Vague Coder`) - len(`TITLE`))),
					tabWidthWhitespace,
					`FOLDER`,
					// Width difference of max tab width and previous cell (15 - 6 = 9) is padded to maintain the table format.
					// Since the padding is higher than the tab width (9 > 8), it is split into 1 + 8.
					whitespace(1),
					tabWidthWhitespace,
					tabWidthWhitespace,
					`ID`,
					tabWidthWhitespace,
					`PARENT`,
					tabWidthWhitespace, // Closing tab added.
				},
				{
					`---`,
					// Width difference of max tab width and previous cell (29 - 3 = 26) is padded to maintain the table format.
					// Since the padding is too higher than the tab width (26 > 8), it is split into 2 + 8 + 8 + 8
					whitespace(2),
					tabWidthWhitespace,
					tabWidthWhitespace,
					tabWidthWhitespace,
					tabWidthWhitespace,
					`-----`,
					// Width difference of max tab width and previous cell (11 - 5 = 6) is padded to maintain the table format
					whitespace(uint(len(`Vague Coder`) - len(`-----`))),
					tabWidthWhitespace,
					`------`,
					// Width difference of max tab width and previous cell (15 - 6 = 9) is padded to maintain the table format.
					// Since the padding is higher than the tab width (9 > 8), it is split into 1 + 8.
					whitespace(1),
					tabWidthWhitespace,
					tabWidthWhitespace,
					`--`,
					tabWidthWhitespace,
					`------`,
					tabWidthWhitespace, // Closing tab added.
				},
				{
					`https://github.com/vaguecoder`,
					tabWidthWhitespace,
					`Vague Coder`,
					tabWidthWhitespace,
					`Profiles/GitHub`,
					tabWidthWhitespace,
					`1`,
					// Width difference of max tab width and previous cell (2 - 1 = 1) is padded to maintain the table format
					whitespace(uint(len(`ID`) - len(`1`))),
					tabWidthWhitespace,
					`0`,
					// Width difference of max tab width and previous cell (6 - 1 = 5) is padded to maintain the table format
					whitespace(uint(len(`PARENT`) - len(`0`))),
					tabWidthWhitespace, // Closing tab added.
				},
				{
					`https://github.com/random`,
					// Width difference of max tab width and previous cell (29 - 25 = 4) is padded to maintain the table format
					whitespace(uint(len(`https://github.com/vaguecoder`) - len(`https://github.com/random`))),
					tabWidthWhitespace,
					`Random`,
					// Width difference of max tab width and previous cell (11 - 6 = 5) is padded to maintain the table format
					whitespace(uint(len(`Vague Coder`) - len(`Random`))),
					tabWidthWhitespace,
					`Profiles/GitHub`,
					tabWidthWhitespace,
					`2`,
					// Width difference of max tab width and previous cell (2 - 1 = 1) is padded to maintain the table format
					whitespace(uint(len(`ID`) - len(`2`))),
					tabWidthWhitespace,
					`0`,
					// Width difference of max tab width and previous cell (6 - 1 = 5) is padded to maintain the table format
					whitespace(uint(len(`PARENT`) - len(`0`))),
					tabWidthWhitespace, // Closing tab added.
				},
			},
			toggles: toggles{
				isExpectedError:     true,
				isFileWriter:        true,
				isErrorAtTableWrite: true,
			},
		},
		{
			name: "Failure-At-Table-Write-With-Non-File-Writer",
			fields: fields{
				enableHeader: true,
				filename:     "firefox-bookmarks.txt",
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
			expected: [][]string{
				{
					`URL`,
					// Width difference of max tab width and previous cell (29 - 3 = 26) is padded to maintain the table format.
					// Since the padding is too higher than the tab width (26 > 8), it is split into 2 + 8 + 8 + 8
					whitespace(2),
					tabWidthWhitespace,
					tabWidthWhitespace,
					tabWidthWhitespace,
					tabWidthWhitespace,
					`TITLE`,
					// Width difference of max tab width and previous cell (11 - 5 = 6) is padded to maintain the table format
					whitespace(uint(len(`Vague Coder`) - len(`TITLE`))),
					tabWidthWhitespace,
					`FOLDER`,
					// Width difference of max tab width and previous cell (15 - 6 = 9) is padded to maintain the table format.
					// Since the padding is higher than the tab width (9 > 8), it is split into 1 + 8.
					whitespace(1),
					tabWidthWhitespace,
					tabWidthWhitespace,
					`ID`,
					tabWidthWhitespace,
					`PARENT`,
					tabWidthWhitespace, // Closing tab added.
				},
				{
					`---`,
					// Width difference of max tab width and previous cell (29 - 3 = 26) is padded to maintain the table format.
					// Since the padding is too higher than the tab width (26 > 8), it is split into 2 + 8 + 8 + 8
					whitespace(2),
					tabWidthWhitespace,
					tabWidthWhitespace,
					tabWidthWhitespace,
					tabWidthWhitespace,
					`-----`,
					// Width difference of max tab width and previous cell (11 - 5 = 6) is padded to maintain the table format
					whitespace(uint(len(`Vague Coder`) - len(`-----`))),
					tabWidthWhitespace,
					`------`,
					// Width difference of max tab width and previous cell (15 - 6 = 9) is padded to maintain the table format.
					// Since the padding is higher than the tab width (9 > 8), it is split into 1 + 8.
					whitespace(1),
					tabWidthWhitespace,
					tabWidthWhitespace,
					`--`,
					tabWidthWhitespace,
					`------`,
					tabWidthWhitespace, // Closing tab added.
				},
				{
					`https://github.com/vaguecoder`,
					tabWidthWhitespace,
					`Vague Coder`,
					tabWidthWhitespace,
					`Profiles/GitHub`,
					tabWidthWhitespace,
					`1`,
					// Width difference of max tab width and previous cell (2 - 1 = 1) is padded to maintain the table format
					whitespace(uint(len(`ID`) - len(`1`))),
					tabWidthWhitespace,
					`0`,
					// Width difference of max tab width and previous cell (6 - 1 = 5) is padded to maintain the table format
					whitespace(uint(len(`PARENT`) - len(`0`))),
					tabWidthWhitespace, // Closing tab added.
				},
				{
					`https://github.com/random`,
					// Width difference of max tab width and previous cell (29 - 25 = 4) is padded to maintain the table format
					whitespace(uint(len(`https://github.com/vaguecoder`) - len(`https://github.com/random`))),
					tabWidthWhitespace,
					`Random`,
					// Width difference of max tab width and previous cell (11 - 6 = 5) is padded to maintain the table format
					whitespace(uint(len(`Vague Coder`) - len(`Random`))),
					tabWidthWhitespace,
					`Profiles/GitHub`,
					tabWidthWhitespace,
					`2`,
					// Width difference of max tab width and previous cell (2 - 1 = 1) is padded to maintain the table format
					whitespace(uint(len(`ID`) - len(`2`))),
					tabWidthWhitespace,
					`0`,
					// Width difference of max tab width and previous cell (6 - 1 = 5) is padded to maintain the table format
					whitespace(uint(len(`PARENT`) - len(`0`))),
					tabWidthWhitespace, // Closing tab added.
				},
			},
			toggles: toggles{
				isExpectedError:     true,
				isFileWriter:        false,
				isErrorAtTableWrite: true,
			},
		},
	}
	for _, testCase = range tests {
		t.Run(testCase.name, func(t *testing.T) {

			var writeErr error = fmt.Errorf("some error")

			filename = ""
			if testCase.toggles.isFileWriter {

				filename = testCase.fields.filename
				fileWriter.On("Name").Return(testCase.fields.filename).Once()

				if len(testCase.expected) == 0 {
					// To flush, write the enpty byte to tabwriter
					fileWriter.On("Write", []byte(nil)).Return(0, nil).Once()
				} else {
					if testCase.toggles.isErrorAtTableWrite {
						fileWriter.On("Write", []byte(testCase.expected[0][0])).Return(len(testCase.expected[0][0]), writeErr).Once()
					} else {
						for _, record := range testCase.expected {
							for index, cell = range record {
								fileWriter.On("Write", []byte(cell)).Return(len(cell), nil).Once()
							}

							if index == len(record)-1 {
								// Add line feed (LF) character at the end records
								fileWriter.On("Write", lfChar).Return(len(lfChar), nil).Once()
							}
						}

						// To flush, write the enpty byte to tabwriter.
						//
						// In case of error at actual Write(), i.e., non flush Write(), the error is
						// returned while flushing, regardless of what error is returned in flush Write().
						fileWriter.On("Write", []byte{}).Return(0, nil).Once()
					}
				}

			} else {
				if len(testCase.expected) == 0 {
					// To flush, write the enpty byte to tabwriter
					nonFileWriter.On("Write", []byte(nil)).Return(0, nil).Once()
				} else {
					if testCase.toggles.isErrorAtTableWrite {
						nonFileWriter.On("Write", []byte(testCase.expected[0][0])).Return(len(testCase.expected[0][0]), writeErr).Once()
					} else {
						for _, record := range testCase.expected {
							for index, cell = range record {
								nonFileWriter.On("Write", []byte(cell)).Return(len(cell), nil).Once()
							}

							if index == len(record)-1 {
								// Add line feed (LF) character at the end records
								nonFileWriter.On("Write", lfChar).Return(len(lfChar), nil).Once()
							}
						}

						// To flush, write the enpty byte to tabwriter.
						//
						// In case of error at actual Write(), i.e., non flush Write(), the error is
						// returned while flushing, regardless of what error is returned in flush Write().
						nonFileWriter.On("Write", []byte{}).Return(0, nil).Once()
					}
				}
			}

			encoder = NewEncoder(testCase.fields.outputWriter, testCase.fields.enableHeader)

			err = encoder.Encode(testCase.args.bookmarks)

			if testCase.toggles.isExpectedError {
				assert.Error(t, err, "Expected error from encode")
			} else {
				assert.NoError(t, err, "Unexpected error from encode")
			}

			assert.Equal(t, filename, encoder.Filename(), "Mismatch of filename in writer")
			assert.Equal(t, constants.TabularFormat.String(), encoder.String(), "Mismatch of encoder name string")

			mock.AssertExpectationsForObjects(t, fileWriter, nonFileWriter)
		})
	}
}
