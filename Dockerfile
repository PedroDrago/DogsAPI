FROM golang:1.24-bookworm 
# use alpine eventually
WORKDIR /app

COPY . .

ENV GIN_MODE="release"

RUN go mod download

RUN go build -o api ./cmd/api/main.go

EXPOSE 8080

CMD ["./api"]
