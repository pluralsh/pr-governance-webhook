ROOT_DIRECTORY := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

PROJECT_NAME := pr-governance-webhook

IMAGE_REGISTRIES := ghcr.io
IMAGE_REPOSITORY := plural

IMG ?= pr-governance-webhook:latest

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)


ifndef GOPATH
$(error $$GOPATH environment variable not set)
endif


##@ General

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Run

.PHONY: run
run: build ## run webhook
	go run cmd/*.go

##@ Build

.PHONY: build
build: ## build webhook
	go build -o bin/deployment-agent cmd/*.go

docker-build: ## build image
	docker build -t ${IMG} .

docker-push: ## push image
	docker push ${IMG}

##@ Tests

.PHONY: lint
lint: $(PRE) ## run linters
	golangci-lint run ./...

.PHONY: fix
fix: $(PRE) ## fix issues found by linters
	golangci-lint run --fix ./...

release-vsn: # tags and pushes a new release
	@read -p "Version: " tag; \
	git checkout main; \
	git pull --rebase; \
	git tag -a $$tag -m "new release"; \
	git push origin $$tag

delete-tag:  ## deletes a tag from git locally and upstream
	@read -p "Version: " tag; \
	git tag -d $$tag; \
	git push origin :$$tag

