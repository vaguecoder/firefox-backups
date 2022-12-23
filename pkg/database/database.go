package database

import (
	"context"
	"fmt"

	"github.com/vaguecoder/firefox-backups/pkg/database/sqlite"
	"github.com/vaguecoder/firefox-backups/pkg/logs"
	"github.com/vaguecoder/firefox-backups/pkg/util"
)

type DatabaseOperator struct {
	db      sqlite.DBConnection
	filters filterMap
}

type bookmark struct {
	URL    *string `json:"url"`
	Title  string  `json:"title"`
	Folder string  `json:"folder"`

	id     int
	parent int
}

type BookmarkOperator interface {
	GetBookmarks(context.Context) ([]bookmark, error)
}

const (
	queryStr = `SELECT bookmarks.id, bookmarks.parent, places.URL, bookmarks.title
				FROM moz_places as places
				RIGHT JOIN moz_bookmarks as bookmarks 
				ON places.id = bookmarks.fk`
)

func NewDatabaseOperator(conn sqlite.DBConnection, rawoutput, ignoreDefaults bool) BookmarkOperator {
	filters := make(filterMap)

	d := DatabaseOperator{
		db: conn,
	}

	if rawoutput {
		return &d
	}
	filters[denormalizeFilterName] = d.denormalize

	if ignoreDefaults {
		filters[defaultsCleanUpFilterName] = d.removeDefaults
	}

	d.filters = filters

	return &d
}

func (d *DatabaseOperator) GetBookmarks(ctx context.Context) ([]bookmark, error) {
	logger := logs.FromContext(ctx).With().
		Str("query", util.StrWhitespacesCleanup(queryStr)).Logger()

	query, err := d.db.Query(queryStr)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to query DB")
		return nil, fmt.Errorf("failed to query db: %v", err)
	}

	var bookmarks []bookmark
	for query.Next() {
		var b bookmark

		err = query.Scan(&b.id, &b.parent, &b.URL, &b.Title)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to execute query")
			return nil, fmt.Errorf("failed to execute query: %v", err)
		}

		bookmarks = append(bookmarks, b)
	}

	logger.Info().Msg("Successfully executed query and scanned fields")

	if err = d.applyFilters(ctx, &bookmarks); err != nil {
		logger.Error().Err(err).Msg("Failed to apply filter(s)")
		return nil, fmt.Errorf("failed to apply filter(s): %v", err)
	}

	return bookmarks, nil
}
