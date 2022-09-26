FROM alpine:latest

COPY ./content ./content
COPY ./app .

ENTRYPOINT ["/app"]