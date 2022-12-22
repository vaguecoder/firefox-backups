#!/usr/bin/env make

.PHONY: build-firefox-bookmarks
build-firefox-bookmarks:
	@ go build ./cmd/firefox-bookmarks

.PHONY: run-firefox-bookmarks
run-firefox-bookmarks: build-firefox-bookmarks
	@ ./firefox-bookmarks \
		--sqlite-filename $(shell find ~/.var/app/org.mozilla.firefox/ -name 'places.sqlite') \
		--raw=false \
		--ignore-defaults \
		--write-to-file \
		--output-filename "firefox-bookmarks.json" \
		--silent=false

.PHONY: unit-test
unit-test:
	go test -v -race -coverprofile cover.out ./...
	go tool cover -func cover.out