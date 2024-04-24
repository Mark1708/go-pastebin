# Миграции

## Установка Goose
```shell
go install github.com/pressly/goose/v3/cmd/goose@latest
# или 
git clone https://github.com/pressly/goose.git
cd goose
go build -tags="no_mysql no_sqlite3 no_mssql no_redshift no_tidb no_clickhouse no_vertica no_ydb no_duckdb" -o goose ./cmd/goose
```
## Правила создания миграций
Все миграции расположенны в директории `migrations`.

### Создание 
Создался новый файл с именем 20230416135213_new_user_table.sql, 
где  20230416135213 — это временная метка (год, месяц, число, час, минуты и секунды)

Автоматически в этом файле указывается две области: для написания миграции `-- +goose Up`, 
которая применяется при накатывании, и миграции `-- +goose Down`, которая выполняется при откатывании миграции.
```shell
./bin/goose -dir migrations create new_user_table sql
```

### Накатываем миграцию
```shell
./bin/goose -dir migrations postgres "postgresql://user:pass@localhost:5432/db?sslmode=disable" up
# или
export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=postgresql://user:pass@localhost:5432/db?sslmode=disable
./bin/goose -dir migrations up
```

### Откатываем миграцию
```shell
./bin/goose -dir migrations down-to <VERSION>
```