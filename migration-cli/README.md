# [migration-cli](https://dev.to/lucasnevespereira/sql-migrations-with-cobra-golang-migrate-3f75)

### Создание и запуск миграции

1. Создать миграцию
   ```shell
   migrate create -ext sql -dir migrations -seq MIGRATION_NAME 
   ```
2. Изменить файл миграции
3. Сбилдить программу-мигратор
   ```shell
   go build -o mcli
   ```
4. Переместить в нужную папку с БД и запустить
   ```shell
   ./mcli migrate up
   ```
5. Откатить миграцию
   ```shell
   ./mcli migrate down
   ```