FROM golang:1.23.2-alpine3.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o ./tmp/main .

EXPOSE 8181

CMD ["./tmp/main"]