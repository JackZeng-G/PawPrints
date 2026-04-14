# Build stage
FROM node:22-alpine AS frontend
WORKDIR /app/web
COPY web/package*.json ./
RUN npm ci
COPY web/ ./
RUN npm run build

FROM golang:1.24-alpine AS builder
WORKDIR /app
RUN apk add --no-cache git
COPY server/go.mod server/go.sum* ./
RUN go mod download || go mod tidy
COPY server/ ./
COPY --from=frontend /app/web/dist ./pkg/embedfs/dist
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o pawprints ./cmd/app/

# Runtime stage
FROM alpine:3.21
WORKDIR /app
RUN apk add --no-cache ca-certificates tzdata
COPY --from=builder /app/pawprints .
COPY config.yaml .
RUN mkdir -p data uploads
EXPOSE 8080
HEALTHCHECK --interval=30s --timeout=3s CMD wget -q --spider http://localhost:8080/api/ping || exit 1
ENTRYPOINT ["./pawprints"]