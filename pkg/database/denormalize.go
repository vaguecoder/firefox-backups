package database

import (
	"context"
	"fmt"
	"sort"

	"github.com/vaguecoder/firefox-backups/pkg/logs"
)

const denormalizeFilterName filterName = "denormalize"

func (d *DatabaseOperator) denormalize(ctx context.Context, bookmarks *[]bookmark) error {
	logger := logs.FromContext(ctx).With().Int("initial-count", len(*bookmarks)).
		Stringer("filter", denormalizeFilterName).Logger()

	sort.Slice(*bookmarks, func(i, j int) bool {
		return (*bookmarks)[i].id < (*bookmarks)[j].id
	})
	logger.Info().Msg("Bookmarks sorted on ID")

	var prevLen int
	for len(*bookmarks) != prevLen {
		prevLen = len(*bookmarks)
		for _, b := range *bookmarks {
			result := d.updatePathInChildBookmarks(b, *bookmarks)
			if result != nil {
				(*bookmarks) = result
				break
			}
		}
	}

	return nil
}

func (d *DatabaseOperator) updatePathInChildBookmarks(parentBookmark bookmark, inputBookmarks []bookmark) []bookmark {
	var resultantBookmarks []bookmark
	var isUpdated bool

	for _, b := range inputBookmarks {
		if parentBookmark.id == b.id {
			continue
		}

		if b.parent == parentBookmark.id {
			isUpdated = true

			current := bookmark{
				id:     b.id,
				parent: b.parent,
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
