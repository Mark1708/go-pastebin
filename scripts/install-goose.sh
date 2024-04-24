#!/bin/bash
if [ -f ./bin/goose ]; then
  echo "Goose already installed"
else
  git clone https://github.com/pressly/goose.git ./bin/goose-repo
  cd ./bin/goose-repo && go build -tags="no_mysql no_sqlite3 no_mssql no_redshift no_tidb no_clickhouse no_vertica no_ydb no_duckdb" -o ../goose ./cmd/goose
  cd ../ && rm -rf goose-repo/
fi