SOURCES = $(shell find . -type f -name '*.go')
BUILDTAGS = remote,exclude_graphdriver_btrfs,btrfs_noversion,exclude_graphdriver_devicemapper,containers_image_openpgp

.PHONY: default
default: all

.PHONY: all
all: darwin linux windows

.PHONY: darwin
darwin: darwin-amd64 darwin-arm64

.PHONY: darwin-amd64
darwin-amd64: $(SOURCES) go.mod go.sum
	fyne-cross darwin -arch amd64 -tags $(BUILDTAGS)

.PHONY: darwin-arm64
darwin-arm64: $(SOURCES) go.mod go.sum
	fyne-cross darwin -arch arm64 -tags $(BUILDTAGS)

.PHONY: linux
linux: linux-amd64 linux-arm64

.PHONY: linux-amd64
linux-amd64: $(SOURCES) go.mod go.sum
	fyne-cross linux -arch amd64 -tags $(BUILDTAGS)

.PHONY: linux-arm64
linux-arm64: $(SOURCES) go.mod go.sum
	fyne-cross linux -arch arm64 -tags $(BUILDTAGS)

.PHONY: windows
windows: windows-amd64

.PHONY: windows-amd64
windows-amd64: $(SOURCES) go.mod go.sum
	fyne-cross windows -arch amd64 -tags $(BUILDTAGS)

.PHONY: clean
clean:
	rm -rf fyne-cross
