# Тестовое задание БВС
Был разработан веб-сервис согласно тестовому заданию, также дополнительно был
собран Docker контейнер. Для быстрый и удобной проверки также дополниельно был
собран Docker-compose.
## О приложении
В веб-сервере использовался fiber и go-redis пакеты. Приложение построено с использованием
чистой архитектуры (handler, usecase, repository), что позволит легко расширять и модифицировать его
## Запуск сервера и бэнчмарка
Используется контейнер Redis, собирается контейнер с веб приложением, он же заполняет
тестовые данные, и затем запускается apache-becnmark (ab) с 20 конкурентными
запросами. Наружу экспортируется порт веб-сервера 8010, согласно заданию.
```shell
docker-compose up
```