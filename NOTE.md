## Пояснительная записка к тестовому заданию "Платежная система"

### Описание
В соотвествие с требованиями реализован бекенд приложения.

Все требования по заданию выполнены. 

Приложение предоставляет HTTP API:
	добавление, получение кошелька, депозит, перевод с кошелька, отчет по операциям, версия сервиса.

Отрицательный баланс контролируется на уровне ограничений БД и логики в обработчике перевода денежных средств. 

Дублирование операций контролируется по совпадению всех полей запроса deposit, transfer и времени последней такой операции (в коде захаркожено 2 минуты).
Валюта в домене не используется и предполагалась как единственная.

В качестве хранилища используется СУБД Postgresql. 

Миграции на БД накатываются автоматически при старте сервиса.

При запуске сервиса порт по-умолчанию 8000.

### Описание API

С помощью go-swag сгенерирован сваггер файл, который находится в [docs](./docs/swagger.json)

[Примеры вызов API](README.md) 

### Запуск приложения

#### Требования:

* linux
* golang 1.12 или выше
* docker
* docker-compose
* make

#### Запуск
```shell script
docker network create payment_system_network
make build run
```


#### Запуск тестов
```shell script
docker network create payment_system_network
make run-local
make test
```

### О тестах:
Для экономии времени реализовано всего пара модульных тестов. Покрытие около процента. 

При запуске интеграционных тестов должны быть подняты зависимости: Postgresql. 
Тесты выполняются с помощью go test (см запуск тестов). 
Если порт занят, то необходимо исправить конфигурацию в tests/payment.yaml

Реализованы три интеграционных теста: добавление и получение кошелька. Депозит на кошелек.

### Качество кода:
* Местами отсутствуют комментарии
* Линтер выдает предупреждения, где-то обоснованно, где-то ложные срабатывания, нужно бы подавить
