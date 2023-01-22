package filters

import (
	"context"
	"fmt"

	"github.com/vaguecoder/firefox-backups/pkg/bookmark"
)

type filterManager struct {
	bookmarks []bookmark.Bookmark
	filters   []Filter
}

func NewFilterManager() *filterManager {
	return &filterManager{
		filters: []Filter{},
	}
}

func (f *filterManager) Filter(filter Filter) *filterManager {
	f.filters = append(f.filters, filter)

	return f
}

func (f *filterManager) Bookmarks(bookmarks []bookmark.Bookmark) *filterManager {
	f.bookmarks = bookmarks

	return f
}

func (f *filterManager) Apply(ctx context.Context) ([]bookmark.Bookmark, error) {
	if f.bookmarks == nil {
		return nil, fmt.Errorf("bookmarks missing in chaining")
	}

	var err error

	for _, filter := range f.filters {
		if filter == nil {
			continue
		}

		f.bookmarks, err = filter.Apply(ctx, f.bookmarks)
		if err != nil {
			return nil, fmt.Errorf("failed to apply filter %q: %v", filter, err)
		}
	}

	return f.bookmarks, nil
}
