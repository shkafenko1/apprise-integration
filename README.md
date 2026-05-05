# Apprise Integration

Сервис с интеграцией **Apprise** для отправки уведомлений в Telegram, Discord, Slack, Email и другие каналы.

## О проекте

Этот проект использует [Apprise](https://github.com/caronc/apprise) для централизованной отправки уведомлений из приложения, фоновых задач, cron-процессов или CI/CD пайплайнов.

Apprise позволяет отправлять сообщения через единый интерфейс и общий URL-формат для большого числа notification-сервисов, включая Telegram, Discord, Slack, Amazon SNS, Gotify, email и другие. Также библиотека поддерживает CLI, конфигурационные файлы, теги и вложения для тех сервисов, где это доступно.

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
