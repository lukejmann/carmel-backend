FROM golang:1.16-buster as builder

WORKDIR /app

COPY . .
RUN go mod download

RUN go build ./cmd/sync/sync.go

FROM debian:buster-slim

COPY --from=builder /app/sync /sync

CMD ["/sync"]