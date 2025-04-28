FROM golang:alpine

ENV GO_VERSION=1.24.1

RUN apk update && apk add --no-cache \
    wget \
    ca-certificates \
    && wget https://golang.org/dl/go$GO_VERSION.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go$GO_VERSION.linux-amd64.tar.gz \
    && rm go$GO_VERSION.linux-amd64.tar.gz \
    && ln -s /usr/local/go/bin/go /usr/bin/go

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
