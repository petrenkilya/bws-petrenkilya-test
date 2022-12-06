FROM golang:latest AS build
## Prebuild
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /server/ ./...

## Deploy
FROM ubuntu

WORKDIR /

COPY --from=build /server /server
COPY entrypoint.sh /server
WORKDIR /server

ENV envListenAddrKey=:8010
EXPOSE 8010

ENTRYPOINT /server/entrypoint.sh