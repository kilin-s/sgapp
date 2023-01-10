FROM alpine:3.17.1


RUN useradd app && \
    mkdir -p /app && \

USER starter

COPY --chown=app --from=builder main /main

