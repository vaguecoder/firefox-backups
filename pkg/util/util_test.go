package util_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/vaguecoder/firefox-backups/pkg/util"
)

func TestStrWhitespacesCleanup(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Empty-Input",
			input: "",
			want:  "",
		},
		{
			name:  "All-Spaces",
			input: "  \t\n",
			want:  " ",
		},
		{
			name:  "Ignore-Single-Whitespaces",
			input: " Hey! World. ",
			want:  " Hey! World. ",
		},
		{
			name:  "Replace-Multiple-Whitespaces",
			input: "  Hey!  World.  ",
			want:  " Hey! World. ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := util.StrWhitespacesCleanup(tt.input); got != tt.want {
				t.Errorf("StrWhitespacesCleanup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapKeys(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]any
		want []string
	}{
		{
			name: "No-Items",
			m:    map[string]interface{}{},
			want: []string{},
		},
		{
			name: "With-Keys",
			m: map[string]interface{}{
				"Luffy": "King of the pirates",
				"Zoro":  "World's greatest swordsman",
			},
			want: []string{"Luffy", "Zoro"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := util.MapKeys(tt.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToStringers(t *testing.T) {
	type args struct {
		elems []any
	}
	tests := []struct {
		name    string
		args    args
		want    []fmt.Stringer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := util.ToStringers(tt.args.elems)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToStringers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStringers() = %v, want %v", got, tt.want)
			}
		})
	}
}
