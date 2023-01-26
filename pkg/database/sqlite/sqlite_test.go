package sqlite

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/vaguecoder/firefox-backups/pkg/files"
	"github.com/vaguecoder/firefox-backups/pkg/util"
)

func TestNewDB(t *testing.T) {
	ctx := context.Background()
	filesOps := files.NewOperator(ctx)

	type args struct {
		dbFilename string
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		permission *int
	}{
		{
			name: "Valid-File",
			args: args{
				dbFilename: "testdata/places.sqlite",
			},
			wantErr:    false,
			permission: nil,
		},
		{
			name: "Invalid-File",
			args: args{
				dbFilename: "testdata/random.sqlite",
			},
			wantErr:    true,
			permission: util.PtrInt(555),
		},
	}
	for _, tt := range tests {
		var (
			permission  int
			revertChmod files.ChmodRevertFunc
			err         error
		)
		if tt.permission != nil {
			permission = *tt.permission
			revertChmod, err = filesOps.Chmod(tt.args.dbFilename, permission)
			if err != nil {
				t.Fatalf("Failed to change the test file %q 's permission to %d: %v",
					tt.args.dbFilename, permission, err)
			}

			defer func() {
				fileMode, err := revertChmod(tt.args.dbFilename)
				if err != nil {
					t.Fatalf("failed to revert file %q 's permissions to %d: %v",
						tt.args.dbFilename, fileMode, err)
				}
			}()
		}
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewDB(tt.args.dbFilename)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
