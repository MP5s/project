# Calc Service

Это веб-сервис для вычисления арифметических выражений.

## Инструкция по установке

1. Клонируйте проект.

> git clone https://github.com/MP5s/project.git

2. ```bash
cd project
```

3. Запуск сервера

```bash
go run ./cmd/calc_service/main.go
```
4. Запуск тестов

```bash
go test ./...
```

## Как использовать

Сервис принимает POST-запросы на `/api/v1/calculate` с телом запроса в формате JSON:

```json
{
    "expression": "выражение, которое ввёл пользователь"
}
```
Примеры использования
Успешный запрос:
### 1. Простое выражение
> curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d "{\"expression\": \"3 + 5\"}"

Ответ:
```json
{
    "result": "8.000000"
}
```
### 2. Выражение по-сложнее:
> curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d "{\"expression\": \"10 * 2 + 3 / 1 - 5\"}"

Ответ:
```json
{
    "result": "18.000000"
}
```
### 3. Запрос с недопустимым методом:
> curl -X GET http://localhost:8080/api/v1/calculate

Ответ:
```json
{
    "error": "Method is not allowed"
}
```
### 4. Ошибка 422 (недопустимое выражение):

> curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d "{\"expression\": \"3 + a\"}"

Ответ:
```json
{
    "error": "Expression is not valid"
}
```
### 5. Ошибка 400:
> curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json"

Ответ:
```json
{
    "error": "Invalid request body"
}
```
~_Сам проверял через POSTMAN_~
