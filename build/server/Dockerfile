FROM golang:1.21.6-bookworm AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./server ./cmd/server/main.go

FROM gcr.io/distroless/base-debian12:debug-nonroot-amd64

LABEL org.opencontainers.image.source=https://github.com/gcleroux/Projet-H24
LABEL org.opencontainers.image.description="Testing WASM code in docker"

WORKDIR /app

COPY --from=build /app/server ./

EXPOSE 8888

ENTRYPOINT [ "/app/server" ]

