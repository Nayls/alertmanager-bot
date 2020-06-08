OK_COLOR=\033[32;01m
NO_COLOR=\033[0m

SHELL=/bin/bash
DOCKER?=$(shell grep alias\ docker= ~/.bashrc | awk -F"'" '{print $$2}')
DOCKER_COMPOSE?=$(shell grep alias\ docker-compose= ~/.bashrc | awk -F"'" '{print $$2}')

# ifeq ($(CI_SERVER),)
# 	include .env
# 	export $(shell sed 's/=.*//' .env)
# endif

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' Makefile | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

DOCKER_BUILDKIT=1
CI_REGISTRY_IMAGE?=nayls/alertmanager-bot
CI_COMMIT_REF_SLUG?=$(shell git symbolic-ref --short -q HEAD | sed 's/\//-/')
CI_PIPELINE_URL?=local-build
CI_COMMIT_SHA?=$(shell git rev-parse -q HEAD)
CI_COMMIT_SHORT_SHA?=$(shell git rev-parse --short=8 -q HEAD)

COMMAND?=/bin/bash
DOCKERFILE?=${PWD}/Dockerfile
DOCKER_CONTEXT?=${PWD}
IMAGE_NAME?=${CI_REGISTRY_IMAGE}:${CI_COMMIT_REF_SLUG}

CONTAINER_NAME?=alertmanager-bot
APP_URL?=http
APP_SCHEME?=localhost
PUBLIC_PORT?=8080
INTERNAL_PORT?=8080

.DEFAULT_GOAL := help

all: download build run

.PHONY: force-build ## Build with vendor and force rebuild
force-build: download vendor build-force run

.PHONY: run
run: ## Run alertmanager-bot
	@printf "$(OK_COLOR)==>$(NO_COLOR) Run alertmanager-bot\n"
	@./bin/alertmanager-bot

.PHONY: download
download: ## Download packages
	@printf "$(OK_COLOR)==>$(NO_COLOR) Download packages\n"
	@go mod download

.PHONY: vendor
vendor: ## Vendoring packages
	@printf "$(OK_COLOR)==>$(NO_COLOR) Vendoring packages\n"
	@go mod vendor

.PHONY: build
build: ## Build alertmanager-bot
	@printf "$(OK_COLOR)==>$(NO_COLOR) Build alertmanager-bot\n"
	@CGO_ENABLED=0 GOOS=linux \
	go build \
	-ldflags " -X 'main.CI_COMMIT_REF_SLUG=${CI_COMMIT_REF_SLUG}' \
		-X 'main.CI_PIPELINE_URL=${CI_PIPELINE_URL}'  \
		-X 'main.CI_COMMIT_TITLE=${CI_COMMIT_TITLE}' \
		-X 'main.CI_COMMIT_SHA=${CI_COMMIT_SHA}' \
		-X 'main.GIT_USER_NAME=${GIT_USER_NAME}' \
		-X 'main.GIT_USER_EMAIL=${GIT_USER_EMAIL}' \
		-X 'main.BuildDate=$(date)'" \
	-installsuffix cgo -o ./bin/alertmanager-bot ./main.go

.PHONY: build-force
build-force: ## Force build alertmanager-bot
	@printf "$(OK_COLOR)==>$(NO_COLOR) Build alertmanager-bot\n"
	@CGO_ENABLED=0 GOOS=linux \
	go build \
	-mod vendor \
	-a \
	-ldflags " -X 'main.CI_COMMIT_REF_SLUG=${CI_COMMIT_REF_SLUG}' \
		-X 'main.CI_PIPELINE_URL=${CI_PIPELINE_URL}'  \
		-X 'main.CI_COMMIT_TITLE=${CI_COMMIT_TITLE}' \
		-X 'main.CI_COMMIT_SHA=${CI_COMMIT_SHA}' \
		-X 'main.GIT_USER_NAME=${GIT_USER_NAME}' \
		-X 'main.GIT_USER_EMAIL=${GIT_USER_EMAIL}' \
		-X 'main.BuildDate=$(date)'" \
	-installsuffix cgo -o ./bin/alertmanager-bot ./main.go

.PHONY: docker-build
docker-build: ## Build docker image
	@printf "$(OK_COLOR)==>$(NO_COLOR) Build docker image\n"
	@${DOCKER} build --rm --compress --pull --progress plain \
		--build-arg BUILDKIT_INLINE_CACHE=1 \
		--tag ${IMAGE_NAME} \
		--target runtime-image \
		--file Dockerfile \
		./

.PHONY: docker-compose-up
docker-compose-up: ## Run service with docker-compose
	@printf "$(OK_COLOR)==>$(NO_COLOR) Build docker image\n"
	@${DOCKER_COMPOSE} up -d

.PHONY: docker-compose-down
docker-compose-down: ## Stop service with docker-compose
	@printf "$(OK_COLOR)==>$(NO_COLOR) Build docker image\n"
	@${DOCKER_COMPOSE} down


.PHONY: develop-prometheus-up
develop-prometheus-up: ## Develop: run prometus with docker-compose
	@printf "$(OK_COLOR)==>$(NO_COLOR) Start prometheus stack\n"
	@${DOCKER_COMPOSE} -f ./ops/develop/prometheus/docker-compose.yaml up -d --force-recreate

.PHONY: develop-prometheus-down
develop-prometheus-down: ## Develop: stop prometus with docker-compose
	@printf "$(OK_COLOR)==>$(NO_COLOR) Stop prometheus stack\n"
	@${DOCKER_COMPOSE} -f ./ops/develop/prometheus/docker-compose.yaml down
