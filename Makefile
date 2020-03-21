APP?=xnotify
ARCH?=amd64
BUILD_DIR?=bin
LOG_DIR?=doc/log/build_history
OS?=linux
PKGER_DIR?=pkg/read
PKGER_INCLUDES?=-include /pkg/tmpl
VERSION?=v0.0.1

FILE_ARCH=$(OS)_$(ARCH)
BUILD_PATH=$(BUILD_DIR)/$(VERSION)/$(FILE_ARCH)
INSTALL_PATH=$(HOME)/.local/bin/
LOG_PATH=$(LOG_DIR)/$(VERSION)/$(FILE_ARCH)

.PHONY: dirprep pkger security tidy vendor build clean version install

default: clean build

dirprep:
	@mkdir -p $(BUILD_PATH)
	@mkdir -p $(LOG_PATH) && touch $(LOG_PATH)/make.log

pkger:
	@pkger $(PKGER_INCLUDES) -o $(PKGER_DIR) 2>&1 \
	  | tee -a $(BUILD_PATH)/make.log
	@echo "[OK] Package embedded files completed!" \
	  | tee -a $(BUILD_PATH)/make.log

security:
	@gosec ./... 2>&1 \
	  | tee -a $(BUILD_PATH)/make.log
	@echo "[OK] Go security check completed!" \
	  | tee -a $(BUILD_PATH)/make.log

tidy:
	@go mod tidy 2>&1 \
	  | tee -a $(BUILD_PATH)/make.log
	@echo "[OK] Go modules tidy completed!" \
	  | tee -a $(BUILD_PATH)/make.log

vendor:
	@go mod vendor 2>&1 \
	  | tee -a $(BUILD_PATH)/make.log
	@echo "[OK] Vendored code import completed!" \
	  | tee -a $(BUILD_PATH)/make.log

build: dirprep pkger tidy vendor security
	@GOARCH=$(ARCH) GOOS=$(OS) \
	  go build \
	  -a \
	  -v \
	  -x \
	  -trimpath \
	  -mod=vendor \
	  -o=$(BUILD_PATH)/$(APP) \
	  -ldflags "-X 'main.VERSION=$(VERSION)' -X 'main.COMMIT=$$(git rev-parse --verify HEAD)' -X 'main.APP=$(APP)'" 2>&1 \
	  | tee -a $(BUILD_PATH)/make.log
	@cp $(BUILD_PATH)/make.log $(LOG_PATH)/make.log 2>&1 \
	  | tee -a $(BUILD_PATH)/make.log
	@echo "[OK] Binary created successfully!" 2>&1 \
	  | tee -a $(BUILD_PATH)/make.log

clean:
	@rm -rf $(BUILD_PATH)
	@>$(LOG_PATH)/make.log
	@echo "[OK] Release $(APP)-$(VERSION)-$(FILE_ARCH) cleaned!"

version:
	@echo $(VERSION)

install:
	install -d -m 755 '$(INSTALL_PATH)'
	install -Z $(BUILD_PATH)/$(APP) $(INSTALL_PATH)/$(APP)
	@echo "[OK] Binary installed successfully!"
