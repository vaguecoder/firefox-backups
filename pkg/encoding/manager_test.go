package encoding

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vaguecoder/firefox-backups/pkg/bookmark"
	"github.com/vaguecoder/firefox-backups/pkg/constants"
	"github.com/vaguecoder/firefox-backups/pkg/mocks"
	"github.com/vaguecoder/firefox-backups/pkg/util"
)

var ptrStr = util.PtrStr

func TestEncodingManager_Write(t *testing.T) {
	type encoderMap map[constants.Constant[constants.OutputFormat]]*mocks.Encoder

	type fields struct {
		encoders  encoderMap
		bookmarks []bookmark.Bookmark
	}

	type toggles struct {
		isAllEncodeErr    bool
		isSingleEncodeErr bool
	}

	type testData struct {
		name                 string
		fields               fields
		toggles              toggles
		indexOfEncodeFailure int
		wantErr              bool
	}

	var (
		err            error
		someErr        error = fmt.Errorf("some error")
		encoderManager *EncodingManager
		encoder        *mocks.Encoder
		encoderName    constants.Constant[constants.OutputFormat]
		testCase       testData
		encoderIndex   int

		ctx                   = context.Background()
		csvEncoder            = new(mocks.Encoder)
		jsonEncoder           = new(mocks.Encoder)
		tableEncoder          = new(mocks.Encoder)
		yamlEncoder           = new(mocks.Encoder)
		encoderNameToFilename = map[constants.Constant[constants.OutputFormat]]string{
			constants.CSVFormat:     `firefox-bookmarks.csv`,
			constants.JSONFormat:    `firefox-bookmarks.json`,
			constants.TabularFormat: `firefox-bookmarks.txt`,
			constants.YAMLFormat:    `firefox-bookmarks.yaml`,
		}
	)

	tests := []testData{
		{
			name: "Empty-Bookmarks",
			fields: fields{
				encoders: encoderMap{
					constants.CSVFormat:     csvEncoder,
					constants.JSONFormat:    jsonEncoder,
					constants.TabularFormat: tableEncoder,
					constants.YAMLFormat:    yamlEncoder,
				},
				bookmarks: []bookmark.Bookmark{},
			},
			toggles: toggles{
				isAllEncodeErr:    false,
				isSingleEncodeErr: false,
			},
			indexOfEncodeFailure: -1,
			wantErr:              false,
		},
		{
			name: "No-Encoders",
			fields: fields{
				encoders: encoderMap{},
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
				isAllEncodeErr:    false,
				isSingleEncodeErr: false,
			},
			indexOfEncodeFailure: -1,
			wantErr:              false,
		},
		{
			name: "Failure-Bookmarks-Not-Added",
			fields: fields{
				encoders: encoderMap{
					constants.CSVFormat:     csvEncoder,
					constants.JSONFormat:    jsonEncoder,
					constants.TabularFormat: tableEncoder,
					constants.YAMLFormat:    yamlEncoder,
				},
				bookmarks: nil,
			},
			toggles: toggles{
				isAllEncodeErr:    false,
				isSingleEncodeErr: false,
			},
			indexOfEncodeFailure: -1,
			wantErr:              true,
		},
		{
			name: "Valid-Case",
			fields: fields{
				encoders: encoderMap{
					constants.CSVFormat:     csvEncoder,
					constants.JSONFormat:    jsonEncoder,
					constants.TabularFormat: tableEncoder,
					constants.YAMLFormat:    yamlEncoder,
				},
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
				isAllEncodeErr:    false,
				isSingleEncodeErr: false,
			},
			indexOfEncodeFailure: -1,
			wantErr:              false,
		},
		{
			name: "Nil-Encoder-In-Start",
			fields: fields{
				encoders: encoderMap{
					constants.CSVFormat:     nil,
					constants.JSONFormat:    jsonEncoder,
					constants.TabularFormat: tableEncoder,
					constants.YAMLFormat:    yamlEncoder,
				},
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
				isAllEncodeErr:    false,
				isSingleEncodeErr: false,
			},
			indexOfEncodeFailure: -1,
			wantErr:              false,
		},
		{
			name: "Nil-Encoder-In-The-End",
			fields: fields{
				encoders: encoderMap{
					constants.CSVFormat:     csvEncoder,
					constants.JSONFormat:    jsonEncoder,
					constants.TabularFormat: tableEncoder,
					constants.YAMLFormat:    nil,
				},
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
				isAllEncodeErr:    false,
				isSingleEncodeErr: false,
			},
			indexOfEncodeFailure: -1,
			wantErr:              false,
		},
		{
			name: "Nil-Encoder-In-The-Middle",
			fields: fields{
				encoders: encoderMap{
					constants.CSVFormat:     csvEncoder,
					constants.JSONFormat:    nil,
					constants.TabularFormat: nil,
					constants.YAMLFormat:    yamlEncoder,
				},
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
				isAllEncodeErr:    false,
				isSingleEncodeErr: false,
			},
			indexOfEncodeFailure: -1,
			wantErr:              false,
		},
		{
			name: "All-Nil-Encoders",
			fields: fields{
				encoders: encoderMap{
					constants.CSVFormat:     nil,
					constants.JSONFormat:    nil,
					constants.TabularFormat: nil,
					constants.YAMLFormat:    nil,
				},
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
				isAllEncodeErr:    false,
				isSingleEncodeErr: false,
			},
			indexOfEncodeFailure: -1,
			wantErr:              false,
		},
		{
			name: "Failure-At-Encode",
			fields: fields{
				encoders: encoderMap{
					constants.CSVFormat:     csvEncoder,
					constants.JSONFormat:    jsonEncoder,
					constants.TabularFormat: tableEncoder,
					constants.YAMLFormat:    yamlEncoder,
				},
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
				isAllEncodeErr:    true,
				isSingleEncodeErr: false,
			},
			indexOfEncodeFailure: -1,
			wantErr:              true,
		},
		{
			name: "Failure-At-First-Encode",
			fields: fields{
				encoders: encoderMap{
					constants.CSVFormat:     csvEncoder,
					constants.JSONFormat:    jsonEncoder,
					constants.TabularFormat: tableEncoder,
					constants.YAMLFormat:    yamlEncoder,
				},
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
				isAllEncodeErr:    false,
				isSingleEncodeErr: true,
			},
			indexOfEncodeFailure: 0,
			wantErr:              true,
		},
		{
			name: "Failure-At-Middle-Encode",
			fields: fields{
				encoders: encoderMap{
					constants.CSVFormat:     csvEncoder,
					constants.JSONFormat:    jsonEncoder,
					constants.TabularFormat: tableEncoder,
					constants.YAMLFormat:    yamlEncoder,
				},
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
				isAllEncodeErr:    false,
				isSingleEncodeErr: true,
			},
			indexOfEncodeFailure: 1,
			wantErr:              true,
		},
		{
			name: "Failure-At-Last-Encode",
			fields: fields{
				encoders: encoderMap{
					constants.CSVFormat:     csvEncoder,
					constants.JSONFormat:    jsonEncoder,
					constants.TabularFormat: tableEncoder,
					constants.YAMLFormat:    yamlEncoder,
				},
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
				isAllEncodeErr:    false,
				isSingleEncodeErr: true,
			},
			indexOfEncodeFailure: 3,
			wantErr:              true,
		},
	}
	for _, testCase = range tests {
		t.Run(testCase.name, func(t *testing.T) {
			encoderManager = NewEncoderManager(ctx).Bookmarks(testCase.fields.bookmarks)

			encoderIndex = -1
			for encoderName, encoder = range testCase.fields.encoders {
				encoderIndex++

				encoderManager = encoderManager.Encoder(encoder)

				if encoder != nil {
					if testCase.toggles.isAllEncodeErr || (testCase.toggles.isSingleEncodeErr && testCase.indexOfEncodeFailure == encoderIndex) {
						encoder.On("Encode", testCase.fields.bookmarks).Return(someErr).Once()
					} else {
						encoder.On("Encode", testCase.fields.bookmarks).Return(nil).Once()
					}

					encoder.On("Filename").Return(encoderNameToFilename[encoderName]).Once()
					encoder.On("String").Return(encoderName.String()).Once()
				}
			}

			err = encoderManager.Write()

			if testCase.wantErr {
				assert.Error(t, err, "Expected error from encoder write")
			} else {
				assert.NoError(t, err, "Unexpected error from encoder write")
			}
		})
	}
}
