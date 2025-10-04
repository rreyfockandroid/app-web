# stage: build
FROM golang:1.25-alpine AS builder
RUN apk add --no-cache git
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /out/app ./cmd/app

# stage: runtime
FROM alpine:3.20
RUN apk add --no-cache ca-certificates
COPY --from=builder /out/app /usr/local/bin/app
ENTRYPOINT ["/usr/local/bin/app"]
