#
# Makefile fragment for installing deps
#   Auto-detects between dep and govendor
#

GOTOOLS       = github.com/axw/gocov/gocov \
                github.com/AlekSi/gocov-xml \
                github.com/stretchr/testify/assert \
                github.com/robertkrimen/godocdown/godocdown \
                github.com/golangci/golangci-lint/cmd/golangci-lint

# Determine package dep manager
ifneq ("$(wildcard go.mod)", "")
	VENDOR     = "Go Modules"
	VENDOR_CMD = ${GO_CMD} mod tidy
	GO         = ${GO_CMD}
else ifneq ("$(wildcard Gopkg.toml)","")
	VENDOR     = dep
	VENDOR_CMD = ${VENDOR} ensure
	GOTOOLS    += github.com/golang/dep
	GO         = ${GO_CMD}
else ifneq ("$(wildcard Godeps/*)","")
	VENDOR     = godep
	VENDOR_CMD = echo "Not Implemented"
	GOTOOLS    += github.com/tools/godep
	GO         = godep go
else
	VENDOR     = govendor
	VENDOR_CMD = ${VENDOR} sync
	GOTOOLS    += github.com/kardianos/govendor
	GO         = ${VENDOR}
endif

tools: check-version
	@echo "=== $(PROJECT_NAME) === [ tools            ]: Installing tools required by the project..."
	@$(GO_CMD) install $(GOTOOLS)

tools-update: check-version
	@echo "=== $(PROJECT_NAME) === [ tools-update     ]: Updating tools required by the project..."
	@$(GO_CMD) get -u $(GOTOOLS)

deps: tools deps-only

deps-only:
	@echo "=== $(PROJECT_NAME) === [ deps             ]: Installing package dependencies required by the project..."
	@echo "=== $(PROJECT_NAME) === [ deps             ]: Detected '$(VENDOR)'"
	@$(VENDOR_CMD)

