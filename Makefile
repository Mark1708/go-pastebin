.PHONY: build
build: clean  ## Сборка проекта
	@echo "  >  Building Pastebin binary..."
	go build -ldflags '-w -s' -a -o ./bin/pastebin ./cmd/pastebin

.PHONY: build_image
build_image:
	docker build -f ./build/Dockerfile -t pastebin .

.PHONY: run
run:  ## Запуск проекта
	echo "  >  Run Pastebin binary..."
	./bin/pastebin

.PHONY: build_run
build_run: build run  ## Сборка и запуск проекта

.PHONY: clean
clean:
	@echo "  >  Cleaning project..."
	go clean
	./bin/pastebin

.PHONY: tidy
tidy:  ## Форматирование и чистка несипользуемых модулей
	@echo "  >  Formatting and cleaning not used modules..."
	go fmt ./...
	go mod tidy -v

.PHONY: list
list:  ## Список пакетов
	@echo "  >  List of packages:"
	go list ./...

.PHONY: lint
lint:  ## Запуск линтера
	@echo "  >  Run Linter..."
	golangci-lint run ./...

.PHONY: swagger-doc
swagger-doc:  ## Запуск генератора спецификации OpenAPI
	@echo "  >  Generating OpenAPI specification..."
	swag init -g cmd/pastebin/main.go -o ./api/openapi-spec -ot yaml

.PHONY: all
all: help

.PHONY: help
help:  Makefile ## Помощь
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
