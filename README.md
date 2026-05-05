# Apprise Integration

Сервис с интеграцией **Apprise** для отправки уведомлений в Telegram, Discord, Slack, Email и другие каналы.

## О проекте

Этот проект использует [Apprise](https://github.com/caronc/apprise) для централизованной отправки уведомлений из приложения, фоновых задач, cron-процессов или CI/CD пайплайнов.

Apprise позволяет отправлять сообщения через единый интерфейс и общий URL-формат для большого числа notification-сервисов, включая Telegram, Discord, Slack, Amazon SNS, Gotify, email и другие.

## Запуск проекта

### 1. Требования

- Docker и Docker Compose
- Go 1.26+ (для локального запуска без Docker)

### 2. Настройка окружения

Создайте файл `.env` на основе `.env.example`:

```bash
cp .env.example .env
```

Отредактируйте переменные в `.env`:
- `APP_PORT`: Порт, на котором будет запущен сервис (по умолчанию 8080).
- `APPRISE_BASE_URL`: URL API Apprise (например, `http://localhost:8000`).
- `MAIL_SERVER_URL`: Формат URL для отправки email через Apprise (например, `mailto://user:pass@smtp.gmail.com`).

### 3. Запуск инфраструктуры

Для работы сервиса необходим запущенный контейнер **Apprise API**. Вы можете запустить его с помощью Docker:

```bash
docker run -d \
  --name=apprise \
  -p 8000:8000 \
  caronc/apprise
```

Или через `docker-compose.yaml`:

```bash
docker-compose up -d
```

### 4. Запуск сервиса

#### Локально:

```bash
go run cmd/apprise-mvp/main.go
```

#### Через Docker:

Если у вас настроен `Dockerfile`:

```bash
docker build -t apprise-mvp .
docker run -p 8080:8080 --env-file .env apprise-mvp
```

## API Документация

Для генерации или обновления документации Swagger используйте команду:

```bash
swag init -g cmd/apprise-mvp/main.go -ot json
```

После запуска сервиса документация Swagger доступна по адресу:
`http://localhost:8080/swagger`

## Тестирование (Email)

Для тестирования отправки Email уведомлений рекомендуется использовать [Mailpit](https://github.com/axllent/mailpit) — современный инструмент для перехвата почты.

Запуск Mailpit через Docker:

```bash
  mailpit:
    image: axllent/mailpit
    ports:
      - "1025:1025"
      - "8025:8025"
```

После запуска:
- SMTP сервер: `localhost:1025` (без аутентификации)
- Web-интерфейс: `http://localhost:8025`

Пример `MAIL_SERVER_URL` для `.env`:
`mailto://localhost:1025`

## Возможности

- Единая точка отправки уведомлений.
- Поддержка нескольких каналов одновременно.
- Настройка через переменные окружения.
- Использование как Python API.
- Возможность группировки каналов через конфиг и теги.
- Поддержка вложений в уведомлениях.
- Удобная интеграция в backend, cron, workers и CI/CD.

## Поддерживаемые каналы

Примеры сервисов, с которыми работает Apprise:

- Telegram
- Discord
- Slack
- Email (SMTP / mailto)
- Gotify
- ntfy
- Microsoft Teams
- Pushbullet
- Amazon SNS

Полный список сервисов и форматов URL можно посмотреть в официальной документации:
- [Apprise GitHub](https://github.com/caronc/apprise)
- [Apprise Docs](https://appriseit.com)
