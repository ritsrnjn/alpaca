
FROM golang:1.19.0-bullseye

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o alpaca /cmd/main.go

CMD ["/alpaca"]