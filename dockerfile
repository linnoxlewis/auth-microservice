FROM golang:1.17

WORKDIR /app

COPY go.mod go.sum ./
COPY . .

RUN go mod download && go mod verify

EXPOSE 8080

CMD ["go","run","src/main.go"]