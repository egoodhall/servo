
.PHONY: example servoc plugins

default: servoc plugins

example: servoc plugins
	go generate ./example

servoc:
	go generate ./internal/...
	go install ./cmd/servoc

plugins:
	go install ./cmd/servoc-ext_goecho
	go install ./cmd/servoc-ext_gotype
