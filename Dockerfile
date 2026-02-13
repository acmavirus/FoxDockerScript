# Copyright by AcmaTvirus
# Stage 1: Build Frontend
FROM node:18-alpine AS frontend-builder
WORKDIR /web
COPY web/package*.json ./
RUN npm install
COPY web/ .
RUN npm run build

# Stage 2: Build Backend
FROM golang:1.21-alpine AS backend-builder
WORKDIR /app
COPY . .
COPY --from=frontend-builder /web/dist /app/web/dist
RUN go mod tidy
RUN go build -o fox-admin ./cmd/fox-admin/main.go

# Stage 3: Final Image
FROM alpine:latest
WORKDIR /app
COPY --from=backend-builder /app/fox-admin .
COPY --from=backend-builder /app/templates ./templates

EXPOSE 8888
CMD ["./fox-admin"]
