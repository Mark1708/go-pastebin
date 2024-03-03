PROJECTNAME=$(shell basename "$(PWD)")

# Сборка проекта
build:
	@echo "  >  Building Pastebin binary..."
	go build -ldflags '-w -s' -a -o ./bin/pastebin ./cmd/pastebin

# Запуск проекта
run:
	@echo "  >  Run Pastebin binary..."
	./bin/pastebin

# Сборка и запуск проекта
all: build run

# Список пакетов
list:
	@echo "  >  List of packages:"
	go list ./...

# Список пакетов
list:
	@echo "  >  List of packages:"
	go list ./...

# Запск линтера
lint:
	@echo "  >  Run Linter..."
	golangci-lint run ./...
