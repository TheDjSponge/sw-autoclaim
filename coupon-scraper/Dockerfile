## We define a builder image to perform the building step
FROM golang:1.25.5-alpine3.23 AS builder
ARG TARGETARCH
# Install project dependencies
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# disabling CGO_ENABLED removes C dependencies, ldflags removes debug symbols
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build -ldflags='-s -w' -o /app/scraper

## We define a final lightweight image that will run the app
FROM alpine:latest AS final

WORKDIR /app
RUN apk add --no-cache dcron curl

ENV SUPERCRONIC_VERSION=v0.2.41
ARG TARGETARCH

RUN if [ "$TARGETARCH" = "amd64" ]; then \
        URL="https://github.com/aptible/supercronic/releases/download/${SUPERCRONIC_VERSION}/supercronic-linux-amd64"; \
        SHA="f70ad28d0d739a96dc9e2087ae370c257e79b8d7"; \
        SUPERCRONIC=supercronic-linux-amd64; \
    elif [ "$TARGETARCH" = "arm64" ]; then \
        URL="https://github.com/aptible/supercronic/releases/download/${SUPERCRONIC_VERSION}/supercronic-linux-arm64"; \
        SHA="44e10e33e8d98b1d1522f6719f15fb9469786ff0"; \
        SUPERCRONIC=supercronic-linux-arm64; \
    fi && \
    curl -fsSLO "$URL" \
    && echo "$SHA  $SUPERCRONIC" | sha1sum -c - \
    && chmod +x "$SUPERCRONIC" \
    && mv "$SUPERCRONIC" "/usr/local/bin/${SUPERCRONIC}" \
    && ln -s "/usr/local/bin/${SUPERCRONIC}" /usr/local/bin/supercronic

COPY --from=builder /app/scraper /app/scraper
COPY entrypoint.sh /app/entrypoint.sh
COPY scraper-cron /etc/cron.d/scraper-cron

RUN chmod +x /app/entrypoint.sh

ENTRYPOINT ["sh","/app/entrypoint.sh"]
