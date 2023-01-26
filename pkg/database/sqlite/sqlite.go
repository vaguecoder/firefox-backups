package sqlite

import (
	"database/sql"
	"fmt"
)

type DBConnection interface {
	Query(query string, args ...any) (*sql.Rows, error)
}

const sqliteDriverName = `sqlite3`

func NewDB(dbFilename string) (DBConnection, error) {
	// Ignore connection err as ping returns the actual error here
	conn, _ := sql.Open(sqliteDriverName, dbFilename)
	if err := conn.Ping(); err != nil {
		// When ping to DB failed
		return nil, fmt.Errorf("failed to ping SQLite DB: %v", err)
	}

	// When connection to DB was successful
	return conn, nil
}
