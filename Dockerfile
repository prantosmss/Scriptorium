FROM --platform=$BUILDPLATFORM golang:1.27.1-alpine3.24 AS builder

WORKDIR /src

ENV CGO_ENABLED=0 GOWORK=off

ARG TARGETOS
ARG TARGETARCH
ARG GOPROXY=https://proxy.golang.org,direct

ENV GOPROXY=$GOPROXY

COPY go.mod go.sum ./
COPY third_party/litellm/go.mod ./third_party/litellm/go.mod
RUN go mod download

COPY . .

RUN GOOS=$TARGETOS GOARCH=$TARGETARCH \
    go build -trimpath -ldflags="-s -w" \
    -o /out/scriptorium \
    ./cmd/scriptorium

FROM alpine:3.24.1

RUN apk add --no-cache \
    ca-certificates \
    python3 \
    tzdata

WORKDIR /workspace

ENV SCRIPTORIUM_RUNS_DIR=/workspace/data/runs

COPY --from=builder /out/scriptorium /usr/local/bin/scriptorium

EXPOSE 8765

ENTRYPOINT ["scriptorium"]
