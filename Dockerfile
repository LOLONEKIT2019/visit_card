FROM alpine:latest

COPY ./tls ./tls
COPY ./content ./content
COPY ./app .

ENTRYPOINT ["/app"]