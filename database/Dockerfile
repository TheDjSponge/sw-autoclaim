# /database/Dockerfile

# Use a lightweight Go image to install Goose and run the migration
FROM golang:alpine AS builder

# Install Goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# --- Final Image ---
FROM alpine:latest

# Install PostgreSQL client tools (pg_isready) for the migrate.sh script
RUN apk update && apk add --no-cache postgresql-client

# Set working directory
WORKDIR /app

# Copy the migration runner script
COPY migrate.sh /app/migrate.sh

# Copy the compiled Goose binary from the builder stage
COPY --from=builder /go/bin/goose /usr/local/bin/goose

# Copy the migration files (your SQL schema files)
COPY schema /app/migrations

# Make the script executable
RUN chmod +x /app/migrate.sh

# The default command will be to run the migration script
ENTRYPOINT ["/app/migrate.sh"]