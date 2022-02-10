go build -o tmp/build ---> в какую директорию собрать go пакет, затем вызов как .sh

go run name.go ---> запуск как скрипт
docker exec golang-last tmp/build - выполнить команду скрипт на рабочем контейнере
docker exec -it golang-last /bin/bash - интерактивно зайти в контейнер
docker exec -it golang-last go build -o tmp/build
