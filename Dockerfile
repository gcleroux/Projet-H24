FROM golang:1.21.6-bookworm AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go env -w CGO_ENABLED=0 &&\
    GOOS=js GOARCH=wasm go build -o ./static/game.wasm ./cmd/game/main.go &&\
    go build -o ./server ./cmd/server/main.go

FROM gcr.io/distroless/static-debian12:nonroot-amd64

LABEL org.opencontainers.image.source=https://github.com/gcleroux/Projet-H24
LABEL org.opencontainers.image.description="Testing WASM code in docker"

WORKDIR /app

COPY --from=build /app/static ./static
COPY --from=build /app/server ./

EXPOSE 80

ENTRYPOINT [ "/app/server" ]
