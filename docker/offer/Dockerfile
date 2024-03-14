ARG NODE_VERSION=16.20
ARG ALPINE_VERSION=3.18
ARG GOLANG_VERSION=1.20.14

FROM node:${NODE_VERSION}-alpine${ALPINE_VERSION} AS node

FROM golang:${GOLANG_VERSION}-alpine${ALPINE_VERSION}

RUN apk add --no-cache gcompat

ENV GOOS=js
ENV GOARCH=wasm

COPY --from=node /usr/lib /usr/lib
COPY --from=node /usr/local/lib /usr/local/lib
COPY --from=node /usr/local/include /usr/local/include
COPY --from=node /usr/local/bin /usr/local/bin

WORKDIR /offer
COPY . .

RUN npm install node-pre-gyp wrtc@0.4.7 --force

RUN go build -o offer main.go

ENTRYPOINT [ "./go_js_wasm_exec" ]
