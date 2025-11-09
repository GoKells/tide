# --- Project settings ---
BINARY = tide
MAIN   = cmd/tide/main.go


# --- Detect OS using Go ---
DETECTED_OS := $(shell go env GOOS)
EXE :=
SEP := /

ifeq ($(DETECTED_OS),windows)
    EXE = .exe
    SEP = \\
endif

# --- Create build directory (works everywhere) ---
define make_dir
	go run -exec "" "fmt" >/dev/null 2>&1 || true
	if not exist build$(SEP)$(1) mkdir build$(SEP)$(1) 2>nul || mkdir -p build/$(1)
endef

# --- Default build (current OS) ---
build:
	@echo Building for $(DETECTED_OS)...
	$(call make_dir,$(DETECTED_OS))
	go build -o build$(SEP)$(DETECTED_OS)$(SEP)$(BINARY)$(EXE) $(MAIN)
	@echo ✅ Build complete: build$(SEP)$(DETECTED_OS)$(SEP)$(BINARY)$(EXE)

# --- Cross compilation targets ---
build-linux:
	@echo Building for Linux...
	go env -w GOOS=linux GOARCH=amd64
	go build -o build/linux/$(BINARY) $(MAIN)
	@echo ✅ Build complete: build/linux/$(BINARY)

build-windows:
	@echo Building for Windows...
	go env -w GOOS=windows GOARCH=amd64
	go build -o build/windows/$(BINARY).exe $(MAIN)
	@echo ✅ Build complete: build/windows/$(BINARY).exe

build-macos:
	@echo Building for macOS...
	go env -w GOOS=darwin GOARCH=arm64
	go build -o build/macos/$(BINARY) $(MAIN)
	@echo ✅ Build complete: build/macos/$(BINARY)

# --- Cleanup ---
clean:
	@echo Cleaning build directory...
	@if exist build (rmdir /s /q build) 2>nul || rm -rf build
	@echo 🧹 Clean complete.
