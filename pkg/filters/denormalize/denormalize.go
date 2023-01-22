package denormalize

import (
	"context"
	"fmt"
	"sort"

	"github.com/vaguecoder/firefox-backups/pkg/bookmark"
	"github.com/vaguecoder/firefox-backups/pkg/constants"
	"github.com/vaguecoder/firefox-backups/pkg/filters"
	"github.com/vaguecoder/firefox-backups/pkg/logs"
)

var FilterName = filters.ToFilterName(constants.DenormalizeFilter)

func init() {
	filters.AllFilterNames = append(filters.AllFilterNames, FilterName)
}

type Denormalizer struct{}

func (d *Denormalizer) Apply(ctx context.Context, bookmarks []bookmark.Bookmark) ([]bookmark.Bookmark, error) {
	logger := logs.FromContext(ctx).With().Int("initial-count", len(bookmarks)).
		Stringer("filter", FilterName).Logger()

	sort.Slice(bookmarks, func(i, j int) bool {
		return (bookmarks)[i].ID < (bookmarks)[j].ID
	})
	logger.Info().Msg("Bookmarks sorted on ID")

	var prevLen int
	for len(bookmarks) != prevLen {
		prevLen = len(bookmarks)
		for _, b := range bookmarks {
			result := updatePathInChildBookmarks(b, bookmarks)
			if result != nil {
				bookmarks = result
				break
			}
		}
	}

	return bookmarks, nil
}

func (d *Denormalizer) String() string {
	return FilterName.String()
}

func updatePathInChildBookmarks(parentBookmark bookmark.Bookmark, inputBookmarks []bookmark.Bookmark) []bookmark.Bookmark {
	var resultantBookmarks []bookmark.Bookmark
	var isUpdated bool

	for _, b := range inputBookmarks {
		if parentBookmark.ID == b.ID {
			continue
		}

		if b.Parent == parentBookmark.ID {
			isUpdated = true

			current := bookmark.Bookmark{
				ID:     b.ID,
				Parent: b.Parent,
				URL:    b.URL,
				Title:  b.Title,
				Folder: "",
			}

			parentTitle := parentBookmark.Title
			if parentBookmark.Folder != "" {
				parentTitle = fmt.Sprintf("%s/%s", parentBookmark.Folder, parentTitle)
			}

			if b.Folder != "" {
				current.Folder = fmt.Sprintf("%s/%s", parentTitle, b.Folder)
			} else {
				current.Folder = parentTitle
			}

			resultantBookmarks = append(resultantBookmarks, current)
			continue
		}

		resultantBookmarks = append(resultantBookmarks, b)
	}
	if isUpdated {
		return resultantBookmarks
	}

	return nil
}
