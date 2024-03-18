FROM golang:1.21.6-bookworm AS build

ARG SERVER_URL

WORKDIR /app

COPY . .

RUN GOOS=js GOARCH=wasm \
    go build \
    -o out/assets/game.wasm \
    internal/game/main.go && \
    \
    go build \
    -o out/client \
    ./cmd/client/main.go && \
    \
    cp $(go env GOROOT)/misc/wasm/wasm_exec.js out/assets/ && \
    cp -r ./static/. out/assets/

FROM gcr.io/distroless/base-debian12:debug-nonroot-amd64

LABEL org.opencontainers.image.source=https://github.com/gcleroux/Projet-H24
LABEL org.opencontainers.image.description="Testing WASM code in docker"

WORKDIR /app

COPY --chown=nonroot --from=build /app/out ./

EXPOSE 8080

ENTRYPOINT [ "/app/client" ]
