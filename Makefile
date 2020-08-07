VERSION := $(shell bash scripts/version.sh; cat VERSION)

.PHONY: version test bin release dev

test:
	@echo "Running tests..."
	@go test -cover ./pkg/...

version:
	@bash scripts/version.sh
	@echo "Version: $(shell cat VERSION)"

bin:
	@mkdir -p bin
	@GOOS=linux GOARCH=amd64 go build -o bin/atlas-$(VERSION)-linux-amd64 -ldflags "-X github.com/MohamedBeydoun/atlas/cmd.version=$(shell cat VERSION)"
	@echo "Generated bin/atlas-$(VERSION)-linux-amd64"
	@GOOS=windows GOARCH=amd64 go build -o bin/atlas-$(VERSION)-windows-amd64 -ldflags "-X github.com/MohamedBeydoun/atlas/cmd.version=$(shell cat VERSION)"
	@echo "Generated bin/atlas-$(VERSION)-windows-amd64"
	@GOOS=darwin GOARCH=amd64 go build -o bin/atlas-$(VERSION)-darwin-amd64 -ldflags "-X github.com/MohamedBeydoun/atlas/cmd.version=$(shell cat VERSION)"
	@echo "Generated bin/atlas-$(VERSION)-darwin-amd64"

release: version
	@bash scripts/release.sh $(shell cat VERSION)
	@echo "Release complete!"

dev: version
	@go install -ldflags "-X github.com/MohamedBeydoun/atlas/cmd.version=$(shell cat VERSION)"