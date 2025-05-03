# Build stage
FROM golang:1.23-alpine3.21 AS builder
WORKDIR /app
COPY . .

# RUN apk add curl
RUN go build -o main main.go
# RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz


# Run stage
FROM alpine:3.21
WORKDIR /app
COPY --from=builder /app/main .
# COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY --from=builder /app/app.env .
COPY --from=builder /app/start.sh .
COPY db/migration ./migration

EXPOSE 8080
CMD ["/app/main"] 
#Luc nay CMD se bi override boi ENTRYPOINT, hoat dong nhu 1 paramenter cua ENTRYPOINT
ENTRYPOINT [ "/app/start.sh" ]
