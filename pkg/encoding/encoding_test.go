package encoding

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vaguecoder/firefox-backups/pkg/constants"
	"github.com/vaguecoder/firefox-backups/pkg/util"
)

var joinStringers = util.JoinAsString[fmt.Stringer]

func Test_encoderNames_String(t *testing.T) {
	tests := []struct {
		name      string
		stringers []fmt.Stringer
		expected  string
	}{
		{
			name:      "Empty-Input",
			stringers: []fmt.Stringer{},
			expected:  "",
		},
		{
			name: "Single-Input",
			stringers: []fmt.Stringer{
				constants.JSONFormat,
			},
			expected: constants.JSONFormat.String(),
		},
		{
			name: "Actual-All-Encoders-Ordered",
			stringers: []fmt.Stringer{
				constants.CSVFormat,
				constants.JSONFormat,
				constants.TabularFormat,
				constants.YAMLFormat,
			},
			expected: joinStringers(
				[]fmt.Stringer{
					constants.CSVFormat,
					constants.JSONFormat,
					constants.TabularFormat,
					constants.YAMLFormat,
				},
				encoderNamesDelimiter,
			),
		},
		{
			name: "Actual-All-Encoders-Unordered",
			stringers: []fmt.Stringer{
				constants.JSONFormat,
				constants.YAMLFormat,
				constants.CSVFormat,
				constants.TabularFormat,
			},
			expected: joinStringers(
				[]fmt.Stringer{
					constants.CSVFormat,
					constants.JSONFormat,
					constants.TabularFormat,
					constants.YAMLFormat,
				},
				encoderNamesDelimiter,
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var encoders encoderNames

			for _, outputFormat := range tt.stringers {
				encoders = append(encoders, ToEncoder(outputFormat))
			}

			assert.Equal(t, tt.expected, encoders.String(), "Mismatch of encoders formatted as string")
		})
	}
}
