FROM golang:1.25-alpine AS builder

WORKDIR /app

# 1. Copy file workspace dan mod dari semua module
COPY go.work ./
COPY go.work.sum ./
COPY common/go.mod common/go.sum ./common/
COPY user-service/go.mod user-service/go.sum ./user-service/

# 2. Download dependencies with cache mount
RUN --mount=type=cache,target=/go/pkg/mod go mod download

# 3. Copy seluruh source code
COPY . .

# 4. Build aplikasi
# Perhatikan path build-nya
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./user-service/main.go

FROM alpine:latest
WORKDIR /root/

RUN apk add --no-cache tzdata

COPY --from=builder /app/main .

# Set environment variable supaya OS tahu pakai zona mana
ENV TZ=Asia/Jakarta

EXPOSE 8080
CMD ["./main"]