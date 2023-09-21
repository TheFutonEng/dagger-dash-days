DPT_BIN := ./build/dpt

.PHONY: help
help: ## Display this help information
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
	  | sort | awk 'BEGIN {FS = ":.*?## "}; \
	  {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

clean: ## Clean the build directory
	rm -rf build

build-cli-linux-amd: ## Build the dpt CLI for Linux on AMD64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="$(BUILD_ARGS)" -o build/dpt .

build-cli-linux-arm: ## Build the dpt CLI for Linux on ARM
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="$(BUILD_ARGS)" -o build/dpt-arm .

build-cli-mac-intel: ## Build the dpt CLI for macOS on AMD64
	GOOS=darwin GOARCH=amd64 go build -ldflags="$(BUILD_ARGS)" -o build/dpt-mac-intel .

build-cli-mac-apple: ## Build the dpt CLI for macOS on ARM
	GOOS=darwin GOARCH=arm64 go build -ldflags="$(BUILD_ARGS)" -o build/dpt-mac-apple .

build-cli-windows-amd: ## Build the dpt CLI for Windows on AMD64
	GOOS=windows GOARCH=amd64 go build -ldflags="$(BUILD_ARGS)" -o build/dpt.exe . ## Build the dpt CLI for Windows on AMD64

build-cli-windows-arm: ## Build the dpt CLI for Windows on ARM
	GOOS=windows GOARCH=arm64 go build -ldflags="$(BUILD_ARGS)" -o build/dpt-arm.exe . ## Build the dpt CLI for Windows on ARM

build-cli-linux: build-cli-linux-amd build-cli-linux-arm ## Build the dpt CLI for Linux on AMD64 and ARM

build-cli: build-cli-linux-amd build-cli-linux-arm build-cli-mac-intel build-cli-mac-apple build-cli-windows-amd build-cli-windows-arm ## Build the CLI
