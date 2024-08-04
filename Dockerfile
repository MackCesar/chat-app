FROM golang:1.20-alpine  AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o chat-app ./cmd

FROM alpine:3.18

WORKDIR /app

COPY --from=build /app/chat-app /app/chat-app

EXPOSE 8080

CMD ["./chat-app"]


