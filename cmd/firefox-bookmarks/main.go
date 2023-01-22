package main

import (
	"context"
	"io"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/vaguecoder/firefox-backups/pkg/bookmark"
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
		err                                                  error
		denormalizeOps, ignoredefaultsOps                    filters.Filter
		csvEncoder, jsonEncoder, tabularEncoder, yamlEncoder pkgEncoding.Encoder
		logger                                               logs.Logger
		inputFlags                                           *flags.Flags
		dbConn                                               sqlite.DBConnection
		dbOps                                                db.BookmarkOperator
		bookmarks                                            []bookmark.Bookmark

		ctx                     = context.Background()
		flagOps                 = flags.NewOperator(os.Args[1:])
		outputStream  io.Writer = os.Stdout
		discardStream io.Writer = io.Discard
		enableHeader  bool      = true
	)

	ctx, logger = logs.NewLogger(ctx, os.Stdout, logs.LevelInfo)

	inputFlags, err = flagOps.Parse()
	if err != nil {
		logger.Fatal().Err(err).Strs("args", os.Args).Msg("Failed to parse flags from command line args")
	}

	if inputFlags.Silent {
		ctx, logger = logs.SilentLogger(ctx)
	}

	logger.Info().Interface("flags", inputFlags).Msg("Input flags")

	cpOps := files.NewOperator(logger)

	if inputFlags.SQLiteDBFilename != placesDBFile {
		if err = cpOps.Copy(inputFlags.SQLiteDBFilename, placesDBFile); err != nil {
			logger.Fatal().Err(err).Str("input-sqlite-file", inputFlags.SQLiteDBFilename).
				Str("temp-file", placesDBFile).Msg("Failed to copy file")
		}

		defer func() {
			if err := cpOps.Delete(placesDBFile); err != nil {
				logger.Fatal().Err(err).Str("temp-file", placesDBFile).Msg("Failed to temp files")
			}
		}()
	}

	dbConn, err = sqlite.NewDB(placesDBFile)
	if err != nil {
		logger.Fatal().Err(err).Str("db-filename", placesDBFile).Msg("Failed to open DB connection")
	}

	dbOps = db.NewDatabaseOperator(dbConn, inputFlags.RawOutput, inputFlags.FilterIgnoreDefaults)
	bookmarks, err = dbOps.GetBookmarks(ctx)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to fetch bookmarks from db")
	}
	logger.Info().Int("count", len(bookmarks)).Msg("Count of bookmarks fetched")

	if inputFlags.FilterDenormalize {
		denormalizeOps = &denormalize.Denormalizer{}
	}

	if inputFlags.FilterIgnoreDefaults {
		ignoredefaultsOps = &ignoredefaults.DefaultsRemover{}
	}

	bookmarks, err = filters.NewFilterManager().
		Bookmarks(bookmarks).
		Filter(denormalizeOps).
		Filter(ignoredefaultsOps).
		Apply(ctx)
	if err != nil {
		logger.Fatal().Err(err).
			Msg("Failed to apply filter(s)")
	}

	csvEncoder = pkgEncodingCSV.NewEncoder(discardStream, enableHeader)
	jsonEncoder = pkgEncodingJSON.NewEncoder(outputStream)
	tabularEncoder = pkgEncodingTab.NewEncoder(discardStream, enableHeader)
	yamlEncoder = pkgEncodingYAML.NewEncoder(discardStream)

	err = pkgEncoding.NewEncoderManager().
		Bookmarks(bookmarks).
		Encoder(csvEncoder).
		Encoder(jsonEncoder).
		Encoder(tabularEncoder).
		Encoder(yamlEncoder).
		Write(ctx)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to encode bookmarks to output stream(s)")
	}
}
