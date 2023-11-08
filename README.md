### Как запустить проект

1. Необходимо установить go
2. необходимо установить пакет gin
    ```shell
     go get github.com/gin-gonic/gin
    ```
3. Запустить проект выполнив следующую команду:
    ```shell
    go run main.go
    ```
4. Перейти на страницу http://0.0.0.0:8080/index

### Пересборка в режиме реального времени

   ```shell
   gin --appPort 8080 --port 80 --excludeDir frontend --all
   ```

### Сборка фронта

   ```shell
   cd frontend && npm run build
   ```
