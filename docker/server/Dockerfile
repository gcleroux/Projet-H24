FROM golang:1.21.6-bookworm AS build

WORKDIR /

COPY . .

RUN CGO_ENABLED=0 go build -tags server -o server

FROM gcr.io/distroless/static-debian12:nonroot-amd64

LABEL org.opencontainers.image.source=https://github.com/gcleroux/Projet-H24
LABEL org.opencontainers.image.description="Server for Golang online multiplayer game"

COPY --chown=nonroot --from=build /server /server

EXPOSE 8080

ENTRYPOINT [ "/server" ]
