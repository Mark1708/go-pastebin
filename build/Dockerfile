# Build environment
# -----------------
FROM golang:1.22-alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0

RUN apk update --no-cache gcc musl-dev && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY ../ .
RUN go build -ldflags '-w -s' -a -o ./bin/pastebin ./cmd/pastebin

# Deployment environment
# ----------------------
FROM alpine:3.19.1

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

WORKDIR /app
COPY --from=builder /build/bin/pastebin /app/pastebin
COPY --from=builder /build/bin/goose /app/goose
COPY --from=builder /build/migrations /app/migrations
COPY --from=builder /build/api/openapi/openapi.json /app/api/openapi/openapi.json

CMD ["/app/pastebin"]
EXPOSE 8080