.PHONY: build clean test dist serve run-compose-test
VERSION := $(shell git describe --always |sed -e "s/^v//")

build:
	@echo "Compiling source"
	@mkdir -p build
	go build $(GO_EXTRA_BUILD_ARGS) -ldflags "-s -w -X main.version=$(VERSION)" -o build/mioty-bssci-adapter cmd/main.go

clean:
	@echo "Cleaning up workspace"
	@rm -rf build
	@rm -rf dist

test:
	@echo "Running tests"
	@rm -f coverage.out
	@staticcheck ./...
	@go vet ./...
	@go test -cover -coverprofile coverage.out -p 1 ./...

dist:
	@goreleaser

snapshot:
	@goreleaser --snapshot --clean

dev-requirements:
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/goreleaser/goreleaser/v2@latest
	go install github.com/goreleaser/nfpm/v2/cmd/nfpm@latest


# shortcuts for development

serve: build
	./build/mioty-bssci-adapter

run-compose-test:
	docker compose run --rm mioty-bssci-adapter make test
