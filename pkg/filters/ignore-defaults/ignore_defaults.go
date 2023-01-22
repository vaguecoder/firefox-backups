package ignoredefaults

import (
	"context"

	"github.com/vaguecoder/firefox-backups/pkg/bookmark"
	"github.com/vaguecoder/firefox-backups/pkg/constants"
	"github.com/vaguecoder/firefox-backups/pkg/filters"
)

const (
	mozillaFirefoxFolder = `menu/Mozilla Firefox`
)

var FilterName = filters.ToFilterName(constants.IgnoreDefaultsFilter)

func init() {
	filters.AllFilterNames = append(filters.AllFilterNames, FilterName)
}

type DefaultsRemover struct{}

func (d *DefaultsRemover) Apply(ctx context.Context, bookmarks []bookmark.Bookmark) ([]bookmark.Bookmark, error) {
	var result []bookmark.Bookmark
	var mozillaFirefoxTitleId int

	for _, bm := range bookmarks {
		if bm.URL == nil {
			continue
		}

		if bm.Folder == mozillaFirefoxFolder {
			mozillaFirefoxTitleId = bm.ID
			continue
		}

		if bm.Parent == mozillaFirefoxTitleId {
			continue
		}

		result = append(result, bm)
	}

	return result, nil
}

func (d *DefaultsRemover) String() string {
	return FilterName.String()
}
