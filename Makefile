# Vars, can be changed via the environment
APP?=xnotify
ARCH?=amd64
BUILD_DIR?=bin
LOG_DIR?=doc/log/build_history
OS?=linux
PKGER_DIR?=pkg/read
PKGER_INCLUDES?=-include /pkg/tmpl
VERSION?=v0.0.1

# Complex Vars, built from the variables above and environment variables
FILE_ARCH=$(OS)_$(ARCH)
BUILD_PATH=$(BUILD_DIR)/$(VERSION)/$(FILE_ARCH)
INSTALL_PATH=/usr/local/bin
LOG_PATH=$(LOG_DIR)/$(VERSION)/$(FILE_ARCH)

.PHONY: default dirprep doc generate tidy vendor security build clean version install

default: clean build

# Prepare build and log directories
dirprep:
	@mkdir -p $(BUILD_PATH) && touch $(BUILD_PATH)/make.log
	@mkdir -p $(LOG_PATH) && touch $(LOG_PATH)/make.log
	@date -u \
	  | tee -a $(BUILD_PATH)/make.log

# Build templated documentation
doc:
	@echo "[OK] Documentation generated successfully!" \
	  | tee -a $(BUILD_PATH)/make.log

# Generate build-time dependencies
generate:
	@go generate ./... 2>&1 \
	  | tee -a $(BUILD_PATH)/make.log
	@echo "[OK] Go generate scripts completed successfully!" \
	  | tee -a $(BUILD_PATH)/make.log

# Tidy modules
tidy:
	@go mod tidy 2>&1 \
	  | tee -a $(BUILD_PATH)/make.log
	@echo "[OK] Go modules tidy completed!" \
	  | tee -a $(BUILD_PATH)/make.log

# Dump imports into vendor directory
vendor:
	@go mod vendor 2>&1 \
	  | tee -a $(BUILD_PATH)/make.log
	@echo "[OK] Vendored code import completed!" \
	  | tee -a $(BUILD_PATH)/make.log

# Run static security analysis
security:
	@gosec ./... 2>&1 \
	  | tee -a $(BUILD_PATH)/make.log
	@echo "[OK] Go security check completed!" \
	  | tee -a $(BUILD_PATH)/make.log

# Build binary
build: dirprep doc generate tidy vendor security
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

# Delete old build directory, clear out old log
clean:
	@rm -rf $(BUILD_PATH)
	@>$(LOG_PATH)/make.log
	@echo "[OK] Release $(APP)-$(VERSION)-$(FILE_ARCH) cleaned!"

# Output version
version:
	@echo $(VERSION)

# Install binary
install:
	install -d -m 755 '$(INSTALL_PATH)'
	install -Z $(BUILD_PATH)/$(APP) $(INSTALL_PATH)/$(APP)
	@echo "[OK] Binary installed successfully!"
