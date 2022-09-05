# syntax=docker/dockerfile:1

## Build
FROM golang:1.18-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN ./build.sh

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /bin/server /server

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/server"]
