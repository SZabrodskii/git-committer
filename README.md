# Git Committer CLI

Этот сервис предназначен для автоматического создания коммитов в Git-репозитории на основе параметров, указанных в конфигурационном файле.

## Установка

1. Склонируйте репозиторий:

    ```bash
    git clone https://github.com/user/repo.git
    ```

2. Скопируйте пример конфига:

    ```bash
    cp config/config.example.json config/config.json
    ```

3. Настройте конфигурационный файл `config/config.json` под свои нужды.

## Пример конфигурации

Пример файла конфигурации (`config/config.json`):

```json
{
  "min_commits": 1,               // Минимальное количество коммитов в день
  "max_commits": 5,               // Максимальное количество коммитов в день
  "days": 30,                     // Количество дней, в течение которых будут создаваться коммиты
  "include_weekends": false,       // Включать ли выходные дни для создания коммитов (true/false)
  "weekend_commits": {             // Настройки коммитов для выходных (актуально, если include_weekends = true)
    "min_commits": 1,             // Минимальное количество коммитов в выходные
    "max_commits": 3              // Максимальное количество коммитов в выходные
  },
  "repo_url": "https://github.com/user/repo.git",  // URL репозитория, в который будут отправляться коммиты
  "commit_template": "feat: some commit message"    // Шаблон для сообщений коммитов (например, в формате Conventional Commits)
}

