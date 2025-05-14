<div align="center">
  <a href="https://git.io/typing-svg"><img src="https://readme-typing-svg.herokuapp.com?font=Tektur&pause=1000&color=DC00F7&center=true&width=435&lines=TELEGRAM+REMINDER+BOT" alt="Typing SVG" /></a>
</div>

# О проекте

Telegram Reminder Bot — бот для создания напоминаний в чатах. Он реализован полностью на Golang и использует многопоточность для работы с несколькими задачами одновременно.

# Функционал бота

- **Командный функционал**: По команде `@имя_бота ctrl NM`, где N - интервал, M - единица времени, бот ставит напоминание для данного пользователя и отправляет его в чат по истечении времени. Удобно для пользователей веб-версии.

- **Графический функционал**: По нажатию на кнопку «Добавить напоминание» бот последовательно собирает необходимую информацию (текст напоминания, интервал, единица времени). Удобно для пользователей мобильной версии.

# Технологии

- Golang
- PostgreSQL
- REST API
- Docker
- TgBotApi

# Установка бота

1. **Клонирование репозитория**:
   ```sh
   git clone https://github.com/WB-tg-bot/telegram_reminder_bot.git

# tg-bot

2. **Установка зависимостей в директории tg-bot**:
   ```sh
   go mod download
   go mod tidy

3. **Проверка и изменение файлов конфигурации**:
  * Файл .env:
     ```sh
    TELEGRAM_BOT_TOKEN=your_Tg-bot_TOKEN

  * Файл configs/config.yml:
     ```sh
    url_create_task: "http://telegram-reminder-bot:8000/create-task"
    url_get_task: "http://telegram-reminder-bot:8000/tasks"

# Server

4. **Установка зависимостей в директории storage**:
   ```sh
   go mod download
   go mod tidy

5. **Проверка и изменение файлов конфигурации**:
  * Файл .env:
     ```sh
     TELEGRAM_BOT_TOKEN=your_Tg-bot_TOKEN
     DB_PASSWORD=your_DB_PASS

  * Файл configs/config.yml:
     ```sh
     port: "8000"

    db:
      username: "postgres"
      host: "db"
      port: "5432"
      dbname: "postgres"
      sslmode: "disable"
      timezone: "Europe/Moscow"

6. **Для Windows 10**:
  * Установка Scoop:
     ```sh
      irm get.scoop.sh -outfile 'install.ps1'
      .\install.ps1 -RunAsAdmin

  * Установка Migrate:
     ```sh
     scoop install migrate

  * Установка Make:
     ```sh
     @powershell -NoProfile -ExecutionPolicy unrestricted -Command "iex ((new-object net.webclient).DownloadString('https://chocolatey.org/install.ps1'))" && SET PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin
     choco install make


7. **Разворачиваем приложение в Docker**:
   ```sh
   make run

# Настрока и запуск проекта локально

1. **Проверка и изменение файлов конфигурации tg-bot**:

  * Файл configs/config.yml:
     ```sh
     url_create_task: "http://localhost:8000/create-task"
     url_get_task: "http://localhost:8000/tasks"

2. **Проверка и изменение файлов конфигурации сервера**:
  * Файл configs/config.yml:
    ```sh
    host: "localhost"
    port: "5436"

4. **Инициализация и миграция базы данных**:
   ```sh
   make db_init
   make db_migrate

5. **Запуск сервера**:
   ```sh
   go run cmd/main.go

6. **Запуск Telegram-бота**:
   ```sh
   cd ../tg-bot
   go run cmd/main.go

Приложение готово к использованию

















