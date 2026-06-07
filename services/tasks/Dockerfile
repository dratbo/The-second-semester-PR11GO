FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY internal ./internal
COPY services/tasks/cmd ./services/tasks/cmd

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/tasks ./services/tasks/cmd/tasks

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/bin/tasks /app/tasks

EXPOSE 8082

CMD ["/app/tasks"]
