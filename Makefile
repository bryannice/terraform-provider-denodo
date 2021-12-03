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
# Docker Variables
# -----------------------------------------------------------------------------

STEP_1_IMAGE ?= alpine:3.13
DOCKER_IMAGE_TAG ?= $(GIT_VERSION)
DOCKER_IMAGE_NAME := $(GIT_REPOSITORY_NAME)

# -----------------------------------------------------------------------------
# Terraform Provider Denodo Variables
# -----------------------------------------------------------------------------

BINARY=$(GIT_REPOSITORY_NAME)
HOSTNAME=custom.com
NAME=denodo
NAMESPACE=$(GIT_ACCOUNT_NAME)
VERSION=0.1
OS_ARCH := linux_amd64

# -----------------------------------------------------------------------------
# Docker Targets
# -----------------------------------------------------------------------------

.PHONY: docker-image
docker-image: docker-rmi-for-image
	@echo "$(BOLD)$(YELLOW)Building docker image.$(RESET)"
	@docker build \
		--build-arg STEP_1_IMAGE=$(STEP_1_IMAGE) \
		--tag $(DOCKER_IMAGE_NAME):latest \
		--tag $(DOCKER_IMAGE_NAME):$(GIT_VERSION) \
		--file build/package/Dockerfile \
		.
	@echo "$(BOLD)$(GREEN)Completed building docker image.$(RESET)"

.PHONY: docker-rmi-for-image
docker-rmi-for-image:
	-docker rmi --force \
		$(DOCKER_IMAGE_NAME):$(GIT_VERSION) \
		$(DOCKER_IMAGE_NAME):latest

.PHONY: dev-env-up
dev-env-up:
	@echo "$(BOLD)$(YELLOW)Create development environment$(RESET)"
	@docker-compose -f deployments/docker-compose.yml up -d
	@echo "$(BOLD)$(GREEN)Completed creating development environment.$(RESET)"

.PHONY: dev-env-down
dev-env-down:
	@echo "$(BOLD)$(YELLOW)Create development environment$(RESET)"
	@docker-compose -f deployments/docker-compose.yml down
	@echo "$(BOLD)$(GREEN)Completed creating development environment.$(RESET)"

# -----------------------------------------------------------------------------
# Terraform Provider Denodo Targets
# -----------------------------------------------------------------------------
.PHONY: clean-build
clean-build:
	@echo "$(BOLD)$(YELLOW)Cleaning up working directory.$(RESET)"
	@rm -rf ~/.terraform.d
	@echo "$(BOLD)$(GREEN)Completed cleaning up working directory.$(RESET)"

.PHONY: clean-examples
clean-examples:
	@echo "$(BOLD)$(YELLOW)Cleaning up working directory.$(RESET)"
	@rm -rf tests/.terraform
	@rm -rf tests/.terraform.lock.hcl
	@rm -rf tests/terraform.tfstate
	@rm -rf tests/terraform.tfstate.backup
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

.PHONY: test-examples
test-examples: clean-examples
	@echo "$(BOLD)$(YELLOW)Create Virtual Database.$(RESET)"
	@cd tests; terraform fmt; terraform init; terraform apply --auto-approve; cd -
	@echo "$(BOLD)$(YELLOW)Completed Virtual Database Creation.$(RESET)"

.PHONY: destroy-examples
destroy-examples:
	@echo "$(BOLD)$(YELLOW)Destroy Virtual Database.$(RESET)"
	@cd tests; terraform fmt; terraform init; terraform destroy --auto-approve; cd -
	@echo "$(BOLD)$(YELLOW)Completed Virtual Database Destruction.$(RESET)"
.PHONY: test
test:
	@cd denodo; go test; cd -
