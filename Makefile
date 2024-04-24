include $(CURDIR)/.env
export
LOCAL_BIN:=$(CURDIR)/bin

# Local
.PHONY: build
build:  ## Сборка проекта
	@echo "  >  Building Pastebin binary..."
	go build -ldflags '-w -s' -a -o ./bin/pastebin ./cmd/pastebin

.PHONY: run
run:  ## Запуск проекта
	@echo "  >  Run Pastebin binary..."
	./bin/pastebin

.PHONY: build_run
build_run: build run  ## Сборка и запуск проекта

.PHONY: up-local
up-local:  ## Запуск необходимой инфраструктуры для локального запуска в Docker
	@echo "  >  Up project in Docker..."
	docker compose -f ./deployments/local/docker-compose.yml up -d

.PHONY: down-local
down-local:  ## Остановить инфраструктуру для локального запуска в Docker
	@echo "  >  Stop project in Docker..."
	docker compose -f ./deployments/local/docker-compose.yml down

# Docker
.PHONY: build_image
build_image:
	docker build -f ./build/Dockerfile -t pastebin .

.PHONY: up
up:  ## Запуск всего проекта в Docker
	@echo "  >  Up project in Docker..."
	docker compose -f ./deployments/docker/docker-compose.yml up -d

.PHONY: down
down:  ## Остановить проект в Docker
	@echo "  >  Stop project in Docker..."
	docker compose -f ./deployments/docker/docker-compose.yml down

# Other
.PHONY: clean
clean:
	@echo "  >  Cleaning project..."
	go clean
	rm -rf ./bin/**

.PHONY: tidy
tidy:  ## Форматирование и чистка несипользуемых модулей
	@echo "  >  Formatting and cleaning not used modules..."
	go fmt ./...
	go mod tidy -v

.PHONY: list
list:  ## Список пакетов
	@echo "  >  List of packages:"
	go list ./...

.PHONY: install-deps
install-deps: ## Установка необходимых инструментов
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3
	#GOBIN=$(LOCAL_BIN) go install github.com/rakyll/statik@v0.1.7
	sh ./scripts/install-goose.sh

# OpenApi
.PHONY: swagger-json-doc
swagger-json-doc:  ## Запуск генератора спецификации OpenAPI в формате JSON
	@echo "  >  Generating OpenAPI specification..."
	swagger-cli bundle api/openapi/spec/openapi.yaml --outfile api/openapi/openapi.json --type json

# OpenApi
.PHONY: swagger-yaml-doc
swagger-yaml-doc:  ## Запуск генератора спецификации OpenAPI в формате YAML
	@echo "  >  Generating OpenAPI specification..."
	swagger-cli bundle api/openapi/spec/openapi.yaml --outfile api/openapi/openapi.yaml --type yaml

# Linting
.PHONY: lint
lint:  ## Запуск линтера
	@echo "  >  Run Linter..."
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... --config .golangci.yml

.PHONY: lint-fix
lint-fix:  ## Запуск линтера с исправлением
	@echo "  >  Run Linter with fix..."
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... --fix --config .golangci.yml

# Migration
.PHONY: migration-status-local
migration-status-local:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} status -v

.PHONY: migration-up-local
migration-up-local:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR}  up -v

.PHONY: migration-down-local
migration-down-local:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR}  down -v

.PHONY: all
all: help

.PHONY: help
help:  Makefile ## Помощь
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
