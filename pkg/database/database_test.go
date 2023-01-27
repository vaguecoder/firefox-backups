package database

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vaguecoder/firefox-backups/pkg/bookmark"
	"github.com/vaguecoder/firefox-backups/pkg/mocks"
	"github.com/vaguecoder/firefox-backups/pkg/util"

	sqlMock "github.com/DATA-DOG/go-sqlmock"
)

var ptrStr = util.PtrStr

func TestDatabaseOperator_GetBookmarks(t *testing.T) {
	ctx := context.Background()

	sqlDB, mockServer, err := sqlMock.New()
	require.NoError(t, err, "unexpected error at DB mock server creation")

	type fields struct {
		db *mocks.DBConnection
	}

	tests := []struct {
		name       string
		fields     fields
		want       []bookmark.Bookmark
		rows       [][]string
		wantErr    bool
		dbQueryErr bool
	}{
		{
			name: "Valid-Case-With-2-Records",
			fields: fields{
				db: new(mocks.DBConnection),
			},
			want: []bookmark.Bookmark{
				{
					URL:    ptrStr("https://github.com/vaguecoder"),
					Title:  "Vague Coder",
					Folder: "",
					ID:     1,
					Parent: 0,
				},
				{
					URL:    ptrStr("https://github.com/random"),
					Title:  "Random",
					Folder: "",
					ID:     2,
					Parent: 0,
				},
			},
			rows: [][]string{
				{
					"id",
					"parent",
					"url",
					"title",
				},
				{
					"1",
					"0",
					"https://github.com/vaguecoder",
					"Vague Coder",
				},
				{
					"2",
					"0",
					"https://github.com/random",
					"Random",
				},
			},
			wantErr:    false,
			dbQueryErr: false,
		},
		{
			name: "Failure-At-DB-Query",
			fields: fields{
				db: new(mocks.DBConnection),
			},
			want: []bookmark.Bookmark(nil),
			rows: [][]string{
				{
					"id",
					"parent",
					"url",
					"title",
				},
				{
					"1",
					"0",
					"https://github.com/vaguecoder",
					"Vague Coder",
				},
				{
					"2",
					"0",
					"https://github.com/random",
					"Random",
				},
			},
			wantErr:    true,
			dbQueryErr: true,
		},
		{
			name: "Failure-At-Scan-Due-To-Invalid-Data-Type",
			fields: fields{
				db: new(mocks.DBConnection),
			},
			want: []bookmark.Bookmark(nil),
			rows: [][]string{
				{
					"id",
					"parent",
					"url",
					"title",
				},
				{
					"1APPLE",
					"0",
					"https://github.com/vaguecoder",
					"Vague Coder",
				},
				{
					"2",
					"0",
					"https://github.com/random",
					"Random",
				},
			},
			wantErr:    true,
			dbQueryErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDatabaseOperator(tt.fields.db)

			if tt.dbQueryErr {
				tt.fields.db.On("Query", queryStr).Return(nil, fmt.Errorf("some error"))
			} else {
				rows, err := sqlRows(sqlDB, mockServer, tt.rows)
				require.NoError(t, err, "Unexpected error at mock rows creation")

				tt.fields.db.On("Query", queryStr).Return(rows, nil)
			}

			got, err := d.GetBookmarks(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("DatabaseOperator.GetBookmarks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DatabaseOperator.GetBookmarks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func sqlRows(db *sql.DB, mockServer sqlMock.Sqlmock, data [][]string) (*sql.Rows, error) {
	if len(data) == 0 {
		return &sql.Rows{}, nil
	}

	mockRows := sqlMock.NewRows(data[0])

	for _, record := range data[1:] {
		var values []driver.Value

		for _, val := range record {
			values = append(values, val)
		}

		mockRows = mockRows.AddRow(values...)
	}

	mockServer.ExpectQuery(queryStr).WillReturnRows(mockRows)

	resultRows, err := db.Query(queryStr)
	if err != nil {
		return nil, fmt.Errorf("Failed to query the mock rows: %v", err)
	}

	return resultRows, nil
}
