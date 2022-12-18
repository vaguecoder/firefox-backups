package database

import (
	"context"

	"github.com/vaguecoder/firefox-backups/pkg/logs"
	"github.com/vaguecoder/firefox-backups/pkg/util"
)

type filterMap map[filterName]func(context.Context, []bookmark) []bookmark

type filterName string

func (f filterName) String() string {
	return string(f)
}

func (d *DatabaseOperator) applyFilters(ctx context.Context, bookmarks []bookmark) []bookmark {
	logger := logs.FromContext(ctx)

	filterNames := util.ToStringers(util.MapKeys(d.filters))

	logger.Info().Stringers("filters", filterNames).Msg("Added filters")
	for filterName, filterFunc := range d.filters {
		bookmarks = filterFunc(ctx, bookmarks)
		logger.Info().Stringer("filter", filterName).Msg("Filtered successfully")
	}

	return bookmarks
}
