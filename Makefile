REPO_PREFIX := github.com/zexi/ncmbox
VERSION_PKG := $(REPO_PREFIX)/pkg/version
ROOT_DIR := $(CURDIR)
BUILD_DIR := $(ROOT_DIR)/_output
BIN_DIR := $(BUILD_DIR)/bin

GIT_COMMIT := $(shell git rev-parse --short HEAD)
GIT_BRANCH := $(shell git name-rev --name-only HEAD)
GIT_VERSION := $(shell git describe --tags --abbrev=14 $(GIT_COMMIT)^{commit})
GIT_TREE_STATE := $(shell s=`git status --porcelain 2>/dev/null`; if [ -z "$$s" ]; then echo "clean"; else echo "dirty"; fi)
BUILD_DATE := $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

LDFLAGS := "-w \
	-X $(VERSION_PKG).gitVersion=$(GIT_VERSION) \
	-X $(VERSION_PKG).gitCommit=$(GIT_COMMIT) \
	-X $(VERSION_PKG).gitBranch=$(GIT_BRANCH) \
	-X $(VERSION_PKG).buildDate=$(BUILD_DATE) \
	-X $(VERSION_PKG).gitTreeState=$(GIT_TREE_STATE) \
	-X $(VERSION_PKG).gitMajor=0 \
	-X $(VERSION_PKG).gitMinor=0"

GO_BUILD := go build -ldflags $(LDFLAGS)

build:
	$(GO_BUILD) -o $(BIN_DIR)/ncmbox $(REPO_PREFIX)/cmd/ncmbox

.PHONY: build
