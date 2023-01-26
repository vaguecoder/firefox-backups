package bookmark

import (
	"reflect"
	"strings"
	"testing"

	"github.com/vaguecoder/firefox-backups/pkg/util"
)

var whitespaceCharacters = strings.Repeat(` `, 3)

func TestBookmarksTable(t *testing.T) {
	var (
		ptrStr = util.PtrStr
	)

	type args struct {
		bookmarks    []Bookmark
		enableHeader bool
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Empty-Bookmarks-Header-Disabled",
			args: args{
				bookmarks:    []Bookmark{},
				enableHeader: false,
			},
			want: [][]string(nil),
		},
		{
			name: "Empty-Bookmarks-Header-Enabled",
			args: args{
				bookmarks:    []Bookmark{},
				enableHeader: true,
			},
			want: [][]string(nil),
		},
		{
			name: "Blank-Record-Header-Disabled",
			args: args{
				bookmarks: []Bookmark{
					{
						URL:    nil,
						Title:  "",
						Folder: "",
						ID:     0,
						Parent: 0,
					},
				},
				enableHeader: false,
			},
			want: [][]string{
				{
					"",
					"",
					"",
					"0",
					"0",
				},
			},
		},
		{
			name: "Blank-Record-Header-Enabled",
			args: args{
				bookmarks: []Bookmark{
					{
						URL:    nil,
						Title:  "",
						Folder: "",
						ID:     0,
						Parent: 0,
					},
				},
				enableHeader: true,
			},
			want: [][]string{
				{
					"URL",
					"TITLE",
					"FOLDER",
					"ID",
					"PARENT",
				},
				{
					"---",
					"-----",
					"------",
					"--",
					"------",
				},
				{
					"",
					"",
					"",
					"0",
					"0",
				},
			},
		},
		{
			name: "Single-Record-Header-Disabled",
			args: args{
				bookmarks: []Bookmark{
					{
						URL:    ptrStr("https://github.com/vaguecoder"),
						Title:  "Vague Coder",
						Folder: "Profiles/GitHub",
						ID:     1,
						Parent: 0,
					},
				},
				enableHeader: false,
			},
			want: [][]string{
				{
					"https://github.com/vaguecoder",
					"Vague Coder",
					"Profiles/GitHub",
					"1",
					"0",
				},
			},
		},
		{
			name: "Single-Record-Header-Enabled",
			args: args{
				bookmarks: []Bookmark{
					{
						URL:    ptrStr("https://github.com/vaguecoder"),
						Title:  "Vague Coder",
						Folder: "Profiles/GitHub",
						ID:     1,
						Parent: 0,
					},
				},
				enableHeader: true,
			},
			want: [][]string{
				{
					"URL",
					"TITLE",
					"FOLDER",
					"ID",
					"PARENT",
				},
				{
					"---",
					"-----",
					"------",
					"--",
					"------",
				},
				{
					"https://github.com/vaguecoder",
					"Vague Coder",
					"Profiles/GitHub",
					"1",
					"0",
				},
			},
		},
		{
			name: "Multiple-Records-Header-Disabled",
			args: args{
				bookmarks: []Bookmark{
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
				enableHeader: false,
			},
			want: [][]string{
				{
					"https://github.com/vaguecoder",
					"Vague Coder",
					"Profiles/GitHub",
					"1",
					"0",
				},
				{
					"https://github.com/random",
					"Random",
					"Profiles/GitHub",
					"2",
					"0",
				},
			},
		},
		{
			name: "Multiple-Records-Header-Enabled",
			args: args{
				bookmarks: []Bookmark{
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
				enableHeader: true,
			},
			want: [][]string{
				{
					"URL",
					"TITLE",
					"FOLDER",
					"ID",
					"PARENT",
				},
				{
					"---",
					"-----",
					"------",
					"--",
					"------",
				},
				{
					"https://github.com/vaguecoder",
					"Vague Coder",
					"Profiles/GitHub",
					"1",
					"0",
				},
				{
					"https://github.com/random",
					"Random",
					"Profiles/GitHub",
					"2",
					"0",
				},
			},
		},
		{
			name: "Multiple-Records-Blank-First-Record-Header-Disabled",
			args: args{
				bookmarks: []Bookmark{
					{
						URL:    nil,
						Title:  "",
						Folder: "",
						ID:     0,
						Parent: 0,
					},
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
				enableHeader: false,
			},
			want: [][]string{
				{
					"",
					"",
					"",
					"0",
					"0",
				},
				{
					"https://github.com/vaguecoder",
					"Vague Coder",
					"Profiles/GitHub",
					"1",
					"0",
				},
				{
					"https://github.com/random",
					"Random",
					"Profiles/GitHub",
					"2",
					"0",
				},
			},
		},
		{
			name: "Multiple-Records-Blank-First-Record-Header-Enabled",
			args: args{
				bookmarks: []Bookmark{
					{
						URL:    nil,
						Title:  "",
						Folder: "",
						ID:     0,
						Parent: 0,
					},
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
				enableHeader: true,
			},
			want: [][]string{
				{
					"URL",
					"TITLE",
					"FOLDER",
					"ID",
					"PARENT",
				},
				{
					"---",
					"-----",
					"------",
					"--",
					"------",
				},
				{
					"",
					"",
					"",
					"0",
					"0",
				},
				{
					"https://github.com/vaguecoder",
					"Vague Coder",
					"Profiles/GitHub",
					"1",
					"0",
				},
				{
					"https://github.com/random",
					"Random",
					"Profiles/GitHub",
					"2",
					"0",
				},
			},
		},
		{
			name: "Multiple-Records-Blank-Middle-Record-Header-Disabled",
			args: args{
				bookmarks: []Bookmark{
					{
						URL:    ptrStr("https://github.com/vaguecoder"),
						Title:  "Vague Coder",
						Folder: "Profiles/GitHub",
						ID:     1,
						Parent: 0,
					},
					{
						URL:    nil,
						Title:  "",
						Folder: "",
						ID:     0,
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
				enableHeader: false,
			},
			want: [][]string{
				{
					"https://github.com/vaguecoder",
					"Vague Coder",
					"Profiles/GitHub",
					"1",
					"0",
				},
				{
					"",
					"",
					"",
					"0",
					"0",
				},
				{
					"https://github.com/random",
					"Random",
					"Profiles/GitHub",
					"2",
					"0",
				},
			},
		},
		{
			name: "Multiple-Records-Blank-Middle-Record-Header-Enabled",
			args: args{
				bookmarks: []Bookmark{
					{
						URL:    ptrStr("https://github.com/vaguecoder"),
						Title:  "Vague Coder",
						Folder: "Profiles/GitHub",
						ID:     1,
						Parent: 0,
					},
					{
						URL:    nil,
						Title:  "",
						Folder: "",
						ID:     0,
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
				enableHeader: true,
			},
			want: [][]string{
				{
					"URL",
					"TITLE",
					"FOLDER",
					"ID",
					"PARENT",
				},
				{
					"---",
					"-----",
					"------",
					"--",
					"------",
				},
				{
					"https://github.com/vaguecoder",
					"Vague Coder",
					"Profiles/GitHub",
					"1",
					"0",
				},
				{
					"",
					"",
					"",
					"0",
					"0",
				},
				{
					"https://github.com/random",
					"Random",
					"Profiles/GitHub",
					"2",
					"0",
				},
			},
		},
		{
			name: "Multiple-Records-Blank-Last-Record-Header-Disabled",
			args: args{
				bookmarks: []Bookmark{
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
					{
						URL:    nil,
						Title:  "",
						Folder: "",
						ID:     0,
						Parent: 0,
					},
				},
				enableHeader: false,
			},
			want: [][]string{
				{
					"https://github.com/vaguecoder",
					"Vague Coder",
					"Profiles/GitHub",
					"1",
					"0",
				},
				{
					"https://github.com/random",
					"Random",
					"Profiles/GitHub",
					"2",
					"0",
				},
				{
					"",
					"",
					"",
					"0",
					"0",
				},
			},
		},
		{
			name: "Multiple-Records-Blank-Last-Record-Header-Enabled",
			args: args{
				bookmarks: []Bookmark{
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
					{
						URL:    nil,
						Title:  "",
						Folder: "",
						ID:     0,
						Parent: 0,
					},
				},
				enableHeader: true,
			},
			want: [][]string{
				{
					"URL",
					"TITLE",
					"FOLDER",
					"ID",
					"PARENT",
				},
				{
					"---",
					"-----",
					"------",
					"--",
					"------",
				},
				{
					"https://github.com/vaguecoder",
					"Vague Coder",
					"Profiles/GitHub",
					"1",
					"0",
				},
				{
					"https://github.com/random",
					"Random",
					"Profiles/GitHub",
					"2",
					"0",
				},
				{
					"",
					"",
					"",
					"0",
					"0",
				},
			},
		},
		{
			name: "Multiple-Records-Trim-Space-Header-Disabled",
			args: args{
				bookmarks: []Bookmark{
					{
						URL:    ptrStr(whitespaceCharacters + "https://github.com/vaguecoder" + whitespaceCharacters),
						Title:  whitespaceCharacters + "Vague Coder" + whitespaceCharacters,
						Folder: whitespaceCharacters + "Profiles/GitHub" + whitespaceCharacters,
						ID:     1,
						Parent: 0,
					},
					{
						URL:    ptrStr(whitespaceCharacters + "https://github.com/random" + whitespaceCharacters),
						Title:  whitespaceCharacters + "Random" + whitespaceCharacters,
						Folder: whitespaceCharacters + "Profiles/GitHub" + whitespaceCharacters,
						ID:     2,
						Parent: 0,
					},
				},
				enableHeader: false,
			},
			want: [][]string{
				{
					"https://github.com/vaguecoder",
					"Vague Coder",
					"Profiles/GitHub",
					"1",
					"0",
				},
				{
					"https://github.com/random",
					"Random",
					"Profiles/GitHub",
					"2",
					"0",
				},
			},
		},
		{
			name: "Multiple-Records-Trim-Space-Header-Enabled",
			args: args{
				bookmarks: []Bookmark{
					{
						URL:    ptrStr(whitespaceCharacters + "https://github.com/vaguecoder" + whitespaceCharacters),
						Title:  whitespaceCharacters + "Vague Coder" + whitespaceCharacters,
						Folder: whitespaceCharacters + "Profiles/GitHub" + whitespaceCharacters,
						ID:     1,
						Parent: 0,
					},
					{
						URL:    ptrStr(whitespaceCharacters + "https://github.com/random" + whitespaceCharacters),
						Title:  whitespaceCharacters + "Random" + whitespaceCharacters,
						Folder: whitespaceCharacters + "Profiles/GitHub" + whitespaceCharacters,
						ID:     2,
						Parent: 0,
					},
				},
				enableHeader: true,
			},
			want: [][]string{
				{
					"URL",
					"TITLE",
					"FOLDER",
					"ID",
					"PARENT",
				},
				{
					"---",
					"-----",
					"------",
					"--",
					"------",
				},
				{
					"https://github.com/vaguecoder",
					"Vague Coder",
					"Profiles/GitHub",
					"1",
					"0",
				},
				{
					"https://github.com/random",
					"Random",
					"Profiles/GitHub",
					"2",
					"0",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BookmarksTable(tt.args.bookmarks, tt.args.enableHeader); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BookmarksTable() = %v , want %v", got, tt.want)
			}
		})
	}
}

func Test_trimSpace(t *testing.T) {
	str := `Luffy!`
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Blank-String",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "Leading-Space-Should-Be-Trimmed",
			args: args{
				s: whitespaceCharacters + str,
			},
			want: str,
		},
		{
			name: "Trailing-Space-Should-Be-Trimmed",
			args: args{
				s: str + whitespaceCharacters,
			},
			want: str,
		},
		{
			name: "Space-In-Middle-Of-String-Should-Not-Trim",
			args: args{
				s: str + whitespaceCharacters + str,
			},
			want: str + whitespaceCharacters + str,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimSpace(tt.args.s); got != tt.want {
				t.Errorf("trimSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerUnderline(t *testing.T) {
	type args struct {
		header []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Empty-Input",
			args: args{
				header: []string{},
			},
			want: []string{},
		},
		{
			name: "Single-Column",
			args: args{
				header: []string{
					"Sno.",
				},
			},
			want: []string{
				"----",
			},
		},
		{
			name: "Multiple-Column",
			args: args{
				header: []string{
					"Sno.", "Name",
				},
			},
			want: []string{
				"----", "----",
			},
		},
		{
			name: "Actual-Input",
			args: args{
				header: []string{
					"URL", "TITLE", "FOLDER", "ID", "PARENT",
				},
			},
			want: []string{
				"---", "-----", "------", "--", "------",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := headerUnderline(tt.args.header); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("headerUnderline() = %v, want %v", got, tt.want)
			}
		})
	}
}
