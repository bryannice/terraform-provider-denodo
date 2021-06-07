# -----------------------------------------------------------------------------
# Terraform Provider Denodo
# -----------------------------------------------------------------------------

# -----------------------------------------------------------------------------
# Internal Variables
# -----------------------------------------------------------------------------

#BOLD :=$(shell tput bold)
#GREEN :=$(shell tput setaf 2)
#RED :=$(shell tput setaf 1)
#RESET :=$(shell tput sgr0)
#YELLOW :=$(shell tput setaf 3)

# -----------------------------------------------------------------------------
# Git Variables
# -----------------------------------------------------------------------------

GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
GIT_REPOSITORY_NAME := $(shell git config --get remote.origin.url | cut -d'/' -f2 | cut -d'.' -f1)
GIT_ACCOUNT_NAME := $(shell git config --get remote.origin.url | cut -d'/' -f1 | cut -d':' -f2)
GIT_SHA := $(shell git log --pretty=format:'%H' -n 1)
GIT_TAG ?= $(shell git describe --always --tags | awk -F "-" '{print $$1}')
GIT_TAG_END ?= HEAD
GIT_VERSION := $(shell git describe --always --tags --long --dirty | sed -e 's/\-0//' -e 's/\-g.......//')
GIT_VERSION_LONG := $(shell git describe --always --tags --long --dirty)

# -----------------------------------------------------------------------------
# Terraform Provider Denodo Variables
# -----------------------------------------------------------------------------

BINARY=$(GIT_REPOSITORY_NAME)
HOSTNAME=custom.com
NAME=denodo
NAMESPACE=$(GIT_ACCOUNT_NAME)
VERSION=0.1
OS_ARCH := $(shell go env GOHOSTOS)_$(shell go env GOHOSTARCH)

# -----------------------------------------------------------------------------
# Terraform Provider Denodo Targets
# -----------------------------------------------------------------------------
.PHONY: clean-build
clean-build:
	@echo "$(BOLD)$(YELLOW)Cleaning up working directory.$(RESET)"
	@rm -rf ~/.terraform.d
	@echo "$(BOLD)$(GREEN)Completed cleaning up working directory.$(RESET)"

.PHONY: clean-data-example
clean-data-example:
	@echo "$(BOLD)$(YELLOW)Cleaning up working directory.$(RESET)"
	@rm -rf examples/data_source/.terraform
	@rm -rf examples/data_source/.terraform.lock.hcl
	@rm -rf examples/data_source/terraform.tfstate
	@rm -rf examples/data_source/terraform.tfstate.backup
	@echo "$(BOLD)$(GREEN)Completed cleaning up working directory.$(RESET)"

.PHONY: clean-resource-example
clean-resource-example:
	@echo "$(BOLD)$(YELLOW)Cleaning up working directory.$(RESET)"
	@rm -rf examples/resource/.terraform
	@rm -rf examples/resource/.terraform.lock.hcl
	@rm -rf examples/resource/terraform.tfstate
	@rm -rf examples/resource/terraform.tfstate.backup
	@echo "$(BOLD)$(GREEN)Completed cleaning up working directory.$(RESET)"

.PHONY: fmt
fmt:
	@go fmt ./...

.PHONY: build
build: fmt
	go build -o ${GIT_REPOSITORY_NAME}

.PHONY: install
install: clean-build build
	@mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	@mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}/${BINARY}_v${VERSION}

.PHONY: test-data-example
test-data-example: clean-data-example
	@cd examples/data_source; terraform fmt; terraform init; terraform apply --auto-approve; cd -

.PHONY: test-resource-example
test-resource-example: clean-resource-example
	@cd examples/resource; terraform fmt; terraform init; terraform apply --auto-approve; cd -

.PHONY: test
test:
	@cd denodo; go test; cd -