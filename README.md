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
  "min_commits": 1,
  "max_commits": 5,
  "days": 30,
  "include_weekends": false,
  "weekend_commits": {
    "min_commits": 1,
    "max_commits": 3
  },
  "repo_url": "https://github.com/user/repo.git",
  "commit_template": "feat: some commit message"
}

