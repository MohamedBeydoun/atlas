VERSION := $(shell bash scripts/version.sh; cat VERSION)

.PHONY: version test bin build coverage

test:
	go test -cover ./pkg/...

version:
	bash scripts/version.sh

bin:
	mkdir -p bin
	GOOS=linux GOARCH=amd64 go build -o bin/atlas-$(VERSION)-linux-amd64
	GOOS=windows GOARCH=amd64 go build -o bin/atlas-$(VERSION)-windows-amd64
	GOOS=darwin GOARCH=amd64 go build -o bin/atlas-$(VERSION)-darwin-amd64
