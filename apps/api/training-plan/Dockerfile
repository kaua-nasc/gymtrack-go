# Build Stage
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Install dependencies
RUN apk add --no-cache ca-certificates

# Copy the workspace configuration
COPY go.work go.work.sum ./

# Copy all go.mod and go.sum files first to leverage Docker layer caching
COPY apps/api/identity/go.mod* apps/api/identity/go.sum* ./apps/api/identity/
COPY apps/api/training-plan/go.mod* apps/api/training-plan/go.sum* ./apps/api/training-plan/
COPY libs/auth/go.mod* libs/auth/go.sum* ./libs/auth/
COPY libs/db/go.mod* libs/db/go.sum* ./libs/db/
COPY libs/email/go.mod* libs/email/go.sum* ./libs/email/
COPY libs/cache/go.mod* libs/cache/go.sum* ./libs/cache/
COPY libs/config/go.mod* libs/config/go.sum* ./libs/config/
COPY libs/storage/go.mod* libs/storage/go.sum* ./libs/storage/

RUN go mod download

COPY libs/ ./libs/
COPY apps/api/training-plan/ ./apps/api/training-plan/

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /training-plan ./apps/api/training-plan/cmd/main.go

FROM alpine:latest

RUN adduser -D -u 1000 appuser
USER appuser

WORKDIR /home/appuser/

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /training-plan .

ENV PORT=:80
EXPOSE 80

CMD ["./training-plan"]
