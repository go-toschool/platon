FROM alpine

RUN apk add --update ca-certificates

COPY bin/plato /usr/bin/plato

EXPOSE 8004

ENTRYPOINT [ "plato" ]

CMD [ "-postgres-dsn=${POSTGRES_DSN}" ]