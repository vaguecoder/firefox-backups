package encoding

import (
	"context"
	"fmt"
	"reflect"

	"github.com/vaguecoder/firefox-backups/pkg/bookmark"
	"github.com/vaguecoder/firefox-backups/pkg/logs"
)

// EncodingManager holds bookmarks and
// target encoders to parse bookmarks to
type EncodingManager struct {
	encoders  []Encoder
	bookmarks []bookmark.Bookmark
	logger    logs.Logger
}

// NewEncoderManager initiates new EncodingManager
func NewEncoderManager(ctx context.Context) *EncodingManager {
	return &EncodingManager{
		encoders:  []Encoder{},
		bookmarks: nil,
		logger:    logs.FromContext(ctx),
	}
}

// Encoder appends input encoder to encoder list in manager.
// Both receiver and return value are of same type to implement builder's pattern.
func (e *EncodingManager) Encoder(encoder Encoder) *EncodingManager {
	if reflect.ValueOf(encoder).IsNil() {
		// When encoder is empty.
		// This is a possible case when caller initializes
		// encoder as interface and doesn't assign a value.
		e.logger.Info().Msg("Empty encoder")

		return e
	}

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
func (e *EncodingManager) Write() error {
	if e.bookmarks == nil {
		// When no bookmarks provided
		return fmt.Errorf("bookmarks missing in chaining")
	}

	if len(e.encoders) == 0 || len(e.bookmarks) == 0 {
		// When no encoders provided or no bookmarks in the list.
		// Skipping.
		return nil
	}

	var (
		err       error
		encoder   Encoder
		subLogger logs.Logger
	)

	// Iterate over encoders in manager
	for _, encoder = range e.encoders {
		// Sub-logger to hold current encoder's filename and encoder name
		subLogger = logs.FromRawLogger(e.logger.With().Str("filename", encoder.Filename()).
			Stringer("encoder", encoder).Logger())

		if err = encoder.Encode(e.bookmarks); err != nil {
			// When encountered error while encoding
			subLogger.Error().Err(err).Msg("Failed to encode")

			return fmt.Errorf("failed to encode to %q: %v", encoder, err)
		}

		// When encoding is successful
		subLogger.Info().Msg("Successfully encoded to output stream/file")
	}

	return nil
}
