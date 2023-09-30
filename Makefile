SOURCES = $(shell find . -type f -name '*.go')
BUILDTAGS = remote exclude_graphdriver_btrfs btrfs_noversion exclude_graphdriver_devicemapper containers_image_openpgp

.PHONY: default
default: all

.PHONY: all
all: build

.PHONY: build
build: build/viuws

build/viuws: $(SOURCES) go.mod go.sum
	go build -tags "$(BUILDTAGS)" -o $@ ./cmd/viuws

.PHONY: run
run: $(SOURCES) go.mod go.sum
	go run -tags "$(BUILDTAGS)" ./cmd/viuws/main.go

.PHONY: dist
dist: dist/darwin-amd64/viuws dist/darwin-arm64/viuws dist/linux-amd64/viuws dist/linux-arm64/viuws dist/windows-amd64/viuws.exe

dist/darwin-amd64/viuws: $(SOURCES) go.mod go.sum
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -tags "$(BUILDTAGS)" -o $@ ./cmd/viuws

dist/darwin-arm64/viuws: $(SOURCES) go.mod go.sum
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -tags "$(BUILDTAGS)" -o $@ ./cmd/viuws

dist/linux-amd64/viuws: $(SOURCES) go.mod go.sum
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags "$(BUILDTAGS)" -o $@ ./cmd/viuws

dist/linux-arm64/viuws: $(SOURCES) go.mod go.sum
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -tags "$(BUILDTAGS)" -o $@ ./cmd/viuws

dist/windows-amd64/viuws.exe: $(SOURCES) go.mod go.sum
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -tags "$(BUILDTAGS)" -o $@ ./cmd/viuws

.PHONY: clean
clean:
	rm -rf build/ dist/
