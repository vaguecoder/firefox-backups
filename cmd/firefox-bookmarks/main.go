package main

import (
	"context"
	"encoding/json"
	"io"
	"os"

	_ "github.com/mattn/go-sqlite3"

	db "github.com/vaguecoder/firefox-backups/pkg/database"
	"github.com/vaguecoder/firefox-backups/pkg/database/sqlite"
	"github.com/vaguecoder/firefox-backups/pkg/files"
	"github.com/vaguecoder/firefox-backups/pkg/flags"
	"github.com/vaguecoder/firefox-backups/pkg/logs"
)

const (
	placesDBFile = `places.sqlite`
	stdout       = `stdout`
)

func main() {
	ctx := context.Background()
	logger := logs.FromContext(ctx)
	flagOps := flags.NewOperator(os.Args[1:])

	f, err := flagOps.Parse()
	if err != nil {
		logger.Fatal().Err(err).Strs("args", os.Args).Msg("Failed to parse flags from command line args")
	}

	if f.Silent {
		ctx, logger = logs.SilentLogger(ctx)
	}

	logger.Info().Interface("flags", f).Msg("Input flags")

	var outputFile io.Writer = os.Stdout
	var outputFilename = stdout
	cpOps := files.NewOperator(logger)

	if f.OutputFilename != nil {
		outputFilename = *f.OutputFilename

		outputFile, err = cpOps.Open(*f.OutputFilename)
		if err != nil {
			logger.Fatal().Err(err).Str("outputJSONFile", *f.OutputFilename).
				Msg("Failed to open file")
		}
	}

	if f.SQLiteDBFilename != placesDBFile {
		if err = cpOps.Copy(f.SQLiteDBFilename, placesDBFile); err != nil {
			logger.Fatal().Err(err).Str("input-sqlite-file", f.SQLiteDBFilename).
				Str("temp-file", placesDBFile).Msg("Failed to copy file")
		}

		defer func() {
			if err := cpOps.Delete(placesDBFile); err != nil {
				logger.Fatal().Err(err).Str("temp-file", placesDBFile).Msg("Failed to temp files")
			}
		}()
	}

	dbConn, err := sqlite.NewDB(placesDBFile)
	if err != nil {
		logger.Fatal().Err(err).Str("db-filename", placesDBFile).Msg("Failed to open DB connection")
	}

	dbOps := db.NewDatabaseOperator(dbConn, f.RawOutput, f.IgnoreDefaults)
	bookmarks, err := dbOps.GetBookmarks(ctx)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to fetch bookmarks from db")
	}
	logger.Info().Int("count", len(bookmarks)).Msg("Count of bookmarks fetched")

	jsonEncoder := json.NewEncoder(outputFile)
	jsonEncoder.SetIndent("", "\t")
	if err := jsonEncoder.Encode(bookmarks); err != nil {
		logger.Fatal().Err(err).Str("output-stream", outputFilename).
			Msg("Failed to JSON marshal to output stream")
	}
	logger.Info().Str("output-stream", outputFilename).
		Msg("Successfully JSON marshalled and written to output stream")
}
