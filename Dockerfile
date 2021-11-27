FROM ghcr.io/dopos/golang-alpine:v1.16.10-alpine3.14.3

WORKDIR /opt/drocopla
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.version=`git describe --tags --always`" -a .

FROM scratch

MAINTAINER Aleksei Kovrizhkin <lekovr+dopos@gmail.com>

WORKDIR /
COPY --from=0 /opt/drocopla/drocopla .

EXPOSE 3000
ENTRYPOINT ["/drocopla"]

