#!/usr/bin/env make

.PHONY: dev-setup
dev-setup:
	go install github.com/vektra/mockery/v2@latest

.PHONY: build-firefox-bookmarks
build-firefox-bookmarks:
	@ go build ./cmd/firefox-bookmarks

.PHONY: run-firefox-bookmarks
run-firefox-bookmarks: build-firefox-bookmarks
	@ ./firefox-bookmarks \
		--input-sqlite-file $(shell find ~/.var/app/org.mozilla.firefox/ -name 'places.sqlite') \
		--raw=false \
		--ignore-defaults \
		--denormalize \
		--silent=false \
		--stdout-format="" \
		--output-files=yaml:firefox-backups.yaml,json:firefox-backups.json,table:firefox-backups.txt,csv:firefox-backups.csv,json:firefox-backups-copy.json

.PHONY: unit-test
unit-test:
	go test -v -race -coverprofile cover.out ./...
	go tool cover -func cover.out

.PHONY: generate-mocks
generate-mocks:
	rm -rf pkg/mocks/*
	mockery --all --output=pkg/mocks --dir ./pkg