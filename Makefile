APP?=xnotify
ARCH?=amd64
BUILD_DIR?=bin
LOG_DIR?=doc/log/build_history
OS?=linux
PKGER_DIR?=pkg/read
VERSION?=v0.0.1

FILE_ARCH=$(OS)_$(ARCH)
BUILD_PATH=$(BUILD_DIR)/$(VERSION)/$(FILE_ARCH)
INSTALL_PATH=$(HOME)/.local/bin/
LOG_PATH=$(LOG_DIR)/$(VERSION)/$(FILE_ARCH)

.PHONY: pkger security tidy vendor build clean version install

pkger:
	@pkger -o $(PKGER_DIR)
	@echo "[OK] Package embedded files completed!"

security:
	@gosec ./...
	@echo "[OK] Go security check completed!"

tidy:
	@go mod tidy
	@echo "[OK] Go modules tidy completed!"

vendor:
	@go mod vendor
	@echo "[OK] Vendored code import completed!"

build: pkger tidy vendor security
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
	@echo "[OK] Binary created successfully!"

clean:
	@rm -rf $(BUILD_PATH)
	@echo "[OK] Release $(APP)-$(VERSION)-$(FILE_ARCH) cleaned!"

version:
	@echo $(VERSION)

install:
	install -d -m 755 '$(INSTALL_PATH)'
	install -Z $(BUILD_PATH)/$(APP) $(INSTALL_PATH)/$(APP)
	@echo "[OK] Binary installed successfully!"
