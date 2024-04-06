FROM golang:1.22-alpine

WORKDIR /app

RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]