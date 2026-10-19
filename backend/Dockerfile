# ==============================================================================
# Stage 1 - Base
# ==============================================================================

FROM golang:1.25-alpine AS base

RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# ==============================================================================
# Stage 2 - Development
# ==============================================================================

FROM base AS development

RUN go install github.com/air-verse/air@latest

COPY . .

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]

# ==============================================================================
# Stage 3 - Builder
# ==============================================================================

FROM base AS builder

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
    -ldflags="-s -w" \
    -o docsai-api \
    ./cmd/api

# ==============================================================================
# Stage 4 - Production
# ==============================================================================

FROM alpine:3.22

RUN apk add --no-cache ca-certificates tzdata

RUN addgroup -S docsai && adduser -S docsai -G docsai

WORKDIR /app

COPY --from=builder /app/docsai-api .

USER docsai

EXPOSE 8080

ENTRYPOINT ["./docsai-api"]