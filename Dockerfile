FROM alpine:3.10

COPY ./bin /usr/local/bin

ENTRYPOINT ["/usr/local/bin/incert"]

CMD ["version"]
