FROM golang:1.21-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main ./cmd/main.go

EXPOSE 8000

# What the container should run when it is started.
ENTRYPOINT [ "/app/main" ]
