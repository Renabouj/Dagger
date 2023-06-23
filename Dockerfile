FROM alpine:latest

RUN mkdir /app
COPY ./dagger/build/dagger-poc /app/dagger-poc

ENTRYPOINT ["/app/dagger-poc"]