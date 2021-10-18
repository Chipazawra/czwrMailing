# syntax=docker/dockerfile:1
FROM golang:1.17.1-alpine

WORKDIR $GOPATH/src/czwrMailing/

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./bin/czwrMailing ./cmd/app/main.go

EXPOSE 8080

ENTRYPOINT ["./bin/czwrMailing", "-config", "./bin/config.yml"]