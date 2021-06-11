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

.PHONY: clean-examples
clean-examples:
	@echo "$(BOLD)$(YELLOW)Cleaning up working directory.$(RESET)"
	@for folder in $$(ls examples); \
	do \
  		if [[ -d "examples/$${folder}" ]]; \
  		then \
  			rm -rf examples/$${folder}/.terraform; \
  			rm -rf examples/$${folder}/.terraform.lock.hcl; \
  			rm -rf examples/$${folder}/terraform.tfstate; \
  			rm -rf examples/$${folder}/terraform.tfstate.backup; \
  		fi; \
  	done
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
	@cd examples/virtual_database; terraform fmt; terraform init; terraform apply --auto-approve; cd -
	@echo "$(BOLD)$(YELLOW)Completed Virtual Database Creation.$(RESET)"
	@echo "$(BOLD)$(YELLOW)Create Folders.$(RESET)"
	@cd examples/folders; terraform fmt; terraform init; terraform apply --auto-approve; cd -
	@echo "$(BOLD)$(YELLOW)Completed Folders Creation.$(RESET)"
	@echo "$(BOLD)$(YELLOW)Create JDBC Data Source.$(RESET)"
	@cd examples/jdbc_data_source; terraform fmt; terraform init; terraform apply --auto-approve; cd -
	@echo "$(BOLD)$(YELLOW)Completed JDBC Data Source Creation.$(RESET)"
	@echo "$(BOLD)$(YELLOW)Create Base Views.$(RESET)"
	@cd examples/base_views; terraform fmt; terraform init; terraform apply --auto-approve; cd -
	@echo "$(BOLD)$(YELLOW)Completed Base Views Creation.$(RESET)"
	@echo "$(BOLD)$(YELLOW)Create Roles.$(RESET)"
	@cd examples/roles; terraform fmt; terraform init; terraform apply --auto-approve; cd -
	@echo "$(BOLD)$(YELLOW)Completed Roles Creation.$(RESET)"
	@echo "$(BOLD)$(YELLOW)Create Users.$(RESET)"
	@cd examples/users; terraform fmt; terraform init; terraform apply --auto-approve; cd -
	@echo "$(BOLD)$(YELLOW)Completed Users Creation.$(RESET)"

.PHONY: destroy-examples
destroy-examples:
	@echo "$(BOLD)$(YELLOW)Destroy Users.$(RESET)"
	@cd examples/users; terraform fmt; terraform init; terraform destroy --auto-approve; cd -
	@echo "$(BOLD)$(YELLOW)Completed Users Destruction.$(RESET)"
	@echo "$(BOLD)$(YELLOW)Destroy Roles.$(RESET)"
	@cd examples/roles; terraform fmt; terraform init; terraform destroy --auto-approve; cd -
	@echo "$(BOLD)$(YELLOW)Completed Roles Destruction.$(RESET)"
	@echo "$(BOLD)$(YELLOW)Destroy Base Views.$(RESET)"
	@cd examples/base_views; terraform fmt; terraform init; terraform destroy --auto-approve; cd -
	@echo "$(BOLD)$(YELLOW)Completed Base Views Destruction.$(RESET)"
	@echo "$(BOLD)$(YELLOW)Destroy JDBC Data Source.$(RESET)"
	@cd examples/jdbc_data_source; terraform fmt; terraform init; terraform destroy --auto-approve; cd -
	@echo "$(BOLD)$(YELLOW)Completed JDBC Data Source Destruction.$(RESET)"
	@echo "$(BOLD)$(YELLOW)Destroy Folders.$(RESET)"
	@cd examples/folders; terraform fmt; terraform init; terraform destroy --auto-approve; cd -
	@echo "$(BOLD)$(YELLOW)Completed Folders Destruction.$(RESET)"
	@echo "$(BOLD)$(YELLOW)Destroy Virtual Database.$(RESET)"
	@cd examples/virtual_database; terraform fmt; terraform init; terraform destroy --auto-approve; cd -
	@echo "$(BOLD)$(YELLOW)Completed Virtual Database Destruction.$(RESET)"

.PHONY: test
test:
	@cd denodo; go test; cd -
