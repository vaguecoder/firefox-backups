package bookmark

import (
	"fmt"
	"strings"
)

type Bookmark struct {
	URL    *string `json:"url" yaml:"url"`
	Title  string  `json:"title" yaml:"title"`
	Folder string  `json:"folder" yaml:"folder"`
	ID     int     `json:"id" yaml:"id"`
	Parent int     `json:"parent" yaml:"parent"`
}

func BookmarksTable(bookmarks []Bookmark, enableHeader bool) [][]string {
	var (
		sheet [][]string
		url   string

		header = []string{"URL", "TITLE", "FOLDER", "ID", "PARENT"}
	)

	if enableHeader {
		sheet = append(sheet, header)
		sheet = append(sheet, headerUnderline(header))
	}

	for _, b := range bookmarks {
		if b.URL == nil {
			url = ""
		} else {
			url = *b.URL
		}

		sheet = append(sheet, []string{
			trimSpace(url),
			trimSpace(b.Title),
			trimSpace(b.Folder),
			trimSpace(fmt.Sprint(b.ID)),
			trimSpace(fmt.Sprint(b.Parent)),
		})
	}

	return sheet
}

func trimSpace(s string) string {
	return strings.TrimSpace(s)
}

func headerUnderline(header []string) []string {
	underline := []string{}
	for _, title := range header {
		underline = append(underline, strings.Repeat("-", len(title)))
	}

	return underline
}
