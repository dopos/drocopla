FROM ghcr.io/dopos/golang-alpine:v1.16.10-alpine3.14.3

ENV DROCOPLA_VERSION 0.1.0
RUN apk add --no-cache git curl

WORKDIR /opt/drocopla
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.version=`git describe --tags --always`" -a ./cmd/drocopla

FROM scratch

MAINTAINER Aleksei Kovrizhkin <lekovr+dopos@gmail.com>

WORKDIR /
COPY --from=0 /opt/drocopla/drocopla .

EXPOSE 3000
ENTRYPOINT ["/drocopla"]

