## Compile executable
FROM alpine:3 AS compiler

RUN apk add --no-cache go

# Set the working directory inside the container
WORKDIR /app/HealthCheckGo

# Copy the entire project directory into the container
COPY . .

RUN CGO_ENABLED=0 go build -ldflags "-s -w" .

FROM rexezugedockerutils/upx AS upx

## Compress executable
FROM debian:12-slim AS compressor

COPY --from=upx /upx /usr/local/bin/upx

COPY --from=0 /app/HealthCheckGo/health-check-go /HealthCheck-Go

RUN upx --best --lzma /HealthCheck-Go

# Final stage
FROM busybox:stable-musl

COPY --from=compressor /HealthCheck-Go /HealthCheck-Go

# Set default command
ENTRYPOINT ["/HealthCheck-Go"]
