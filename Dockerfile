FROM alpine:3.17.1


RUN addgroup -S app && adduser -S app -G app

USER app

COPY --chown=app main /main

