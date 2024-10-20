FROM golang:1.23.2-alpine3.20

ARG USER=weather_user
ARG GROUP=weather_group

RUN addgroup -S GROUP && adduser -S USER -G GROUP

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./*.go ./

RUN go build -o main .

EXPOSE 8181

CMD ["./main"]