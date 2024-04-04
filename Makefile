
.PHONY: example install-*

example: install-servoc install-plugins
	go generate ./example

install-servoc:
	go generate ./internal/...
	go install ./cmd/servoc

install-plugins:
	go install ./cmd/servoc-ext_goecho
	go install ./cmd/servoc-ext_gotype
