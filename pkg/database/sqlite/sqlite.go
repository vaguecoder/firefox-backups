package sqlite

import "database/sql"

type DB struct {
}

type DBConnection interface {
	Query(query string, args ...any) (*sql.Rows, error)
}

const sqliteDriverName = `sqlite3`

func NewDB(dbFilename string) (DBConnection, error) {
	conn, err := sql.Open(sqliteDriverName, dbFilename)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
