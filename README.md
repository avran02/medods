# MEDODS TEST TASK

### Запуск

``` bash
# клинтруем репозиторий
git clone https://github.com/avran02/medods
# заходим в папку с проектом
cd medods
# получаем шаблон .env файла
cp example.env .env
# Поднимаем всё в фоне со сборкой образа приложения
docker-compose up -d --build
```
___
### Доступные эндпоинты

- localhost:8080 - UI к БДшке
- __localhost:3000/api/v1/get-tokens__ - получение токенов через _GET_ запрос 
- __localhost:3000/api/v1/refresh-tokens__ - Обновление токенов оп refresh токену через _POST_ запрос. Данные передаются в JSON 
