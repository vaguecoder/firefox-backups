package main

import (
	"context"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/vaguecoder/firefox-backups/pkg/bookmark"
	"github.com/vaguecoder/firefox-backups/pkg/constants"
	db "github.com/vaguecoder/firefox-backups/pkg/database"
	"github.com/vaguecoder/firefox-backups/pkg/database/sqlite"
	pkgEncoding "github.com/vaguecoder/firefox-backups/pkg/encoding"
	pkgEncodingCSV "github.com/vaguecoder/firefox-backups/pkg/encoding/csv"
	pkgEncodingJSON "github.com/vaguecoder/firefox-backups/pkg/encoding/json"
	pkgEncodingTab "github.com/vaguecoder/firefox-backups/pkg/encoding/tabular"
	pkgEncodingYAML "github.com/vaguecoder/firefox-backups/pkg/encoding/yaml"
	"github.com/vaguecoder/firefox-backups/pkg/files"
	"github.com/vaguecoder/firefox-backups/pkg/filters"
	"github.com/vaguecoder/firefox-backups/pkg/filters/denormalize"
	ignoredefaults "github.com/vaguecoder/firefox-backups/pkg/filters/ignore-defaults"
	"github.com/vaguecoder/firefox-backups/pkg/flags"
	"github.com/vaguecoder/firefox-backups/pkg/logs"
)

const (
	placesDBFile = `places.sqlite`
	stdout       = `stdout`
)

func main() {
	var (
		err                               error
		denormalizeOps, ignoredefaultsOps filters.Filter
		encoder                           pkgEncoding.Encoder
		encoderManager                    *pkgEncoding.EncodingManager
		logger                            logs.Logger
		inputFlags                        *flags.Flags
		outputFile                        files.File
		dbConn                            sqlite.DBConnection
		dbOps                             db.BookmarkOperator
		bookmarks                         []bookmark.Bookmark
		fileOps                           files.FileOperator
		outputFileSet                     flags.OutputFile

		ctx          = context.Background()
		enableHeader = true

		// Initialize flag operator with input arguments
		flagOps = flags.NewOperator(os.Args[1:])
	)

	// Create new logger and add to context for easy propagation
	ctx, logger = logs.NewLogger(ctx, os.Stdout, logs.LevelInfo)

	// Read the input flags
	inputFlags, err = flagOps.Parse()
	if err != nil {
		logger.Fatal().Err(err).Strs("args", os.Args).Msg("Failed to parse flags from command line args")
	}

	// When silent mode is enabled in input flags, replace logger with silent logger.
	if inputFlags.Silent {
		ctx, logger = logs.SilentLogger(ctx)
	}

	// Log input flag values
	logger.Info().Interface("flags", inputFlags).Msg("Input flags")

	// Initiate files operator for file creation, copying, deletion, etc.
	fileOps = files.NewOperator(ctx)

	if inputFlags.SQLiteDBFilename != placesDBFile {
		// When input sqlite DB file is not same as places.sqlite,
		// i.e., when input file either has a non-default name, or
		// it is in a different location, the file has to be copied to operate,
		// to avoid conflicts, locking or corruption.
		if err = fileOps.Copy(inputFlags.SQLiteDBFilename, placesDBFile); err != nil {
			// When copying of input file as places.sqlite failed
			logger.Fatal().Err(err).Str("input-sqlite-file", inputFlags.SQLiteDBFilename).
				Str("temp-file", placesDBFile).Msg("Failed to copy file")
		}

		// Delete the copied file (and matching files) after completion or failure
		defer func() {
			if err = fileOps.Delete(placesDBFile); err != nil {
				// When deletion of copied input file failed
				logger.Fatal().Err(err).Str("temp-file", placesDBFile).Msg("Failed to temp files")
			}
		}()
	}

	// Initiate database connection
	dbConn, err = sqlite.NewDB(placesDBFile)
	if err != nil {
		// When initialization of database connection failed
		logger.Fatal().Err(err).Str("db-filename", placesDBFile).Msg("Failed to open DB connection")
	}

	// Initiate database operator
	dbOps = db.NewDatabaseOperator(dbConn)
	bookmarks, err = dbOps.GetBookmarks(ctx)
	if err != nil {
		// When initialization of database operator failed
		logger.Fatal().Err(err).Msg("Failed to fetch bookmarks from db")
	}

	// When initialization of database operator was successful
	logger.Info().Int("count", len(bookmarks)).Msg("Count of bookmarks fetched")

	// Filters
	if inputFlags.FilterDenormalize {
		// When denormalize filter is enabled
		denormalizeOps = &denormalize.Denormalizer{}
	}
	if inputFlags.FilterIgnoreDefaults {
		// When ignore-defaults filter is enabled
		ignoredefaultsOps = &ignoredefaults.DefaultsRemover{}
	}

	// Filter the resultant bookmarks
	bookmarks, err = filters.NewFilterManager().
		Bookmarks(bookmarks).
		Filter(denormalizeOps).
		Filter(ignoredefaultsOps).
		Apply(ctx)
	if err != nil {
		// When filtering failed
		logger.Fatal().Err(err).Msg("Failed to apply filter(s)")
	}

	// Initialize encoder manager and add bookmarks to manager
	encoderManager = pkgEncoding.NewEncoderManager().Bookmarks(bookmarks)

	if inputFlags.StdOutFormat != nil {
		// When stdout printer is also enabled
		encoderManager = encoderManager.Encoder(inputFlags.StdOutFormat)
	}

	// Iterate over list of input format-filename flag value sets
	for _, outputFileSet = range inputFlags.OutputFiles {
		// Create output file
		outputFile, err = fileOps.Open(outputFileSet.Filename)
		if err != nil {
			// When creation of output file failed
			logger.Fatal().Err(err).Str("output-filename", outputFileSet.Filename).
				Stringer("output-format", outputFileSet.Format).Msg("Failed to create output file")
		}

		// Map output file format against the encoder type
		switch outputFileSet.Format {
		case constants.CSVFormat:
			// CSV format
			encoder = pkgEncodingCSV.NewEncoder(outputFile, enableHeader)
		case constants.JSONFormat:
			// JSON format
			encoder = pkgEncodingJSON.NewEncoder(outputFile)
		case constants.TabularFormat:
			// Table format
			encoder = pkgEncodingTab.NewEncoder(outputFile, enableHeader)
		case constants.YAMLFormat:
			// YAML format
			encoder = pkgEncodingYAML.NewEncoder(outputFile)
		default:
			// Input format is already validated at input flags
		}

		// Append the encoder to the list in manager
		encoderManager = encoderManager.Encoder(encoder)
	}

	// Encode bookmarks against all output formats (stdout or file formats)
	err = encoderManager.Write(ctx)
	if err != nil {
		// When encoding bookmarks failed
		logger.Fatal().Err(err).Msg("Failed to encode bookmarks to output stream(s)")
	}
}
