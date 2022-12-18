package database

// func TestDatabaseOperator_FindLeafBookmarks(t *testing.T) {
// 	tests := []struct {
// 		name      string
// 		bookmarks []Bookmark
// 		expected  []Bookmark
// 	}{
// 		{
// 			name: "",
// 			bookmarks: []Bookmark{
// 				{
// 					id:     1,
// 					parent: 0,
// 					URL:    ptrStr("url1"),
// 					Title:  "title1",
// 					Folder: "title1",
// 				},
// 				{
// 					id:     2,
// 					parent: 1,
// 					URL:    ptrStr("url2"),
// 					Title:  "title2",
// 					Folder: "title2",
// 				},
// 			},
// 			expected: []Bookmark{
// 				{
// 					id:     2,
// 					parent: 1,
// 					URL:    ptrStr("url2"),
// 					Title:  "title2",
// 					Folder: "title1/title2",
// 				},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			d := &DatabaseOperator{
// 				bookmarks: tt.bookmarks,
// 			}
// 			d.NormalizeBookmarks()
// 			assert.Equal(t, tt.expected, d.bookmarks, "mismatch")
// 		})
// 	}
// }

// func ptrStr(s string) *string {
// 	return &s
// }
