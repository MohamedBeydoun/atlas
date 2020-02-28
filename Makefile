test:
	go test -cover ./pkg/...

.PHONY: version
version:
	bash scripts/version.sh