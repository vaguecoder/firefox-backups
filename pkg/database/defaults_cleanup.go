package database

import (
	"context"
)

const (
	mozillaFirefoxFolder                 = `menu/Mozilla Firefox`
	defaultsCleanUpFilterName filterName = "defaults-clean-up"
)

func (d *DatabaseOperator) removeDefaults(ctx context.Context, bookmarks *[]bookmark) error {
	var result []bookmark
	var mozillaFirefoxTitleId int

	for _, b := range *bookmarks {
		if b.URL == nil {
			continue
		}

		if b.Folder == mozillaFirefoxFolder {
			mozillaFirefoxTitleId = b.id
			continue
		}

		if b.parent == mozillaFirefoxTitleId {
			continue
		}

		result = append(result, b)
	}

	(*bookmarks) = result

	return nil
}
