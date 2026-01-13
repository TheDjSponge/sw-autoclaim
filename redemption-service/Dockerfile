## We define a builder image to perform the building step
FROM golang:alpine AS builder
ARG TARGETARCH
# Install project dependencies
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# disabling CGO_ENABLED removes C dependencies, ldflags removes debug symbols
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build -ldflags='-s -w' -o /app/redeemer



## We define a final lightweight image that will run the app
FROM alpine:latest AS final

WORKDIR /app

COPY --from=builder /app/redeemer /app/redeemer

ENTRYPOINT ["/app/redeemer"]

