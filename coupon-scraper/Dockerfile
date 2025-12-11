## We define a builder image to perform the building step
FROM golang:1.25.5-alpine3.23 AS builder

# Install project dependencies
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# disabling CGO_ENABLED removes C dependencies, ldflags removes debug symbols
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags='-s -w' -o /app/scraper

## We define a final lightweight image that will run the app
FROM alpine:latest AS final

USER appuser
WORKDIR /app

COPY --from=builder /app/scraper /app/scraper

ENTRYPOINT ["/app/scraper"]
