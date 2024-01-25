FROM golang:1.21.6-alpine3.19 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go env -w CGO_ENABLED=0 &&\
    GOOS=js GOARCH=wasm go build -o ./static/game.wasm ./cmd/game/main.go &&\
    go build -o ./server ./cmd/server/main.go

FROM gcr.io/distroless/static as final

WORKDIR /app

COPY --from=build /app/static ./static
COPY --from=build /app/server ./

EXPOSE 80

ENTRYPOINT [ "/app/server" ]
