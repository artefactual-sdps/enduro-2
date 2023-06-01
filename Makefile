PROJECT := enduro-sdps
MAKEDIR := hack/make

include hack/make/bootstrap.mk
include hack/make/dep_golangci_lint.mk
include hack/make/dep_ent.mk
include hack/make/dep_goa.mk
include .bingo/Variables.mk

define NEWLINE


endef

IGNORED_PACKAGES := \
	github.com/artefactual-sdps/enduro/hack/genpkgs \
	github.com/artefactual-sdps/enduro/internal/api/auth/fake \
	github.com/artefactual-sdps/enduro/internal/api/design \
	github.com/artefactual-sdps/enduro/internal/api/gen/http/cli/enduro \
	github.com/artefactual-sdps/enduro/internal/api/gen/http/package_/client \
	github.com/artefactual-sdps/enduro/internal/api/gen/http/package_/server \
	github.com/artefactual-sdps/enduro/internal/api/gen/http/storage/client \
	github.com/artefactual-sdps/enduro/internal/api/gen/http/storage/server \
	github.com/artefactual-sdps/enduro/internal/api/gen/http/swagger/client \
	github.com/artefactual-sdps/enduro/internal/api/gen/http/swagger/server \
	github.com/artefactual-sdps/enduro/internal/api/gen/http/upload/client \
	github.com/artefactual-sdps/enduro/internal/api/gen/http/upload/server \
	github.com/artefactual-sdps/enduro/internal/api/gen/package_ \
	github.com/artefactual-sdps/enduro/internal/api/gen/package_/views \
	github.com/artefactual-sdps/enduro/internal/api/gen/storage \
	github.com/artefactual-sdps/enduro/internal/api/gen/storage/views \
	github.com/artefactual-sdps/enduro/internal/api/gen/swagger \
	github.com/artefactual-sdps/enduro/internal/api/gen/upload \
	github.com/artefactual-sdps/enduro/internal/batch/fake \
	github.com/artefactual-sdps/enduro/internal/package_/fake \
	github.com/artefactual-sdps/enduro/internal/storage/fake \
	github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db \
	github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db/enttest \
	github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db/hook \
	github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db/location \
	github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db/migrate \
	github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db/pkg \
	github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db/predicate \
	github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db/runtime \
	github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/schema \
	github.com/artefactual-sdps/enduro/internal/storage/persistence/fake \
	github.com/artefactual-sdps/enduro/internal/temporal/testutil \
	github.com/artefactual-sdps/enduro/internal/upload/fake \
	github.com/artefactual-sdps/enduro/internal/watcher/fake
PACKAGES		:= $(shell go list ./...)
TEST_PACKAGES	:= $(filter-out $(IGNORED_PACKAGES),$(PACKAGES))
TFORMAT := short

export PATH:=$(GOBIN):$(PATH)

.DEFAULT_GOAL := help

$(GOBIN)/bingo:
	$(GO) install github.com/bwplotka/bingo@latest

bingo: $(GOBIN)/bingo

env:  ## Print Go env variables.
	$(GO) env

tools: bingo  ## Install developer tools managed by bingo.
	bingo get
	bingo list

tparse:  ## Run all tests and outputs a coverage report using tparse.
	@$(GO) test -count=1 -json -cover  $(TEST_PACKAGES) | $(TPARSE) -follow -all -notests

test:  ## Run all tests and outputs a summary using gotestsum.
	@$(GOTESTSUM) --format=$(TFORMAT) $(TEST_PACKAGES)

test-race:  ## Run all tests with the race detector.
	@$(GOTESTSUM) --format=$(TFORMAT) $(TEST_PACKAGES) -- -race

ignored:  ## Print a list of packages ignored by tests.
	$(foreach PACKAGE,$(IGNORED_PACKAGES),@echo $(PACKAGE)$(NEWLINE))

lint: OUT_FORMAT ?= colored-line-number
lint: LINT_FLAGS ?= --timeout=5m --fix
lint: $(GOLANGCI_LINT)  ## Lint the project Go files with golangci-lint.
	golangci-lint run --out-format $(OUT_FORMAT) $(LINT_FLAGS)

gen-goa: $(GOA)  ## Generate Goa assets.
	goa gen github.com/artefactual-sdps/enduro/internal/api/design -o internal/api

gen-dashboard-client:  ## Generate Dashboard web client from OpenAPI spec.
	@rm -rf $(CURDIR)/dashboard/src/openapi-generator
	@docker container run --rm --user $(shell id -u):$(shell id -g) --volume $(CURDIR):/local openapitools/openapi-generator-cli:v6.1.0 \
		generate \
			--input-spec /local/internal/api/gen/http/openapi3.yaml \
			--generator-name typescript-fetch \
			--output /local/dashboard/src/openapi-generator/ \
			-p "generateAliasAsModel=false" \
			-p "withInterfaces=true" \
			-p "supportsES6=true"
	@echo "@@@@ Please, review all warnings generated by openapi-generator-cli above!"

gen-mock:  ## Generate mocks.
	$(MOCKGEN) -destination=./internal/package_/fake/mock_package_.go -package=fake github.com/artefactual-sdps/enduro/internal/package_ Service
	$(MOCKGEN) -destination=./internal/storage/fake/mock_storage.go -package=fake github.com/artefactual-sdps/enduro/internal/storage Service
	$(MOCKGEN) -destination=./internal/storage/persistence/fake/mock_persistence.go -package=fake github.com/artefactual-sdps/enduro/internal/storage/persistence Storage
	$(MOCKGEN) -destination=./internal/upload/fake/mock_upload.go -package=fake github.com/artefactual-sdps/enduro/internal/upload Service
	$(MOCKGEN) -destination=./internal/watcher/fake/mock_watcher.go -package=fake github.com/artefactual-sdps/enduro/internal/watcher Service
	$(MOCKGEN) -destination=./internal/api/auth/fake/mock_ticket_store.go -package=fake github.com/artefactual-sdps/enduro/internal/api/auth TicketStore

gen-ent: $(ENT)  ## Generate Ent assets.
	ent generate ./internal/storage/persistence/ent/schema \
		--feature sql/versioned-migration \
		--target=./internal/storage/persistence/ent/db

help:  ## Show this help.
	@grep -hE '^[A-Za-z0-9_ \-]*?:.*##.*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: *
