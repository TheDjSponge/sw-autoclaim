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


# Latest releases available at https://github.com/aptible/supercronic/releases
ENV SUPERCRONIC_URL=https://github.com/aptible/supercronic/releases/download/v0.2.41/supercronic-linux-amd64 \
SUPERCRONIC_SHA1SUM=f70ad28d0d739a96dc9e2087ae370c257e79b8d7 \
SUPERCRONIC=supercronic-linux-amd64

RUN curl -fsSLO "$SUPERCRONIC_URL" \
&& echo "${SUPERCRONIC_SHA1SUM}  ${SUPERCRONIC}" | sha1sum -c - \
&& chmod +x "$SUPERCRONIC" \
&& mv "$SUPERCRONIC" "/usr/local/bin/${SUPERCRONIC}" \
&& ln -s "/usr/local/bin/${SUPERCRONIC}" /usr/local/bin/supercronic

COPY --from=builder /app/scraper /app/scraper
COPY entrypoint.sh /app/entrypoint.sh
COPY scraper-cron /etc/cron.d/scraper-cron

RUN chmod +x /app/entrypoint.sh

ENTRYPOINT ["sh", "/app/entrypoint.sh"]
