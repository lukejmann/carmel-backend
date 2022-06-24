
FROM golang:1.16-buster as builder

WORKDIR /app


# reuse cached dependencies.
COPY go.* ./
RUN go mod download

COPY . ./

RUN go build cmd/listen/listen.go


FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

ENV GOOGLE_APPLICATION_CREDENTIALS=/tmp/gcloud-credentials.json


COPY --from=builder /app/listen /listen
COPY --from=builder /app/private/gcloud-credentials.json /tmp/gcloud-credentials.json

CMD ["/listen"]