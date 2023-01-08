package database

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/vaguecoder/firefox-backups/pkg/logs"
	"github.com/vaguecoder/firefox-backups/pkg/util"
)

type filterMap map[filterName]func(context.Context, *[]bookmark) error

type filterName string

func (f filterName) String() string {
	return string(f)
}

type filterNames []filterName

func (f filterNames) String() string {
	var filters []string

	for _, filter := range f {
		filters = append(filters, filter.String())
	}

	sort.Strings(filters)

	return strings.Join(filters, ", ")
}

var AllFilters filterNames

func (d *DatabaseOperator) applyFilters(ctx context.Context, bookmarks *[]bookmark) error {
	logger := logs.FromContext(ctx)

	filterNames, err := util.ToStringers(util.MapKeys(d.filters))
	if err != nil {
		logger.Error().Err(err).Msg("Failed to parse filter names as stringers")
		return fmt.Errorf("failed to parse filter names as stringers %v: %v", d.filters, err)
	}

	logger.Info().Stringers("filters", filterNames).Msg("Added filters")
	for filterName, filterFunc := range d.filters {
		err := filterFunc(ctx, bookmarks)
		if err != nil {
			logger.Error().Err(err).Stringer("filter", filterName).Msg("Failed to apply filter")
			return fmt.Errorf("failed to apply filter %q: %v", filterName, err)
		}
		logger.Info().Stringer("filter", filterName).Msg("Filtered successfully")
	}

	return nil
}
