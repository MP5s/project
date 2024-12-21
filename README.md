# Calc Service

Это веб-сервис для вычисления арифметических выражений.

## Как использовать

Сервис принимает POST-запросы на `/api/v1/calculate` с телом запроса в формате JSON:

```json
{
    "expression": "выражение, которое ввёл пользователь"
}
Примеры использования
Успешный запрос:
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2 + 2"
}'
Ответ:

{
    "result": "4.000000"
}
Ошибка 422 (недопустимое выражение):
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2 + a"
}'
Ответ:

{
    "error": "Expression is not valid"
}
Ошибка 500 (внутренняя ошибка сервера):
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2 / 0"
}'
Ответ:

{
    "error": "Internal server error"
}