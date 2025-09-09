# Subscription Aggregator API

RESTful сервис для агрегации данных об онлайн-подписках пользователей.

## 📋 Описание

Subscription Aggregator API предоставляет возможность управления подписками пользователей на онлайн-сервисы с возможностью аналитики и подсчета стоимости подписок за выбранные периоды.

## 🚀 Основные возможности

- ✅ CRUDL операции над подписками
- ✅ Подсчет суммарной стоимости подписок за период
- ✅ Фильтрация по пользователю и названию сервиса
- ✅ Swagger документация
- ✅ Docker контейнеризация
- ✅ PostgreSQL база данных с миграциями
- ✅ Логирование всех операций

## 🛠 Технологии

- **Go** (Gin framework)
- **PostgreSQL**
- **Docker & Docker Compose**
- **Swagger/OpenAPI**
- **GORM ORM**
- **Logrus** для логирования

## 📦 Структура проекта

```bash
subscription-aggregator-api/
├── cmd/
│ └── server/ # Точка входа приложения
├── internal/
│ ├── app/ # Инициализация приложения
│ ├── config/ # Конфигурация
│ ├── handlers/ # HTTP handlers
│ ├── models/ # Модели данных
│ ├── repository/ # Работа с базой данных
│ ├── service/ # Бизнес-логика
│ └── routes/ # Маршрутизация
├── migrations/ # SQL миграции
├── pkg/
│ └── logger/ # Логирование
├── docs/ # Swagger документация
├── .env # Конфигурационный файл
├── docker-compose.yml # Docker Compose конфигурация
└── Dockerfile # Docker конфигурация
```

## 🚀 Быстрый старт

### С использованием Docker (рекомендуется)

```bash
# Клонирование репозитория
git clone https://github.com/gingerfoxie/subscription-aggregator-api.git
cd subscription-aggregator-api

# Запуск сервиса
docker-compose up --build
```

### Локальный запуск

```bash
# Установка зависимостей
go mod tidy

# Генерация Swagger документации
swag init -g cmd/server/main.go

# Запуск приложения
go run cmd/server/main.go
```

## 📊 API Endpoints

### Подписки

|  Метод |      ENDPOINT    |          Описание            |
|--------|------------------|------------------------------|
| POST |`/api/v1/subscriptions`|Создать подписку|
| GET |`/api/v1/subscriptions`|Получить список подписок|
| GET |`/api/v1/subscriptions/{id}`|Получить подписку по ID|
| PUT |`/api/v1/subscriptions/{id}`|Обновить подписку|
| DELETE |`/api/v1/subscriptions/{id}`|Удалить подписку|


### Аналитика

|МЕТОД|ENDPOINT|Описание|
|--------|------------------|------------------------------|
|GET|`/api/v1/total`|Получить суммарную стоимость подписок|

## Swagger UI

Документация API доступна по адресу: http://localhost:8080/swagger/index.html

## ⚙️ Конфигурация

Создайте файл .env в корне проекта:

```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=subscription_db
DB_SSLMODE=disable
SERVER_PORT=8080
# Для Docker
DATABASE_URL=postgres://postgres:postgres@db:5432/subscription_db?sslmode=disable
# Logger
LOG_OUTPUT=file # значения: file|stdout
```

## 🗄 Миграции базы данных

Миграции автоматически применяются при запуске через Docker Compose. Для ручного управления миграциями используйте migrate tool.

## 🧪 Тестирование

```bash
# Запуск unit тестов
go test ./...

# Запуск с coverage
go test -cover ./...
# Запуск unit тестов
go test ./...

# Запуск с coverage
go test -cover ./...
```

## 📈 Мониторинг

Все операции логируются в формате JSON для удобной интеграции с системами мониторинга.

## 🤝 Вклад в проект

Форкните репозиторий  
Создайте ветку для вашей фичи (git checkout -b feature/AmazingFeature)  
Зафиксируйте изменения (git commit -m 'Add some AmazingFeature')  
Запушьте ветку (git push origin feature/AmazingFeature)  
Откройте Pull Request  

## 📄 Лицензия

Этот проект лицензирован под MIT License - смотрите файл LICENSE.md для подробностей.

## 👤 Автор

[gingerfoxie](https://github.com/gingerfoxie)

## 🙏 Благодарности

Gin Framework  
GORM  
Swagger  
