APP?=xnotify
ARCH?=amd64
BUILD_DIR?=bin
LOG_DIR?=doc/log/build_history
OS?=linux
VERSION?=v0.0.1

FILE_ARCH=$(OS)_$(ARCH)
BUILD_PATH=$(BUILD_DIR)/$(VERSION)/$(FILE_ARCH)
INSTALL_PATH=$(HOME)/.local/bin/
LOG_PATH=$(LOG_DIR)/$(VERSION)/$(FILE_ARCH)

.PHONY: tidy vendor build clean version install

tidy:
	@go mod tidy

vendor:
	@go mod vendor

build: tidy vendor
	@GOARCH=$(ARCH) GOOS=$(OS) \
	  mkdir -p $(BUILD_PATH) && \
	  go build \
	  -a \
	  -v \
	  -x \
	  -trimpath \
	  -mod=vendor \
	  -o=$(BUILD_PATH)/$(APP) \
	  -ldflags "-X 'main.VERSION=$(VERSION)' -X 'main.COMMIT=$$(git rev-parse --verify HEAD)' -X 'main.APP=$(APP)'" \
	  > $(BUILD_PATH)/make.log 2>&1 && \
	  mkdir -p $(LOG_PATH) && \
	  cp $(BUILD_PATH)/make.log $(LOG_PATH)/make.log

clean:
	@rm -rf $(BUILD_DIR)

version:
	@echo $(VERSION)

install:
	install -d -m 755 '$(INSTALL_PATH)'
	install -Z $(BUILD_PATH)/$(APP) $(INSTALL_PATH)/$(APP)
