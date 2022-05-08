# syntax=docker/dockerfile:1

FROM golang:1.18.1-buster AS build

LABEL cloudnative="httpserver"

ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0

WORKDIR /app

COPY go.mod ./

COPY *.go ./

RUN go build -o httpserver .


FROM scratch

COPY --from=build /app/httpserver /

EXPOSE 8080

ENTRYPOINT [ "/httpserver" ]