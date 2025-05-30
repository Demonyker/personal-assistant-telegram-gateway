ifneq ($(wildcard .env),)
include .env
export
else
$(warning WARNING: .env file not found!)
endif

LOCAL_BIN:=$(CURDIR)/bin

deps: ### deps tidy + verify
	go mod tidy && go mod verify
.PHONY: deps

deps-audit: ### check dependencies vulnerabilities
	govulncheck ./...
.PHONY: deps-audit

format: ### Run code formatter
	gofumpt -l -w .
	gci write . --skip-generated -s standard -s default
.PHONY: format

linter-golangci: ### check by golangci linter
	golangci-lint run
.PHONY: linter-golangci

linter-hadolint: ### check by hadolint linter
	git ls-files --exclude='Dockerfile*' --ignored | xargs hadolint
.PHONY: linter-hadolint

bin-deps: ### install tools
	GOBIN=$(LOCAL_BIN) go install tool
.PHONY: bin-deps

pre-commit: format linter-golangci ### run pre-commit
.PHONY: pre-commit
