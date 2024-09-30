FROM golang:1.22.5-alpine

WORKDIR /app

COPY . . 

RUN go mod download

RUN goose up ./db/migrations/

ENTRYPOINT ["go", "run", "./cmd/music/main.go"]
