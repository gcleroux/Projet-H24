ARG NODE_VERSION=16.20
ARG DEBIAN_VERSION=bookworm
ARG GOLANG_VERSION=1.20.14

# FROM node:${NODE_VERSION}-alpine${ALPINE_VERSION} AS node
FROM node:${NODE_VERSION}-${DEBIAN_VERSION} AS node

FROM golang:${GOLANG_VERSION}-${DEBIAN_VERSION}

# RUN apk add --no-cache gcompat

COPY --from=node /usr/lib /usr/lib
COPY --from=node /usr/local/lib /usr/local/lib
COPY --from=node /usr/local/include /usr/local/include
COPY --from=node /usr/local/bin /usr/local/bin

WORKDIR /answer
COPY . .

RUN npm install node-pre-gyp wrtc@0.4.7 --force
RUN go build -o answer main.go

EXPOSE 50000

ENTRYPOINT [ "./answer" ]
