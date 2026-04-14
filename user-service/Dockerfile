# Menggunakan image Go
FROM golang:1.25-alpine

# Set working directory di dalam container
WORKDIR /app

# Copy modul common
COPY common/ /common/

# Copy file go.mod dan go.sum dari user-service
COPY user-service/go.mod user-service/go.sum ./

# Download Dependency
RUN go mod download

# Copy sisa kode user-service
COPY user-service/ .

# Build aplikasi
RUN go build -o main .

# Jalankan aplikasi
EXPOSE 8080
CMD ["./main"]
