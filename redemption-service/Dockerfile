## We define a builder image to perform the building step
FROM golang:alpine AS builder
ARG TARGETARCH
# Install project dependencies
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# disabling CGO_ENABLED removes C dependencies, ldflags removes debug symbols
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build -ldflags='-s -w' -o /app/redeemer ./cmd/server
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build -ldflags='-s -w' -o /app/mock-hive ./cmd/mock-hive


## We define a final lightweight image that will run the app
FROM alpine:latest AS prod

WORKDIR /app

COPY --from=builder /app/redeemer /app/redeemer

ENTRYPOINT ["/app/redeemer"]

## We define another lightweight image to mock the hive server
FROM alpine:latest AS mock

WORKDIR /app

COPY --from=builder /app/mock-hive /app/mock-hive

ENTRYPOINT ["/app/mock-hive"]