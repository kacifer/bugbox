FROM golang:1.19.5 as build

COPY . /app

WORKDIR /app

RUN go build -o /app/main .

#

FROM ubuntu:20.04

WORKDIR /app

COPY --from=build /app/main /app/main

ENTRYPOINT /app/main
