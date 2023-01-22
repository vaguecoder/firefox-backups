package database

import (
	"context"
	"fmt"

	"github.com/vaguecoder/firefox-backups/pkg/bookmark"
	"github.com/vaguecoder/firefox-backups/pkg/database/sqlite"
	"github.com/vaguecoder/firefox-backups/pkg/logs"
	"github.com/vaguecoder/firefox-backups/pkg/util"
)

type DatabaseOperator struct {
	db sqlite.DBConnection
}

type BookmarkOperator interface {
	GetBookmarks(context.Context) ([]bookmark.Bookmark, error)
}

const (
	queryStr = `SELECT bookmarks.id, bookmarks.parent, places.URL, bookmarks.title
				FROM moz_places as places
				RIGHT JOIN moz_bookmarks as bookmarks 
				ON places.id = bookmarks.fk`
)

func NewDatabaseOperator(conn sqlite.DBConnection, rawoutput, ignoreDefaults bool) BookmarkOperator {
	return &DatabaseOperator{
		db: conn,
	}
}

func (d *DatabaseOperator) GetBookmarks(ctx context.Context) ([]bookmark.Bookmark, error) {
	logger := logs.FromContext(ctx).With().
		Str("query", util.StrWhitespacesCleanup(queryStr)).Logger()

	query, err := d.db.Query(queryStr)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to query DB")
		return nil, fmt.Errorf("failed to query db: %v", err)
	}

	var bookmarks []bookmark.Bookmark
	for query.Next() {
		var bm bookmark.Bookmark

		err = query.Scan(&bm.ID, &bm.Parent, &bm.URL, &bm.Title)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to execute query")
			return nil, fmt.Errorf("failed to execute query: %v", err)
		}

		bookmarks = append(bookmarks, bm)
	}

	logger.Info().Msg("Successfully executed query and scanned fields")

	return bookmarks, nil
}
