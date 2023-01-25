package encoding

import (
	"context"
	"fmt"

	"github.com/vaguecoder/firefox-backups/pkg/bookmark"
	"github.com/vaguecoder/firefox-backups/pkg/logs"
)

// EncodingManager holds bookmarks and
// target encoders to parse bookmarks to
type EncodingManager struct {
	encoders  []Encoder
	bookmarks []bookmark.Bookmark
}

// NewEncoderManager initiates new EncodingManager
func NewEncoderManager() *EncodingManager {
	return &EncodingManager{
		encoders: []Encoder{},
	}
}

// Encoder appends input encoder to encoder list in manager.
// Both receiver and return value are of same type to implement builder's pattern.
func (e *EncodingManager) Encoder(encoder Encoder) *EncodingManager {
	e.encoders = append(e.encoders, encoder)

	return e
}

// Encoder appends bookmarks to manager.
// Both receiver and return value are of same type to implement builder's pattern.
func (e *EncodingManager) Bookmarks(bookmarks []bookmark.Bookmark) *EncodingManager {
	e.bookmarks = bookmarks

	return e
}

// Write encodes the bookmarks to all the encoders added to manager
func (f *EncodingManager) Write(ctx context.Context) error {
	if f.bookmarks == nil {
		// When no bookmarks provided
		return fmt.Errorf("bookmarks missing in chaining")
	}

	if len(f.encoders) == 0 {
		// When no encoders provided.
		// Skipping.
		return nil
	}

	var (
		err       error
		encoder   Encoder
		subLogger logs.Logger

		logger = logs.FromContext(ctx).With().Logger()
	)

	// Iterate over encoders in manager
	for _, encoder = range f.encoders {
		// Sub-logger to hold current encoder's filename and encoder name
		subLogger = logs.FromRawLogger(logger.With().Str("filename", encoder.Filename()).
			Stringer("encoder", encoder).Logger())

		if encoder == nil {
			// When encoder is empty.
			// This is a possible case when caller initializes
			// encoder as interface and doesn't assign a value.
			subLogger.Info().Msg("Encoder empty")

			continue
		}

		err = encoder.Encode(f.bookmarks)
		if err != nil {
			// When encountered error while encoding
			subLogger.Error().Err(err).Msg("Failed to encode")

			return fmt.Errorf("failed to encode to %q: %v", encoder, err)
		}

		// When encoding is successful
		subLogger.Info().Msg("Successfully encoded to output stream/file")
	}

	return nil
}
