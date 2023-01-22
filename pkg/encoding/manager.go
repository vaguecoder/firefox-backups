package encoding

import (
	"context"
	"fmt"

	"github.com/vaguecoder/firefox-backups/pkg/bookmark"
)

type encodingManager struct {
	encoders  []Encoder
	bookmarks []bookmark.Bookmark
}

func NewEncoderManager() *encodingManager {
	return &encodingManager{
		encoders: []Encoder{},
	}
}

func (e *encodingManager) Encoder(encoder Encoder) *encodingManager {
	e.encoders = append(e.encoders, encoder)

	return e
}

func (e *encodingManager) Bookmarks(bookmarks []bookmark.Bookmark) *encodingManager {
	e.bookmarks = bookmarks

	return e
}

func (f *encodingManager) Write(ctx context.Context) error {
	if f.bookmarks == nil {
		return fmt.Errorf("bookmarks missing in chaining")
	}

	var err error

	for _, encoder := range f.encoders {
		if encoder == nil {
			continue
		}

		err = encoder.Encode(f.bookmarks)
		if err != nil {
			return fmt.Errorf("failed to encode to %q: %v", encoder, err)
		}
	}

	return nil
}
