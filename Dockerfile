# --- Builder stage ---
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first for dependency caching
COPY go.mod ./

# Copy the rest of the source code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

# --- Final stage ---
FROM alpine:3.19

# Create a user with a known UID/GID within range 10000-20000.
RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid 10014 \
  "choreo"

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/app .
COPY --from=builder /app/response.json .

# Grant write permission to /app for the unprivileged user
RUN chown choreo:choreo /app && chmod 775 /app

# Use the above created unprivileged user
USER 10014

# Run the app
ENTRYPOINT ["./app"]
