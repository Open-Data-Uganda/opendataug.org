FROM golang:1.23.5-alpine AS build 

RUN apk update && apk add --no-cache git

RUN mkdir /app

WORKDIR /app

COPY go.mod go.sum templates ./

COPY . .

RUN go get ./...

RUN CGO_ENABLED=0 GOOS=linux go build -o main -ldflags="-s -w" main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /app/main .

COPY --from=build /app/templates ./templates

COPY --from=build /app/.env .

EXPOSE 8080

CMD ["./main"]