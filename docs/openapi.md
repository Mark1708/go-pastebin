# OpenAPI v3

> Спецификация хранится в директории `api/openapi-spec` 

## Генератор спецификации для API
### Установка
```shell
sudo npm install -g swagger-cli
```

### Генерация
```shell
swagger-cli bundle api/openapi-spec/openapi.yaml --outfile api/openapi-spec/build/openapi.yaml --type yaml
swagger-cli bundle api/openapi-spec/openapi.yaml --outfile api/openapi-spec/build/openapi.json --type json
```

## Полезные источники
- https://habr.com/ru/articles/776538/
- https://davidgarcia.dev/posts/how-to-split-open-api-spec-into-multiple-files/