# syntax=docker/dockerfile:1

FROM golang:1.21 AS builder
WORKDIR /server
COPY server/* .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app

FROM ubuntu:22.04 AS go-server
RUN addgroup -gid 73 go-server
RUN useradd -u 73 -g go-server go-server

RUN mkdir -p /run/go-server/certs/
RUN touch /run/go-server/certs/svid.crt /run/go-server/certs/svid.key /run/go-server/certs/root.crt
RUN chmod 600 /run/go-server/certs/svid.crt /run/go-server/certs/svid.key /run/go-server/certs/root.crt
RUN chown go-server:go-server \ 
    /run/go-server/certs/svid.crt \
    /run/go-server/certs/svid.key \
    /run/go-server/certs/root.crt

USER root
COPY --from=builder /app /app
EXPOSE 33333
ENTRYPOINT ["/app"]

