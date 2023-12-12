build:
	cd ../frontend && npm run build
	go build
	./backend

init-dependency:
	go get -u github.com/antonfisher/nested-logrus-formatter
	go get -u github.com/gin-gonic/gin
	go get -u golang.org/x/crypto
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/sqlite
	go get -u github.com/sirupsen/logrus
	go get -u github.com/joho/godotenv
	go get -u github.com/google/wire

wire:
	wire gen Kogalym/backend/config