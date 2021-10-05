.DEFAULT_GOAL := help
ROOT_DIR := $(dir $(realpath $(firstword $(MAKEFILE_LIST))))
SHELL := /usr/bin/env bash

IMAGE ?= kafkactl
VERSION ?= 0.0.0-development-SNAPSHOT

# used for docker-compose images
CONFLUENT_VERSION := 6.2.1

GOOS ?= linux
GOARCH ?= amd64
BIN_DIR ?= $(ROOT_DIR)/bin

##@ Helpers
.PHONY: print-%
print-%:  ## Print a var
	@echo $($*)

.PHONY: help
help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage: make \033[36m<target>\033[0m\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Main
.PHONY: build
build:  ## Build kafkactl docker image
	@docker build \
		--build-arg VERSION=$(VERSION) \
		--build-arg GOOS=$(GOOS) \
		--build-arg GOARCH=$(GOARCH) \
		-f docker/Dockerfile \
		-t $(IMAGE):$(VERSION) \
		.

.PHONY: cli
cli: build  ## Start a shell inside a kafkactl container
	@docker run --rm -it -v "$(HOME)/.kafkactl.yaml:/root/.kafkactl.yaml" $(IMAGE):$(VERSION) bash

.PHONY: genbin
genbin: build  ## Generate a kafkactl binary
	@docker run --rm -t -v "$(ROOT_DIR)/bin:/output_bin" $(IMAGE):$(VERSION) bash -c "cp /usr/local/bin/kafkactl /output_bin/kafkactl"

.PHONY: gendoc
gendoc: build  ## Generate cobra markup documentation
	@docker run --rm -t -v "$(ROOT_DIR)/docs:/output" $(IMAGE):$(VERSION) gendoc

##@ Development
DOCKER_COMPOSE := env CONFLUENT_VERSION=$(CONFLUENT_VERSION) docker-compose -f $(ROOT_DIR)/docker-compose.yaml
.PHONY: dev-cluster-start
dev-cluster-start:  ## Launch Kafka dev cluster
	@$(DOCKER_COMPOSE) up -d

.PHONY: dev-cluster-stop
dev-cluster-stop:  ## Stop Kafka dev cluster
	@$(DOCKER_COMPOSE) stop

.PHONY: dev-cluster-down
dev-cluster-down:  ## Destroy Kafka dev cluster
	@$(DOCKER_COMPOSE) down

.PHONY: dev-cluster-logs
dev-cluster-logs:  ## Tail logs of Kafka dev cluster
	@$(DOCKER_COMPOSE) logs -f kafka-rest
