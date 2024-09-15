# backend/Dockerfile
FROM golang:1.22-alpine

WORKDIR /go/src/app

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Компилируем приложение в исполняемый файл 'main'
RUN go build -o main ./main.go

# Указываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./main"]
