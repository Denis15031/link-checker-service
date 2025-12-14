# Link Checker Service

HTTP-сервис для проверки доступности ссылок и генерации PDF-отчётов.

## Функции

- `/check` — принимает JSON `{"links": [...]}`, возвращает статусы ссылок (`available`/`not available`) и присваивает набору ID.
- `/report` — принимает `{"links_list": [...]}`, возвращает PDF-отчёт по указанным ID.

## Запуск

### Локально

```bash
go run cmd/server/main.go