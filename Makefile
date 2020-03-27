# Vars, cannot be changed via environment
MINIMUM_GO_VERSION=1.13

# Vars, can be changed via the environment
APP?=xnotify
ARCH?=amd64
BOX_BLOBFILE=pkg/box/blob.go
BUILD_DIR?=bin
HACK_DIR?=hack
INSTALL_PATH=/usr/local/bin
LOG_DIR?=doc/log/build_history
OS?=linux
RELEASE?=false

# Vars, automatically populated
COMMIT=$(shell git rev-parse --verify 'HEAD^{commit}')
CURR_GO_VERSION=$(shell go version | grep go | cut -d " " -f 3)
VERSION=$(shell git describe --always --abbrev=40 --dirty)
SEMVER=$(shell git describe --always --abbrev=0)

# Complex Vars, built from the variables above and environment variables
FILE_ARCH=$(OS)_$(ARCH)
BUILD_PATH=$(BUILD_DIR)/$(SEMVER)/$(FILE_ARCH)
LOG_PATH=$(LOG_DIR)/$(SEMVER)/$(FILE_ARCH)


.PHONY: default build clean code-prep commit dir-prep doc go-generate go-goimports go-lint go-sec go-test go-tidy go-vendor go-verify go-vet install release shellcheck version yaml-lint
.EXPORT_ALL_VARIABLES:

default: clean doc code-prep build

# Build binary
build: dir-prep 
	@hack/build.sh

# Delete old build directory, clear out old log
clean:
	@hack/clean.sh

# Prepare the code for building
code-prep: go-tidy go-vendor go-goimports go-lint go-generate go-vet go-verify go-sec shellcheck yaml-lint

# Create a new commit for all of the changes in the current directory
commit: clean dir-prep code-prep doc
	@hack/commit.sh

# Prepare build and log directories
dir-prep:
	@hack/dir-prep.sh

# Build templated documentation
doc:
	@hack/doc.sh

# Generate build-time dependencies
go-generate:
	@hack/go-generate.sh

# Format code and resolve imports
go-goimports:
	@hack/go-goimports.sh

# Lint code
go-lint:
	@hack/go-lint.sh

# Run static security analysis
go-sec:
	@hack/go-sec.sh

# Run tests
go-test:
	@hack/go-test.sh

# Tidy modules
go-tidy:
	@hack/go-tidy.sh

# Dump imports into vendor directory
go-vendor:
	@hack/go-vendor.sh

# Verify modules
go-verify:
	@hack/go-verify.sh

# Vet code
go-vet:
	@hack/go-vet.sh

# Install binary
install:
	@hack/install.sh

# Build release
release: clean dir-prep code-prep doc
	@hack/release.sh $(VERSION)

# Run shellcheck
shellcheck:
	@hack/shellcheck.sh

# Output version
version:
	@echo $(VERSION)

# Run YAML linter
yaml-lint:
	@hack/yaml-lint.sh
