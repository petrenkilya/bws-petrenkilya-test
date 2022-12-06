FROM ubuntu:latest

RUN apt update && apt install -y apache2-utils

WORKDIR /
COPY benchmark.sh .
RUN ls -la

ENTRYPOINT ./benchmark.sh