FROM golang:1.21.6-bookworm AS build

WORKDIR /

COPY . .

RUN cp "$(go env GOROOT)"/misc/wasm/wasm_exec.js assets/ &&\
    GOOS=js GOARCH=wasm go build -o assets/game.wasm internal/game/main.go &&\
    CGO_ENABLED=0 go build -tags client -o client

FROM gcr.io/distroless/static-debian12:nonroot-amd64

LABEL org.opencontainers.image.source=https://github.com/gcleroux/Projet-H24
LABEL org.opencontainers.image.description="Client for Golang online multiplayer game"

COPY --chown=nonroot --from=build /client /client

EXPOSE 8080

ENTRYPOINT [ "/client" ]
