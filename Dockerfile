FROM debian:stable-slim

EXPOSE 9134

WORKDIR /app
COPY ./bin/nccl-exporter-amd /app/exporter

ENTRYPOINT ["/app/exporter"]