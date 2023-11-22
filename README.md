### Как запустить проект

1. Необходимо установить go
2. Инициализировать проект:
    ```shell
    go mod init main.go
    ```
3. необходимо установить пакет gin
    ```shell
     go get github.com/gin-gonic/gin
    ```
4. Запустить проект выполнив следующую команду:
    ```shell
    go run main.go
    ```
5. Перейти на страницу http://0.0.0.0:8080/index

### Пересборка в режиме реального времени

   ```shell
   gin --appPort 8080 --port 80 --excludeDir frontend --all
   ```

### Сборка фронта

   ```shell
   cd frontend && npm run build
   ```

### Создание и запуск миграции

1. Перейти в папку migration-cli
2. Выполнить команду
   ```shell
   migrate create -ext sql -dir /migrations -seq MIGRATION_NAME 
   ```
3. Изменить файл миграции
4. Сбилдить программу-мигратор
   ```shell
   go build -o mcli
   ```
5. Переместить в нужную папку с БД и запустить
   ```shell
   ./mcli migrate up
   ```
