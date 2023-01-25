package text

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_maxLengthOfColumns(t *testing.T) {
	type args struct {
		data [][]string
	}
	tests := []struct {
		name string
		args args
		want map[uint]uint
	}{
		{
			name: "",
			args: args{
				data: [][]string{
					{
						"url",
						"title",
						"folder",
						"id",
						"parent",
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
			want: map[uint]uint{
				0: uint(len("https://github.com/vaguecoder")),
				1: uint(len("Vague Coder")),
				2: uint(len("Profiles/GitHub")),
				3: uint(len("id")),
				4: uint(len("parent")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateMaxLengthOfColumns(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxLengthOfColumns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTable(t *testing.T) {
	lineStart := `+++`

	type args struct {
		data            [][]string
		headerSeperator bool
		lineStart       string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "No-Data-Without-Header-Seperator-Without-LineStart",
			args: args{
				data:            [][]string{},
				headerSeperator: false,
				lineStart:       "",
			},
			want: []string{},
		},
		{
			name: "No-Data-With-Header-Seperator-Without-LineStart",
			args: args{
				data:            [][]string{},
				headerSeperator: true,
				lineStart:       "",
			},
			want: []string{},
		},
		{
			name: "No-Data-Without-Header-Seperator-With-LineStart",
			args: args{
				data:            [][]string{},
				headerSeperator: false,
				lineStart:       lineStart,
			},
			want: []string{},
		},
		{
			name: "No-Data-With-Header-Seperator-With-LineStart",
			args: args{
				data:            [][]string{},
				headerSeperator: true,
				lineStart:       lineStart,
			},
			want: []string{},
		},
		{
			name: "Only-Header-Without-Header-Seperator-Without-LineStart",
			args: args{
				data: [][]string{
					{
						"url",
						"title",
						"folder",
						"id",
						"parent",
					},
				},
				headerSeperator: false,
				lineStart:       "",
			},
			want: []string{
				("--------------------------------------"),
				("| url | title | folder | id | parent |"),
				("--------------------------------------"),
			},
		},
		{
			name: "Only-Header-With-Header-Seperator-Without-LineStart",
			args: args{
				data: [][]string{
					{
						"url",
						"title",
						"folder",
						"id",
						"parent",
					},
				},
				headerSeperator: true,
				lineStart:       "",
			},
			want: []string{
				"--------------------------------------",
				"| URL | TITLE | FOLDER | ID | PARENT |",
				"--------------------------------------",
			},
		},
		{
			name: "Only-Header-Without-Header-Seperator-With-LineStart",
			args: args{
				data: [][]string{
					{
						"url",
						"title",
						"folder",
						"id",
						"parent",
					},
				},
				headerSeperator: false,
				lineStart:       lineStart,
			},
			want: []string{
				fmt.Sprintf("%s--------------------------------------", lineStart),
				fmt.Sprintf("%s| url | title | folder | id | parent |", lineStart),
				fmt.Sprintf("%s--------------------------------------", lineStart),
			},
		},
		{
			name: "Only-Header-With-Header-Seperator-With-LineStart",
			args: args{
				data: [][]string{
					{
						"url",
						"title",
						"folder",
						"id",
						"parent",
					},
				},
				headerSeperator: true,
				lineStart:       lineStart,
			},
			want: []string{
				fmt.Sprintf("%s--------------------------------------", lineStart),
				fmt.Sprintf("%s| URL | TITLE | FOLDER | ID | PARENT |", lineStart),
				fmt.Sprintf("%s--------------------------------------", lineStart),
			},
		},
		{
			name: "Header-With-Single-Row-Without-Header-Seperator-Without-LineStart",
			args: args{
				data: [][]string{
					{
						"url",
						"title",
						"folder",
						"id",
						"parent",
					},
					{
						"https://github.com/vaguecoder",
						"Vague Coder",
						"Profiles/GitHub",
						"1",
						"0",
					},
				},
				headerSeperator: false,
				lineStart:       "",
			},
			want: []string{
				"-------------------------------------------------------------------------------",
				"| url                           | title       | folder          | id | parent |",
				"| https://github.com/vaguecoder | Vague Coder | Profiles/GitHub | 1  | 0      |",
				"-------------------------------------------------------------------------------",
			},
		},
		{
			name: "Header-With-Single-Row-With-Header-Seperator-Without-LineStart",
			args: args{
				data: [][]string{
					{
						"url",
						"title",
						"folder",
						"id",
						"parent",
					},
					{
						"https://github.com/vaguecoder",
						"Vague Coder",
						"Profiles/GitHub",
						"1",
						"0",
					},
				},
				headerSeperator: true,
				lineStart:       "",
			},
			want: []string{
				"-------------------------------------------------------------------------------",
				"| URL                           | TITLE       | FOLDER          | ID | PARENT |",
				"-------------------------------------------------------------------------------",
				"| https://github.com/vaguecoder | Vague Coder | Profiles/GitHub | 1  | 0      |",
				"-------------------------------------------------------------------------------",
			},
		},
		{
			name: "Header-With-Single-Row-Without-Header-Seperator-With-LineStart",
			args: args{
				data: [][]string{
					{
						"url",
						"title",
						"folder",
						"id",
						"parent",
					},
					{
						"https://github.com/vaguecoder",
						"Vague Coder",
						"Profiles/GitHub",
						"1",
						"0",
					},
				},
				headerSeperator: false,
				lineStart:       lineStart,
			},
			want: []string{
				fmt.Sprintf("%s-------------------------------------------------------------------------------", lineStart),
				fmt.Sprintf("%s| url                           | title       | folder          | id | parent |", lineStart),
				fmt.Sprintf("%s| https://github.com/vaguecoder | Vague Coder | Profiles/GitHub | 1  | 0      |", lineStart),
				fmt.Sprintf("%s-------------------------------------------------------------------------------", lineStart),
			},
		},
		{
			name: "Header-With-Single-Row-With-Header-Seperator-With-LineStart",
			args: args{
				data: [][]string{
					{
						"url",
						"title",
						"folder",
						"id",
						"parent",
					},
					{
						"https://github.com/vaguecoder",
						"Vague Coder",
						"Profiles/GitHub",
						"1",
						"0",
					},
				},
				headerSeperator: true,
				lineStart:       lineStart,
			},
			want: []string{
				fmt.Sprintf("%s-------------------------------------------------------------------------------", lineStart),
				fmt.Sprintf("%s| URL                           | TITLE       | FOLDER          | ID | PARENT |", lineStart),
				fmt.Sprintf("%s-------------------------------------------------------------------------------", lineStart),
				fmt.Sprintf("%s| https://github.com/vaguecoder | Vague Coder | Profiles/GitHub | 1  | 0      |", lineStart),
				fmt.Sprintf("%s-------------------------------------------------------------------------------", lineStart),
			},
		},
		{
			name: "Header-With-Empty-Row-Without-Header-Seperator-Without-LineStart",
			args: args{
				data: [][]string{
					{
						"url",
						"title",
						"folder",
						"id",
						"parent",
					},
					{
						"",
						"",
						"",
						"",
						"",
					},
				},
				headerSeperator: false,
				lineStart:       "",
			},
			want: []string{
				"--------------------------------------",
				"| url | title | folder | id | parent |",
				"|     |       |        |    |        |",
				"--------------------------------------",
			},
		},
		{
			name: "Header-With-Empty-Row-With-Header-Seperator-Without-LineStart",
			args: args{
				data: [][]string{
					{
						"url",
						"title",
						"folder",
						"id",
						"parent",
					},
					{
						"",
						"",
						"",
						"",
						"",
					},
				},
				headerSeperator: true,
				lineStart:       "",
			},
			want: []string{
				"--------------------------------------",
				"| URL | TITLE | FOLDER | ID | PARENT |",
				"--------------------------------------",
				"|     |       |        |    |        |",
				"--------------------------------------",
			},
		},
		{
			name: "Header-With-Empty-Row-Without-Header-Seperator-With-LineStart",
			args: args{
				data: [][]string{
					{
						"url",
						"title",
						"folder",
						"id",
						"parent",
					},
					{
						"",
						"",
						"",
						"",
						"",
					},
				},
				headerSeperator: false,
				lineStart:       lineStart,
			},
			want: []string{
				fmt.Sprintf("%s--------------------------------------", lineStart),
				fmt.Sprintf("%s| url | title | folder | id | parent |", lineStart),
				fmt.Sprintf("%s|     |       |        |    |        |", lineStart),
				fmt.Sprintf("%s--------------------------------------", lineStart),
			},
		},
		{
			name: "Header-With-Empty-Row-With-Header-Seperator-With-LineStart",
			args: args{
				data: [][]string{
					{
						"url",
						"title",
						"folder",
						"id",
						"parent",
					},
					{
						"",
						"",
						"",
						"",
						"",
					},
				},
				headerSeperator: true,
				lineStart:       lineStart,
			},
			want: []string{
				fmt.Sprintf("%s--------------------------------------", lineStart),
				fmt.Sprintf("%s| URL | TITLE | FOLDER | ID | PARENT |", lineStart),
				fmt.Sprintf("%s--------------------------------------", lineStart),
				fmt.Sprintf("%s|     |       |        |    |        |", lineStart),
				fmt.Sprintf("%s--------------------------------------", lineStart),
			},
		},
		{
			name: "Header-With-Multiple-Rows-Without-Header-Seperator-Without-LineStart",
			args: args{
				data: [][]string{
					{
						"url",
						"title",
						"folder",
						"id",
						"parent",
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
				headerSeperator: false,
				lineStart:       "",
			},
			want: []string{
				"-------------------------------------------------------------------------------",
				"| url                           | title       | folder          | id | parent |",
				"| https://github.com/vaguecoder | Vague Coder | Profiles/GitHub | 1  | 0      |",
				"| https://github.com/random     | Random      | Profiles/GitHub | 2  | 0      |",
				"-------------------------------------------------------------------------------",
			},
		},
		{
			name: "Header-With-Multiple-Rows-With-Header-Seperator-Without-LineStart",
			args: args{
				data: [][]string{
					{
						"url",
						"title",
						"folder",
						"id",
						"parent",
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
				headerSeperator: true,
				lineStart:       "",
			},
			want: []string{
				"-------------------------------------------------------------------------------",
				"| URL                           | TITLE       | FOLDER          | ID | PARENT |",
				"-------------------------------------------------------------------------------",
				"| https://github.com/vaguecoder | Vague Coder | Profiles/GitHub | 1  | 0      |",
				"| https://github.com/random     | Random      | Profiles/GitHub | 2  | 0      |",
				"-------------------------------------------------------------------------------",
			},
		},
		{
			name: "Header-With-Multiple-Rows-Without-Header-Seperator-With-LineStart",
			args: args{
				data: [][]string{
					{
						"url",
						"title",
						"folder",
						"id",
						"parent",
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
				headerSeperator: false,
				lineStart:       lineStart,
			},
			want: []string{
				fmt.Sprintf("%s-------------------------------------------------------------------------------", lineStart),
				fmt.Sprintf("%s| url                           | title       | folder          | id | parent |", lineStart),
				fmt.Sprintf("%s| https://github.com/vaguecoder | Vague Coder | Profiles/GitHub | 1  | 0      |", lineStart),
				fmt.Sprintf("%s| https://github.com/random     | Random      | Profiles/GitHub | 2  | 0      |", lineStart),
				fmt.Sprintf("%s-------------------------------------------------------------------------------", lineStart),
			},
		},
		{
			name: "Header-With-Multiple-Rows-With-Header-Seperator-With-LineStart",
			args: args{
				data: [][]string{
					{
						"url",
						"title",
						"folder",
						"id",
						"parent",
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
				headerSeperator: true,
				lineStart:       lineStart,
			},
			want: []string{
				fmt.Sprintf("%s-------------------------------------------------------------------------------", lineStart),
				fmt.Sprintf("%s| URL                           | TITLE       | FOLDER          | ID | PARENT |", lineStart),
				fmt.Sprintf("%s-------------------------------------------------------------------------------", lineStart),
				fmt.Sprintf("%s| https://github.com/vaguecoder | Vague Coder | Profiles/GitHub | 1  | 0      |", lineStart),
				fmt.Sprintf("%s| https://github.com/random     | Random      | Profiles/GitHub | 2  | 0      |", lineStart),
				fmt.Sprintf("%s-------------------------------------------------------------------------------", lineStart),
			},
		},
		{
			name: "Header-With-Multiple-Rows-First-Row-Empty-Without-Header-Seperator-Without-LineStart",
			args: args{
				data: [][]string{
					{
						"url",
						"title",
						"folder",
						"id",
						"parent",
					},
					{
						"",
						"",
						"",
						"",
						"",
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
				headerSeperator: false,
				lineStart:       "",
			},
			want: []string{
				"-------------------------------------------------------------------------------",
				"| url                           | title       | folder          | id | parent |",
				"|                               |             |                 |    |        |",
				"| https://github.com/vaguecoder | Vague Coder | Profiles/GitHub | 1  | 0      |",
				"| https://github.com/random     | Random      | Profiles/GitHub | 2  | 0      |",
				"-------------------------------------------------------------------------------",
			},
		},
		{
			name: "Header-With-Multiple-Rows-First-Row-Empty-With-Header-Seperator-Without-LineStart",
			args: args{
				data: [][]string{
					{
						"url",
						"title",
						"folder",
						"id",
						"parent",
					},
					{
						"",
						"",
						"",
						"",
						"",
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
				headerSeperator: true,
				lineStart:       "",
			},
			want: []string{
				"-------------------------------------------------------------------------------",
				"| URL                           | TITLE       | FOLDER          | ID | PARENT |",
				"-------------------------------------------------------------------------------",
				"|                               |             |                 |    |        |",
				"| https://github.com/vaguecoder | Vague Coder | Profiles/GitHub | 1  | 0      |",
				"| https://github.com/random     | Random      | Profiles/GitHub | 2  | 0      |",
				"-------------------------------------------------------------------------------",
			},
		},
		{
			name: "Header-With-Multiple-Rows-First-Row-Empty-Without-Header-Seperator-With-LineStart",
			args: args{
				data: [][]string{
					{
						"url",
						"title",
						"folder",
						"id",
						"parent",
					},
					{
						"",
						"",
						"",
						"",
						"",
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
				headerSeperator: false,
				lineStart:       lineStart,
			},
			want: []string{
				fmt.Sprintf("%s-------------------------------------------------------------------------------", lineStart),
				fmt.Sprintf("%s| url                           | title       | folder          | id | parent |", lineStart),
				fmt.Sprintf("%s|                               |             |                 |    |        |", lineStart),
				fmt.Sprintf("%s| https://github.com/vaguecoder | Vague Coder | Profiles/GitHub | 1  | 0      |", lineStart),
				fmt.Sprintf("%s| https://github.com/random     | Random      | Profiles/GitHub | 2  | 0      |", lineStart),
				fmt.Sprintf("%s-------------------------------------------------------------------------------", lineStart),
			},
		},
		{
			name: "Header-With-Multiple-Rows-First-Row-Empty-With-Header-Seperator-With-LineStart",
			args: args{
				data: [][]string{
					{
						"url",
						"title",
						"folder",
						"id",
						"parent",
					},
					{
						"",
						"",
						"",
						"",
						"",
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
				headerSeperator: true,
				lineStart:       lineStart,
			},
			want: []string{
				fmt.Sprintf("%s-------------------------------------------------------------------------------", lineStart),
				fmt.Sprintf("%s| URL                           | TITLE       | FOLDER          | ID | PARENT |", lineStart),
				fmt.Sprintf("%s-------------------------------------------------------------------------------", lineStart),
				fmt.Sprintf("%s|                               |             |                 |    |        |", lineStart),
				fmt.Sprintf("%s| https://github.com/vaguecoder | Vague Coder | Profiles/GitHub | 1  | 0      |", lineStart),
				fmt.Sprintf("%s| https://github.com/random     | Random      | Profiles/GitHub | 2  | 0      |", lineStart),
				fmt.Sprintf("%s-------------------------------------------------------------------------------", lineStart),
			},
		},
		{
			name: "Header-With-Multiple-Rows-Middle-Row-Empty-Without-Header-Seperator-Without-LineStart",
			args: args{
				data: [][]string{
					{
						"url",
						"title",
						"folder",
						"id",
						"parent",
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
						"",
						"",
					},
					{
						"https://github.com/random",
						"Random",
						"Profiles/GitHub",
						"2",
						"0",
					},
				},
				headerSeperator: false,
				lineStart:       "",
			},
			want: []string{
				"-------------------------------------------------------------------------------",
				"| url                           | title       | folder          | id | parent |",
				"| https://github.com/vaguecoder | Vague Coder | Profiles/GitHub | 1  | 0      |",
				"|                               |             |                 |    |        |",
				"| https://github.com/random     | Random      | Profiles/GitHub | 2  | 0      |",
				"-------------------------------------------------------------------------------",
			},
		},
		{
			name: "Header-With-Multiple-Rows-Middle-Row-Empty-With-Header-Seperator-Without-LineStart",
			args: args{
				data: [][]string{
					{
						"url",
						"title",
						"folder",
						"id",
						"parent",
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
						"",
						"",
					},
					{
						"https://github.com/random",
						"Random",
						"Profiles/GitHub",
						"2",
						"0",
					},
				},
				headerSeperator: true,
				lineStart:       "",
			},
			want: []string{
				"-------------------------------------------------------------------------------",
				"| URL                           | TITLE       | FOLDER          | ID | PARENT |",
				"-------------------------------------------------------------------------------",
				"| https://github.com/vaguecoder | Vague Coder | Profiles/GitHub | 1  | 0      |",
				"|                               |             |                 |    |        |",
				"| https://github.com/random     | Random      | Profiles/GitHub | 2  | 0      |",
				"-------------------------------------------------------------------------------",
			},
		},
		{
			name: "Header-With-Multiple-Rows-Middle-Row-Empty-Without-Header-Seperator-With-LineStart",
			args: args{
				data: [][]string{
					{
						"url",
						"title",
						"folder",
						"id",
						"parent",
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
						"",
						"",
					},
					{
						"https://github.com/random",
						"Random",
						"Profiles/GitHub",
						"2",
						"0",
					},
				},
				headerSeperator: false,
				lineStart:       lineStart,
			},
			want: []string{
				fmt.Sprintf("%s-------------------------------------------------------------------------------", lineStart),
				fmt.Sprintf("%s| url                           | title       | folder          | id | parent |", lineStart),
				fmt.Sprintf("%s| https://github.com/vaguecoder | Vague Coder | Profiles/GitHub | 1  | 0      |", lineStart),
				fmt.Sprintf("%s|                               |             |                 |    |        |", lineStart),
				fmt.Sprintf("%s| https://github.com/random     | Random      | Profiles/GitHub | 2  | 0      |", lineStart),
				fmt.Sprintf("%s-------------------------------------------------------------------------------", lineStart),
			},
		},
		{
			name: "Header-With-Multiple-Rows-Middle-Row-Empty-With-Header-Seperator-With-LineStart",
			args: args{
				data: [][]string{
					{
						"url",
						"title",
						"folder",
						"id",
						"parent",
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
						"",
						"",
					},
					{
						"https://github.com/random",
						"Random",
						"Profiles/GitHub",
						"2",
						"0",
					},
				},
				headerSeperator: true,
				lineStart:       lineStart,
			},
			want: []string{
				fmt.Sprintf("%s-------------------------------------------------------------------------------", lineStart),
				fmt.Sprintf("%s| URL                           | TITLE       | FOLDER          | ID | PARENT |", lineStart),
				fmt.Sprintf("%s-------------------------------------------------------------------------------", lineStart),
				fmt.Sprintf("%s| https://github.com/vaguecoder | Vague Coder | Profiles/GitHub | 1  | 0      |", lineStart),
				fmt.Sprintf("%s|                               |             |                 |    |        |", lineStart),
				fmt.Sprintf("%s| https://github.com/random     | Random      | Profiles/GitHub | 2  | 0      |", lineStart),
				fmt.Sprintf("%s-------------------------------------------------------------------------------", lineStart),
			},
		},
		{
			name: "Header-With-Multiple-Rows-Last-Row-Empty-Without-Header-Seperator-Without-LineStart",
			args: args{
				data: [][]string{
					{
						"url",
						"title",
						"folder",
						"id",
						"parent",
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
						"",
						"",
					},
				},
				headerSeperator: false,
				lineStart:       "",
			},
			want: []string{
				"-------------------------------------------------------------------------------",
				"| url                           | title       | folder          | id | parent |",
				"| https://github.com/vaguecoder | Vague Coder | Profiles/GitHub | 1  | 0      |",
				"| https://github.com/random     | Random      | Profiles/GitHub | 2  | 0      |",
				"|                               |             |                 |    |        |",
				"-------------------------------------------------------------------------------",
			},
		},
		{
			name: "Header-With-Multiple-Rows-Last-Row-Empty-With-Header-Seperator-Without-LineStart",
			args: args{
				data: [][]string{
					{
						"url",
						"title",
						"folder",
						"id",
						"parent",
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
						"",
						"",
					},
				},
				headerSeperator: true,
				lineStart:       "",
			},
			want: []string{
				"-------------------------------------------------------------------------------",
				"| URL                           | TITLE       | FOLDER          | ID | PARENT |",
				"-------------------------------------------------------------------------------",
				"| https://github.com/vaguecoder | Vague Coder | Profiles/GitHub | 1  | 0      |",
				"| https://github.com/random     | Random      | Profiles/GitHub | 2  | 0      |",
				"|                               |             |                 |    |        |",
				"-------------------------------------------------------------------------------",
			},
		},
		{
			name: "Header-With-Multiple-Rows-Last-Row-Empty-Without-Header-Seperator-With-LineStart",
			args: args{
				data: [][]string{
					{
						"url",
						"title",
						"folder",
						"id",
						"parent",
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
						"",
						"",
					},
				},
				headerSeperator: false,
				lineStart:       lineStart,
			},
			want: []string{
				fmt.Sprintf("%s-------------------------------------------------------------------------------", lineStart),
				fmt.Sprintf("%s| url                           | title       | folder          | id | parent |", lineStart),
				fmt.Sprintf("%s| https://github.com/vaguecoder | Vague Coder | Profiles/GitHub | 1  | 0      |", lineStart),
				fmt.Sprintf("%s| https://github.com/random     | Random      | Profiles/GitHub | 2  | 0      |", lineStart),
				fmt.Sprintf("%s|                               |             |                 |    |        |", lineStart),
				fmt.Sprintf("%s-------------------------------------------------------------------------------", lineStart),
			},
		},
		{
			name: "Header-With-Multiple-Rows-Last-Row-Empty-With-Header-Seperator-With-LineStart",
			args: args{
				data: [][]string{
					{
						"url",
						"title",
						"folder",
						"id",
						"parent",
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
						"",
						"",
					},
				},
				headerSeperator: true,
				lineStart:       lineStart,
			},
			want: []string{
				fmt.Sprintf("%s-------------------------------------------------------------------------------", lineStart),
				fmt.Sprintf("%s| URL                           | TITLE       | FOLDER          | ID | PARENT |", lineStart),
				fmt.Sprintf("%s-------------------------------------------------------------------------------", lineStart),
				fmt.Sprintf("%s| https://github.com/vaguecoder | Vague Coder | Profiles/GitHub | 1  | 0      |", lineStart),
				fmt.Sprintf("%s| https://github.com/random     | Random      | Profiles/GitHub | 2  | 0      |", lineStart),
				fmt.Sprintf("%s|                               |             |                 |    |        |", lineStart),
				fmt.Sprintf("%s-------------------------------------------------------------------------------", lineStart),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Table(tt.args.data, tt.args.headerSeperator, tt.args.lineStart); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Table() = %v, want %v", got, tt.want)
			}
		})
	}
}
