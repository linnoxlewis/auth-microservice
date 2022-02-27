FROM golang:1.17

WORKDIR /app

COPY go.mod go.sum ./
COPY .env ./

RUN go mod download && go mod verify
RUN go build -v -o /usr/local/bin/app ./...

EXPOSE 8080

CMD ["go","run","src/main.go"]