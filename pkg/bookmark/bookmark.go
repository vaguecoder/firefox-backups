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

// BookmarksTable parses the bookmarks data to 2D string table
func BookmarksTable(bookmarks []Bookmark, enableHeader bool) [][]string {
	if len(bookmarks) == 0 {
		// When no bookmarks
		return nil
	}

	var (
		sheet [][]string
		url   string

		// Header for table based on field order, uppercased
		header = []string{"URL", "TITLE", "FOLDER", "ID", "PARENT"}
	)

	if enableHeader {
		// When header toggle is enabled
		sheet = append(sheet, header)
		sheet = append(sheet, headerUnderline(header))
	}

	for _, b := range bookmarks {
		if b.URL == nil {
			// When *url is nil, assign empty string
			url = ""
		} else {
			// Fetch URL string from reference
			url = *b.URL
		}

		// Append record to result
		sheet = append(sheet, []string{
			trimSpace(url),       // Trim leading and trailing whitespace from URL
			trimSpace(b.Title),   // Trim leading and trailing whitespace from title
			trimSpace(b.Folder),  // Trim leading and trailing whitespace from folder name
			fmt.Sprint(b.ID),     // No whitespace trimming required for ID as int is parsed as string
			fmt.Sprint(b.Parent), // No whitespace trimming required for parent ID as int is parsed as string
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
