FROM golang:1.23.5-alpine3.21 AS build
ARG OS
ARG ARCH
COPY . /build/
WORKDIR /build
RUN go mod download && go build -o ncddns

FROM alpine:3.21
ARG VERSION
ARG user=ncddns
ARG group=ncddns
ARG uid=1000
ARG gid=1000
USER root
WORKDIR /app
COPY --from=build /build/ncddns /app/ncddns
COPY container-entrypoint.sh /app/container-entrypoint.sh
RUN apk update && apk --no-cache add bash && addgroup -g ${gid} ${group} && adduser -h /app -u ${uid} -G ${group} -s /bin/bash -D ${user}
RUN chown ncddns:ncddns /app/ncddns && chmod +x /app/ncddns && \
    chown ncddns:ncddns /app/container-entrypoint.sh && chmod +x /app/container-entrypoint.sh
USER ncddns
ENTRYPOINT [ "/app/container-entrypoint.sh"]