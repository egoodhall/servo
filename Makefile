
.PHONY: example servoc plugins

default: servoc plugins

example: servoc plugins
	go generate ./example/...

servoc:
	go generate ./...
	go install ./cmd/servoc

plugins:
	go install ./cmd/servoc-ext_gotype
	go install ./cmd/servoc-ext_gohttp
