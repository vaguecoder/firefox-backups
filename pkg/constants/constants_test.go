package constants

import (
	"testing"
)

func TestConstant_stringer_OutputFormat_String(t *testing.T) {
	tests := []struct {
		name     string
		stringer Constant[OutputFormat]
		want     string
	}{
		{
			name:     "Empty-String",
			stringer: "",
			want:     "",
		},
		{
			name:     "String-Value",
			stringer: "Luffy",
			want:     "Luffy",
		},
		{
			name:     "OutputFormat_JSON-Format",
			stringer: JSONFormat,
			want:     `json`,
		},
		{
			name:     "OutputFormat_YAML-Format",
			stringer: YAMLFormat,
			want:     `yaml`,
		},
		{
			name:     "OutputFormat_CSV-Format",
			stringer: CSVFormat,
			want:     `csv`,
		},
		{
			name:     "OutputFormat_Tabular-Format",
			stringer: TabularFormat,
			want:     `table`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constant[OutputFormat](tt.stringer).String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstant_stringer_Filter_String(t *testing.T) {
	tests := []struct {
		name     string
		stringer Constant[Filter]
		want     string
	}{
		{
			name:     "Empty-String",
			stringer: "",
			want:     "",
		},
		{
			name:     "String-Value",
			stringer: "Luffy",
			want:     "Luffy",
		},
		{
			name:     "Filter_Denormalize-Filter",
			stringer: DenormalizeFilter,
			want:     `denormalize`,
		},
		{
			name:     "Filter_Ignore-Defaults-Filter",
			stringer: IgnoreDefaultsFilter,
			want:     `ignore-defaults`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constant[Filter](tt.stringer).String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstant_stringer_Flag_String(t *testing.T) {
	tests := []struct {
		name     string
		stringer Constant[Flag]
		want     string
	}{
		{
			name:     "Empty-String",
			stringer: "",
			want:     "",
		},
		{
			name:     "String-Value",
			stringer: "Luffy",
			want:     "Luffy",
		},
		{
			name:     "Flag_Input-SQLite-File-Flag",
			stringer: InputSQLiteFileFlag,
			want:     `input-sqlite-file`,
		},
		{
			name:     "Flag_Raw-Flag",
			stringer: RawFlag,
			want:     `raw`,
		},
		{
			name:     "Flag_Ignore-Defaults-Flag",
			stringer: IgnoreDefaultsFlag,
			want:     `ignore-defaults`,
		},
		{
			name:     "Flag_Silent-Flag",
			stringer: SilentFlag,
			want:     `silent`,
		},
		{
			name:     "Flag_Std-Out-Format-Flag",
			stringer: StdOutFormatFlag,
			want:     `stdout-format`,
		},
		{
			name:     "Flag_Denormalize-Flag",
			stringer: DenormalizeFlag,
			want:     `denormalize`,
		},
		{
			name:     "Flag_Output-Files",
			stringer: OutputFiles,
			want:     `output-files`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constant[Flag](tt.stringer).String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
